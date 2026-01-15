// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package codegen

import (
	"fmt"
	"io"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/hashicorp/cli"
	"github.com/volcengine/terraform-provider-volcenginecc/internal/ccschema"
	tfmaps "github.com/volcengine/terraform-provider-volcenginecc/internal/maps"
	"github.com/volcengine/terraform-provider-volcenginecc/internal/naming"
)

// Features of the emitted code.
type Features struct {
	HasRequiredRootProperty       bool // At least one root property is required.
	HasUpdatableProperty          bool // At least one property can be updated.
	HasValidator                  bool // At least one validator.
	UsesFrameworkTypes            bool // Uses a type from the terraform-plugin-framework/types package.
	UsesFrameworkJSONTypes        bool // Uses a type from the terraform-plugin-framework-jsontypes/jsontypes package.
	UsesFrameworkTimeTypes        bool // Uses a type from the terraform-plugin-framework-timetypes/timetypes package.
	UsesInternalDefaultsPackage   bool // Uses a function from the internal/defaults package.
	UsesInternalValidatorsPackage bool // Uses a function from the internal/validators package.
	UsesRegexpInValidation        bool // Uses a type from the Go standard regexp package for attribute validation.

	FrameworkDefaultsPackages     []string // Package names for any terraform-plugin-framework/resource/schema default values. May contain duplicates.
	FrameworkPlanModifierPackages []string // Package names for any terraform-plugin-framework plan modifiers. May contain duplicates.
	FrameworkValidatorsPackages   []string // Package names for any terraform-plugin-framework-validators validators. May contain duplicates.
}

func (f Features) LogicalOr(features Features) Features {
	var result Features

	result.FrameworkDefaultsPackages = slices.Clone(f.FrameworkDefaultsPackages)
	result.FrameworkDefaultsPackages = append(result.FrameworkDefaultsPackages, features.FrameworkDefaultsPackages...)
	result.FrameworkPlanModifierPackages = slices.Clone(f.FrameworkPlanModifierPackages)
	result.FrameworkPlanModifierPackages = append(result.FrameworkPlanModifierPackages, features.FrameworkPlanModifierPackages...)
	result.FrameworkValidatorsPackages = slices.Clone(f.FrameworkValidatorsPackages)
	result.FrameworkValidatorsPackages = append(result.FrameworkValidatorsPackages, features.FrameworkValidatorsPackages...)
	result.HasRequiredRootProperty = f.HasRequiredRootProperty || features.HasRequiredRootProperty
	result.HasUpdatableProperty = f.HasUpdatableProperty || features.HasUpdatableProperty
	result.HasValidator = f.HasValidator || features.HasValidator
	result.UsesInternalDefaultsPackage = f.UsesInternalDefaultsPackage || features.UsesInternalDefaultsPackage
	result.UsesInternalValidatorsPackage = f.UsesInternalValidatorsPackage || features.UsesInternalValidatorsPackage
	result.UsesFrameworkTypes = f.UsesFrameworkTypes || features.UsesFrameworkTypes
	result.UsesFrameworkJSONTypes = f.UsesFrameworkJSONTypes || features.UsesFrameworkJSONTypes
	result.UsesFrameworkTimeTypes = f.UsesFrameworkTimeTypes || features.UsesFrameworkTimeTypes
	result.UsesRegexpInValidation = f.UsesRegexpInValidation || features.UsesRegexpInValidation

	return result
}

var (
	tfMetaArguments = []string{
		"count",
		"depends_on",
		"for_each",
		//Mapped to lifecycle_config
		//"lifecycle",
		// Mapped to "provider_name".
		// "provider",
	}
)

type Emitter struct {
	CfResource   *ccschema.Resource
	IsDataSource bool
	Ui           cli.Ui
	Writer       io.Writer
}

type parent struct {
	computedAndOptional bool
	computedOnly        bool
	parentIfSetList     bool //是否为数组或者Set下的嵌套属性,会结合该属性+readOnly属性，在ListNestedAttribute或者SetNestedAttribute进行属性过滤
	path                []string
	reqd                interface {
		IsRequired(name string) bool
	}
}

// EmitRootPropertiesSchema generates the Terraform Plugin SDK code for a Cloud Control root schema
// and emits the generated code to the emitter's Writer. Code features are returned.
// The root schema is the map of root property names to Attributes.
func (e Emitter) EmitRootPropertiesSchema(tfType string, attributeNameMap map[string]string) (Features, error) {
	var features Features

	cfResource := e.CfResource
	features, err := e.emitSchema(tfType, attributeNameMap, parent{reqd: cfResource}, cfResource.Properties)

	if err != nil {
		return features, err
	}

	for name := range cfResource.Properties {
		for _, tfMetaArgument := range tfMetaArguments {
			if naming.CloudControlPropertyToTerraformAttribute(name) == tfMetaArgument {
				return features, fmt.Errorf("top-level property %s conflicts with Terraform meta-argument: %s", name, tfMetaArgument)
			}
		}

		if cfResource.IsRequired(name) {
			features.HasRequiredRootProperty = true
		}
	}

	return features, nil
}

// emitAttribute generates the Terraform Plugin SDK code for a Cloud Control property's Attributes
// and emits the generated code to the emitter's Writer. Code features are returned.
func (e Emitter) emitAttribute(tfType string, attributeNameMap map[string]string, path []string, name string, property *ccschema.Property, required, parentComputedOnly, parentComputedAndOptional bool, parentIfListOrSet bool) (Features, bool, error) {
	var features Features
	var validators []string
	var planModifiers []string
	var fwPlanModifierPackage, fwPlanModifierType, fwValidatorType string
	if parentIfListOrSet && e.CfResource.ReadOnlyProperties.ContainsPath(path) && !e.IsDataSource {
		e.printf("// Property %s was skipped because its parent is a ListNestedAttribute or SetNestedAttribute, and the property is read-only.\n", name)
		return features, true, nil
	}
	ifListOrSetNestedAttribute := false
	createOnly := e.CfResource.CreateOnlyProperties.ContainsPath(path)
	readOnly := e.CfResource.ReadOnlyProperties.ContainsPath(path)
	writeOnly := e.CfResource.WriteOnlyProperties.ContainsPath(path)
	hasDefaultValue := property.Default != nil

	if readOnly && required {
		e.warnf("%s is ReadOnly and Required", strings.Join(path, "/"))
	}
	if readOnly && writeOnly {
		e.warnf("%s is ReadOnly and WriteOnly", strings.Join(path, "/"))
	}

	var optional, computed bool

	if required && hasDefaultValue {
		e.warnf("%s is Required and has a default value. Emitting as Computed,Optional", strings.Join(path, "/"))

		required = false
		optional = true
	}

	if !required && !readOnly {
		optional = true
	}

	if (readOnly || createOnly) && !required {
		computed = true
	}

	if hasDefaultValue && !computed {
		computed = true
	}

	// All Optional attributes are also Computed.
	if optional && !computed {
		computed = true
	}

	// If our parent is Computed-only then so are we.
	if parentComputedOnly {
		required = false
		optional = false
		computed = true
	}

	var setNotNullValidator bool
	if required && parentComputedAndOptional {
		required = false
		optional = true
		computed = true
		setNotNullValidator = true
	}

	computedOnly := computed && !optional
	computedAndOptional := computed && optional

	switch propertyType := property.Type.String(); propertyType {
	//
	// Primitive types.
	//
	case ccschema.PropertyTypeBoolean:
		e.printf("schema.BoolAttribute{/*START ATTRIBUTE*/\n")
		fwPlanModifierPackage = "boolplanmodifier"
		fwPlanModifierType = "Bool"
		fwValidatorType = "Bool"

	case ccschema.PropertyTypeInteger:
		e.printf("schema.Int64Attribute{/*START ATTRIBUTE*/\n")
		fwPlanModifierPackage = "int64planmodifier"
		fwPlanModifierType = "Int64"
		fwValidatorType = "Int64"

		if f, v, err := integerValidators(path, property); err != nil {
			return features, false, err
		} else if len(v) > 0 {
			features = features.LogicalOr(f)
			validators = append(validators, v...)
		}

	case ccschema.PropertyTypeNumber:
		e.printf("schema.Float64Attribute{/*START ATTRIBUTE*/\n")
		fwPlanModifierPackage = "float64planmodifier"
		fwPlanModifierType = "Float64"
		fwValidatorType = "Float64"

		if f, v, err := numberValidators(path, property); err != nil {
			return features, false, err
		} else if len(v) > 0 {
			features = features.LogicalOr(f)
			validators = append(validators, v...)
		}

	case ccschema.PropertyTypeString:
		e.printf("schema.StringAttribute{/*START ATTRIBUTE*/\n")

		if f, c, err := stringCustomType(path, property); err != nil {
			return features, false, err
		} else if c != "" {
			features = features.LogicalOr(f)
			e.printf("CustomType:%s,\n", c)
		}

		fwPlanModifierPackage = "stringplanmodifier"
		fwPlanModifierType = "String"
		fwValidatorType = "String"

		if f, v, err := stringValidators(path, property); err != nil {
			return features, false, err
		} else if len(v) > 0 {
			features = features.LogicalOr(f)
			validators = append(validators, v...)
		}

	//
	// Complex types.
	//
	case ccschema.PropertyTypeArray:
		arrayType := aggregateType(property)

		if arrayType == aggregateSet {
			//
			// Set.
			//
			var elementType string
			var validatorsGenerator primitiveValidatorsGenerator

			fwPlanModifierPackage = "setplanmodifier"
			fwPlanModifierType = "Set"
			fwValidatorType = "Set"

			switch itemType := property.Items.Type.String(); itemType {
			case ccschema.PropertyTypeBoolean:
				elementType = "types.BoolType"

			case ccschema.PropertyTypeInteger:
				elementType = "types.Int64Type"
				validatorsGenerator = integerValidators

			case ccschema.PropertyTypeNumber:
				elementType = "types.Float64Type"
				validatorsGenerator = numberValidators

			case ccschema.PropertyTypeString:
				if f, c, err := stringCustomType(path, property.Items); err != nil {
					return features, false, err
				} else if c != "" {
					features = features.LogicalOr(f)
					elementType = c
				} else {
					elementType = "types.StringType"
				}

				validatorsGenerator = stringValidators

			case ccschema.PropertyTypeObject:
				if len(property.Items.PatternProperties) > 0 {
					return features, false, unsupportedTypeError(path, "set of key-value map")
				}

				if len(property.Items.Properties) == 0 {
					return features, false, unsupportedTypeError(path, "set of undefined schema")
				}
				ifListOrSetNestedAttribute = true

				e.printf("schema.SetNestedAttribute{/*START ATTRIBUTE*/\n")
				e.printf("NestedObject: schema.NestedAttributeObject{/*START NESTED OBJECT*/\n")
				e.printf("Attributes:")

				f, err := e.emitSchema(
					tfType,
					attributeNameMap,
					parent{
						computedAndOptional: computedAndOptional,
						computedOnly:        computedOnly,
						path:                path,
						reqd:                property.Items,
						parentIfSetList:     true,
					},
					property.Items.Properties)

				if err != nil {
					return features, false, err
				}

				features = features.LogicalOr(f)

				e.printf(",\n")
				e.printf("}/*END NESTED OBJECT*/,\n")

				if v, err := setLengthValidator(path, property); err != nil {
					return features, false, err
				} else if v != "" {
					validators = append(validators, v)
					features.FrameworkValidatorsPackages = append(features.FrameworkValidatorsPackages, "setvalidator")
				}

			default:
				return features, false, unsupportedTypeError(path, fmt.Sprintf("set of %s", itemType))
			}

			if elementType != "" {
				features.UsesFrameworkTypes = true

				e.printf("schema.SetAttribute{/*START ATTRIBUTE*/\n")
				e.printf("ElementType:%s,\n", elementType)

				if v, err := setLengthValidator(path, property); err != nil {
					return features, false, err
				} else if v != "" {
					validators = append(validators, v)
					features.FrameworkValidatorsPackages = append(features.FrameworkValidatorsPackages, "setvalidator")
				}

				if validatorsGenerator != nil {
					if f, v, err := validatorsGenerator(path, property.Items); err != nil {
						return features, false, err
					} else if len(v) > 0 {
						features = features.LogicalOr(f)

						w := &strings.Builder{}
						switch itemType := property.Items.Type.String(); itemType {
						case ccschema.PropertyTypeString:
							fprintf(w, "setvalidator.ValueStringsAre(\n")
						case ccschema.PropertyTypeInteger:
							fprintf(w, "setvalidator.ValueInt64sAre(\n")
						default:
							return features, false, fmt.Errorf("%s is of unsupported type for set item validation: %s", strings.Join(path, "/"), itemType)
						}
						for _, v := range v {
							fprintf(w, "%s,\n", v)
						}
						fprintf(w, ")")
						validators = append(validators, w.String())

						features.FrameworkValidatorsPackages = append(features.FrameworkValidatorsPackages, "setvalidator")
					}
				}
			}
		} else {
			//
			// List.
			//
			var elementType string
			var validatorsGenerator primitiveValidatorsGenerator

			fwPlanModifierPackage = "listplanmodifier"
			fwPlanModifierType = "List"
			fwValidatorType = "List"

			switch itemType := property.Items.Type.String(); itemType {
			case ccschema.PropertyTypeBoolean:
				elementType = "types.BoolType"

			case ccschema.PropertyTypeInteger:
				elementType = "types.Int64Type"
				validatorsGenerator = integerValidators

			case ccschema.PropertyTypeNumber:
				elementType = "types.Float64Type"
				validatorsGenerator = numberValidators

			case ccschema.PropertyTypeString:
				if f, c, err := stringCustomType(path, property.Items); err != nil {
					return features, false, err
				} else if c != "" {
					features = features.LogicalOr(f)
					elementType = c
				} else {
					elementType = "types.StringType"
				}

				validatorsGenerator = stringValidators

			case ccschema.PropertyTypeObject:
				if len(property.Items.PatternProperties) > 0 {
					return features, false, unsupportedTypeError(path, "list of key-value map")
				}

				if len(property.Items.Properties) == 0 {
					return features, false, unsupportedTypeError(path, "list of undefined schema")
				}
				ifListOrSetNestedAttribute = true

				e.printf("schema.ListNestedAttribute{/*START ATTRIBUTE*/\n")
				e.printf("NestedObject: schema.NestedAttributeObject{/*START NESTED OBJECT*/\n")
				e.printf("Attributes:")

				f, err := e.emitSchema(
					tfType,
					attributeNameMap,
					parent{
						computedAndOptional: computedAndOptional,
						computedOnly:        computedOnly,
						path:                path,
						reqd:                property.Items,
						parentIfSetList:     true,
					},
					property.Items.Properties)

				if err != nil {
					return features, false, err
				}

				features = features.LogicalOr(f)

				e.printf(",\n")
				e.printf("}/*END NESTED OBJECT*/,\n")

				if v, err := listLengthValidator(path, property); err != nil {
					return features, false, err
				} else if v != "" {
					validators = append(validators, v)
					features.FrameworkValidatorsPackages = append(features.FrameworkValidatorsPackages, "listvalidator")
				}

				switch arrayType {
				case aggregateOrderedSet:
					validators = append(validators, "listvalidator.UniqueValues()")
					features.FrameworkValidatorsPackages = append(features.FrameworkValidatorsPackages, "listvalidator")
				case aggregateMultiset:
					planModifiers = append(planModifiers, "generic.Multiset()")
				}

			default:
				return features, false, unsupportedTypeError(path, fmt.Sprintf("list of %s", itemType))
			}

			if elementType != "" {
				features.UsesFrameworkTypes = true

				e.printf("schema.ListAttribute{/*START ATTRIBUTE*/\n")
				e.printf("ElementType:%s,\n", elementType)

				if v, err := listLengthValidator(path, property); err != nil {
					return features, false, err
				} else if v != "" {
					validators = append(validators, v)
					features.FrameworkValidatorsPackages = append(features.FrameworkValidatorsPackages, "listvalidator")
				}

				switch arrayType {
				case aggregateOrderedSet:
					validators = append(validators, "listvalidator.UniqueValues()")
					features.FrameworkValidatorsPackages = append(features.FrameworkValidatorsPackages, "listvalidator")
				case aggregateMultiset:
					planModifiers = append(planModifiers, "generic.Multiset()")
				}

				if validatorsGenerator != nil {
					if f, v, err := validatorsGenerator(path, property.Items); err != nil {
						return features, false, err
					} else if len(v) > 0 {
						features = features.LogicalOr(f)

						w := &strings.Builder{}
						switch itemType := property.Items.Type.String(); itemType {
						case ccschema.PropertyTypeString:
							fprintf(w, "listvalidator.ValueStringsAre(\n")
						case ccschema.PropertyTypeInteger:
							fprintf(w, "listvalidator.ValueInt64sAre(\n")
						default:
							return features, false, fmt.Errorf("%s is of unsupported type for list item validation: %s", strings.Join(path, "/"), itemType)
						}
						for _, v := range v {
							fprintf(w, "%s,\n", v)
						}
						fprintf(w, ")")
						validators = append(validators, w.String())

						features.FrameworkValidatorsPackages = append(features.FrameworkValidatorsPackages, "listvalidator")
					}
				}
			}
		}

	case "":
		//
		// If the property has no specified type but has properties then assume it's an object.
		//
		fallthrough

	case ccschema.PropertyTypeObject:
		if patternProperties := property.PatternProperties; len(patternProperties) > 0 {
			//
			// Map.
			//
			if len(property.Properties) > 0 {
				return features, false, fmt.Errorf("%s has both Properties and PatternProperties", strings.Join(path, "/"))
			}

			fwPlanModifierPackage = "mapplanmodifier"
			fwPlanModifierType = "Map"
			fwValidatorType = "Map"

			// Sort the patterns to reduce diffs.
			patterns := tfmaps.Keys(patternProperties)
			sort.Strings(patterns)

			// Ignore all but the first pattern.
			pattern := patterns[0]
			patternProperty := patternProperties[pattern]

			e.printf("// Pattern: %q\n", pattern)
			switch propertyType := patternProperty.Type.String(); propertyType {
			//
			// Primitive types.
			//
			case ccschema.PropertyTypeBoolean:
				e.printf("schema.MapAttribute{/*START ATTRIBUTE*/\n")
				e.printf("ElementType:types.BoolType,\n")

				features.UsesFrameworkTypes = true

			case ccschema.PropertyTypeInteger:
				e.printf("schema.MapAttribute{/*START ATTRIBUTE*/\n")
				e.printf("ElementType:types.Int64Type,\n")

				features.UsesFrameworkTypes = true

			case ccschema.PropertyTypeNumber:
				e.printf("schema.MapAttribute{/*START ATTRIBUTE*/\n")
				e.printf("ElementType:types.Float64Type,\n")

				features.UsesFrameworkTypes = true

			case ccschema.PropertyTypeString:
				e.printf("schema.MapAttribute{/*START ATTRIBUTE*/\n")
				e.printf("ElementType:types.StringType,\n")

				features.UsesFrameworkTypes = true

			//
			// Complex types.
			//
			case ccschema.PropertyTypeArray:
				if aggregateType(patternProperty) == aggregateSet {
					switch itemType := patternProperty.Items.Type.String(); itemType {
					case ccschema.PropertyTypeBoolean:
						e.printf("schema.MapAttribute{/*START ATTRIBUTE*/\n")
						e.printf("ElementType:types.SetType{ElemType:types.BoolType},\n")

					case ccschema.PropertyTypeInteger:
						e.printf("schema.MapAttribute{/*START ATTRIBUTE*/n")
						e.printf("ElementType:types.SetType{ElemType:types.Int64Type},\n")

					case ccschema.PropertyTypeNumber:
						e.printf("schema.MapAttribute{/*START ATTRIBUTE*/\n")
						e.printf("ElementType:types.SetType{ElemType:types.Float64Type},\n")

					case ccschema.PropertyTypeString:
						e.printf("schema.MapAttribute{/*START ATTRIBUTE*/\n")
						e.printf("ElementType:types.SetType{ElemType:types.StringType},\n")

					default:
						return features, false, unsupportedTypeError(path, fmt.Sprintf("key-value map of set of %s", itemType))
					}

					features.UsesFrameworkTypes = true
				} else {
					switch itemType := patternProperty.Items.Type.String(); itemType {
					case ccschema.PropertyTypeBoolean:
						e.printf("schema.MapAttribute{/*START ATTRIBUTE*/\n")
						e.printf("ElementType:types.ListType{ElemType:types.BoolType},\n")

					case ccschema.PropertyTypeInteger:
						e.printf("schema.MapAttribute{/*START ATTRIBUTE*/\n")
						e.printf("ElementType:types.ListType{ElemType:types.Int64Type},\n")

					case ccschema.PropertyTypeNumber:
						e.printf("schema.MapAttribute{/*START ATTRIBUTE*/\n")
						e.printf("ElementType:types.ListType{ElemType:types.Float64Type},\n")

					case ccschema.PropertyTypeString:
						e.printf("schema.MapAttribute{/*START ATTRIBUTE*/\n")
						e.printf("ElementType:types.ListType{ElemType:types.StringType},\n")

					default:
						return features, false, unsupportedTypeError(path, fmt.Sprintf("key-value map of list of %s", itemType))
					}

					features.UsesFrameworkTypes = true
				}

			case ccschema.PropertyTypeObject:
				if len(patternProperty.PatternProperties) > 0 {
					return features, false, unsupportedTypeError(path, "key-value map of key-value map")
				}

				if len(patternProperty.Properties) == 0 {
					return features, false, unsupportedTypeError(path, "key-value map of undefined schema")
				}

				e.printf("schema.MapNestedAttribute{/*START ATTRIBUTE*/\n")
				e.printf("NestedObject: schema.NestedAttributeObject{/*START NESTED OBJECT*/\n")
				e.printf("Attributes:")

				f, err := e.emitSchema(
					tfType,
					attributeNameMap,
					parent{
						computedAndOptional: computedAndOptional,
						computedOnly:        computedOnly,
						path:                path,
						reqd:                property.Items,
						parentIfSetList:     parentIfListOrSet,
					},
					patternProperty.Properties)

				if err != nil {
					return features, false, err
				}

				features = features.LogicalOr(f)

				if !e.IsDataSource {
					if patternProperty.MinItems != nil {
						return features, false, fmt.Errorf("%s has unsupported MinItems", strings.Join(path, "/"))
					}
					if patternProperty.MaxItems != nil {
						return features, false, fmt.Errorf("%s has unsupported MaxItems", strings.Join(path, "/"))
					}
				}

				e.printf(",\n")
				e.printf("}/*END NESTED OBJECT*/,\n")

			default:
				return features, false, unsupportedTypeError(path, fmt.Sprintf("key-value map of %s", propertyType))
			}

			for _, pattern := range patterns[1:] {
				e.printf("// Pattern %q ignored.\n", pattern)
			}

			break
		}

		//
		// Object.
		//
		if len(property.Properties) == 0 {
			// Schemaless object => JSON string.
			features.UsesFrameworkJSONTypes = true

			e.printf("schema.StringAttribute{/*START ATTRIBUTE*/\n")
			e.printf("CustomType:jsontypes.NormalizedType{},\n")

			fwPlanModifierPackage = "stringplanmodifier"
			fwPlanModifierType = "String"
			fwValidatorType = "String"

			break
		}

		fwPlanModifierPackage = "objectplanmodifier"
		fwPlanModifierType = "Object"
		fwValidatorType = "Object"

		e.printf("schema.SingleNestedAttribute{/*START ATTRIBUTE*/\n")
		e.printf("Attributes:")
		f, err := e.emitSchema(
			tfType,
			attributeNameMap,
			parent{
				computedAndOptional: computedAndOptional,
				computedOnly:        computedOnly,
				path:                path,
				reqd:                property,
				parentIfSetList:     parentIfListOrSet,
			},
			property.Properties)

		if err != nil {
			return features, false, err
		}

		features = features.LogicalOr(f)

		e.printf(",\n")

	default:
		return features, false, unsupportedTypeError(path, propertyType)
	}

	if description := property.Description; description != nil {
		replacer := strings.NewReplacer(
			":::tip", "\n  **提示:**",
			":::", "",
			"- ", "  - ",
			"\n", "\n  ",
		)
		replaceStr := replacer.Replace(*description)
		if !e.IsDataSource && ifListOrSetNestedAttribute {
			e.printf("Description:%q,\n", replaceStr+"\n 特别提示: 在使用 SetNestedAttribute 时，必须完整定义其嵌套结构体的所有属性。若定义不完整，Terraform 在执行计划对比时可能会检测到意料之外的差异，从而触发不必要的资源更新，影响资源的稳定性与可预测性。")

		} else {
			e.printf("Description:%q,\n", replaceStr)
		}
	}

	// Return early as attribute validations are not required and additional configurations are not supported for data source.
	if e.IsDataSource {
		e.printf("Computed:true,\n")
		e.printf("}/*END ATTRIBUTE*/")

		return features, false, nil
	}

	if required {
		e.printf("Required:true,\n")
	}
	if optional {
		e.printf("Optional:true,\n")
	}
	if computed {
		e.printf("Computed:true,\n")
	}

	// Handle any default value.
	if f, defaultValue, planModifier, err := attributeDefaultValue(path, property); err != nil {
		return features, false, err
	} else {
		features = features.LogicalOr(f)
		if defaultValue != "" {
			e.printf("Default:%s,\n", defaultValue)
		}
		if planModifier != "" {
			planModifiers = append(planModifiers, planModifier)
		}
	}

	if setNotNullValidator {
		features.UsesInternalValidatorsPackage = true
		validators = append(validators, fmt.Sprintf("fwvalidators.NotNull%s()", fwValidatorType))
	}

	// Don't emit validators for Computed-only attributes.
	if !computedOnly {
		if len(validators) > 0 {
			features.HasValidator = true
			e.printf("Validators:[]validator.%s{/*START VALIDATORS*/\n", fwValidatorType)
			for _, validator := range validators {
				e.printf("%s,\n", validator)
			}
			e.printf("}/*END VALIDATORS*/,\n")
		}
	} else {
		features.FrameworkValidatorsPackages = nil
		features.UsesRegexpInValidation = false
	}

	if computed && !parentComputedOnly {
		// Computed.
		// If our parent is Computed-only (and hence we are) then we don't need our own plan modifier.
		planModifiers = append(planModifiers, fmt.Sprintf("%s.UseStateForUnknown()", fwPlanModifierPackage))
		features.FrameworkPlanModifierPackages = append(features.FrameworkPlanModifierPackages, fwPlanModifierPackage)
	}

	if createOnly {
		// ForceNew.
		if optional && computed {
			planModifiers = append(planModifiers, fmt.Sprintf("%s.RequiresReplaceIfConfigured()", fwPlanModifierPackage))
		} else {
			planModifiers = append(planModifiers, fmt.Sprintf("%s.RequiresReplace()", fwPlanModifierPackage))
		}
		features.FrameworkPlanModifierPackages = append(features.FrameworkPlanModifierPackages, fwPlanModifierPackage)
	}

	if len(planModifiers) > 0 {
		e.printf("PlanModifiers:[]planmodifier.%s{/*START PLAN MODIFIERS*/\n", fwPlanModifierType)
		for _, planModifier := range planModifiers {
			e.printf("%s,\n", planModifier)
		}
		e.printf("}/*END PLAN MODIFIERS*/,\n")
	}

	if writeOnly {
		e.printf("// %s is a write-only property.\n", name)
	}

	if !createOnly && !readOnly {
		features.HasUpdatableProperty = true
	}

	e.printf("}/*END ATTRIBUTE*/")

	return features, false, nil
}

// emitSchema generates the Terraform Plugin SDK code for a Cloud Control property's schema.
// and emits the generated code to the emitter's Writer. Code features are returned.
// A schema is a map of property names to Attributes.
// Property names are sorted prior to code generation to reduce diffs.
func (e Emitter) emitSchema(tfType string, attributeNameMap map[string]string, parent parent, properties map[string]*ccschema.Property) (Features, error) {
	names := tfmaps.Keys(properties)
	sort.Strings(names)

	var features Features

	e.printf("map[string]schema.Attribute{/*START SCHEMA*/\n")
	for _, name := range names {
		tfAttributeName := naming.CloudControlPropertyToTerraformAttribute(name)

		switch {
		case len(parent.path) == 0 && tfAttributeName == "id":
			// Terraform uses "id" as the attribute name for the resource's primary identifier.
			// If the resource has its own "Id" property, swap in a new Terraform attribute name.
			const (
				partCount = 3
			)
			parts := strings.SplitN(tfType, "_", partCount)
			// "volcenginecc_wafv2_regex_pattern_set" -> "regex_pattern_set"
			relativeTfType := parts[2]
			tfAttributeName = relativeTfType + "_id"
			if _, ok := attributeNameMap[tfAttributeName]; ok {
				return features, fmt.Errorf("top-level property %s conflicts with id", tfAttributeName)
			}
			attributeNameMap[tfAttributeName] = name

		case len(parent.path) == 0 && tfAttributeName == "provider":
			// Map "provider" to "provider_name" to avoid conflicts with the meta-argument.
			tfAttributeName = "provider_name"
			if _, ok := attributeNameMap[tfAttributeName]; ok {
				return features, fmt.Errorf("top-level property %s conflicts with provider", tfAttributeName)
			}
			attributeNameMap[tfAttributeName] = name

		case len(parent.path) == 0 && tfAttributeName == "lifecycle":
			// Map "lifecycle" to "lifecycle_config" to avoid conflicts with the meta-argument.
			tfAttributeName = "lifecycle_config"
			if _, ok := attributeNameMap[tfAttributeName]; ok {
				return features, fmt.Errorf("top-level property %s conflicts with provider", tfAttributeName)
			}
			attributeNameMap[tfAttributeName] = name

		default:
			cfPropertyName, ok := attributeNameMap[tfAttributeName]
			if ok {
				if cfPropertyName != name {
					return features, fmt.Errorf("%s overwrites %s for Terraform attribute %s", name, cfPropertyName, tfAttributeName)
				}
			} else {
				attributeNameMap[tfAttributeName] = name
			}
		}

		e.printf("// Property: %s\n", name)

		// Only dump top-level property schemas as nested properties have been expanded here.
		if len(parent.path) == 0 {
			e.printf("// Cloud Control resource type schema:\n")
			// Comment out each line.
			e.printf("%s\n", regexp.MustCompile(`(?m)^`).ReplaceAllString(fmt.Sprintf("%v", properties[name]), "// "))
		}
		schemaBuilder := e.Writer
		attributeBuilder := &strings.Builder{}
		e.Writer = attributeBuilder
		f, skipAttribute, err := e.emitAttribute(
			tfType,
			attributeNameMap,
			append(parent.path, name),
			name,
			properties[name],
			parent.reqd.IsRequired(name),
			parent.computedOnly,
			parent.computedAndOptional,
			parent.parentIfSetList,
		)

		if err != nil {
			return features, err
		}

		features = features.LogicalOr(f)
		e.Writer = schemaBuilder
		if !skipAttribute {
			e.printf("%q:", tfAttributeName)
			e.printf("%s", attributeBuilder.String())
			e.printf(",\n")
		}
	}
	e.printf("}/*END SCHEMA*/")

	return features, nil
}

// printf emits a formatted string to the underlying writer.
func (e Emitter) printf(format string, a ...interface{}) (int, error) {
	return fprintf(e.Writer, format, a...)
}

// warnf emits a formatted warning message to the UI.
func (e Emitter) warnf(format string, a ...interface{}) {
	e.Ui.Warn(fmt.Sprintf(format, a...))
}

// fprintf writes a formatted string to a Writer.
func fprintf(w io.Writer, format string, a ...interface{}) (int, error) {
	return io.WriteString(w, fmt.Sprintf(format, a...))
}

type aggregate int

const (
	aggregateNone aggregate = iota
	aggregateList
	aggregateSet
	aggregateMultiset
	aggregateOrderedSet
)

// aggregate returns the type of a Property.
func aggregateType(property *ccschema.Property) aggregate {
	if property.Type.String() != ccschema.PropertyTypeArray {
		return aggregateNone
	}

	insertionOrder := true
	if property.InsertionOrder != nil {
		insertionOrder = *property.InsertionOrder
	}
	uniqueItems := false
	if property.UniqueItems != nil {
		uniqueItems = *property.UniqueItems
	}

	if uniqueItems && !insertionOrder {
		return aggregateSet
	}

	if uniqueItems && insertionOrder {
		return aggregateOrderedSet
	}

	if !uniqueItems && !insertionOrder {
		return aggregateMultiset
	}

	return aggregateList
}

func unsupportedTypeError(path []string, typ string) error {
	return fmt.Errorf("%s is of unsupported type: %s", strings.Join(path, "/"), typ)
}

// listLengthValidator returns any list length AttributeValidator for the specified Property.
func listLengthValidator(path []string, property *ccschema.Property) (string, error) { //nolint:unparam
	if property.MinItems != nil && property.MaxItems == nil {
		return fmt.Sprintf("listvalidator.SizeAtLeast(%d)", *property.MinItems), nil
	} else if property.MinItems == nil && property.MaxItems != nil {
		return fmt.Sprintf("listvalidator.SizeAtMost(%d)", *property.MaxItems), nil
	} else if property.MinItems != nil && property.MaxItems != nil {
		return fmt.Sprintf("listvalidator.SizeBetween(%d,%d)", *property.MinItems, *property.MaxItems), nil
	}

	return "", nil
}

func setLengthValidator(path []string, property *ccschema.Property) (string, error) { //nolint:unparam
	if property.MinItems != nil && property.MaxItems == nil {
		return fmt.Sprintf("setvalidator.SizeAtLeast(%d)", *property.MinItems), nil
	} else if property.MinItems == nil && property.MaxItems != nil {
		return fmt.Sprintf("setvalidator.SizeAtMost(%d)", *property.MaxItems), nil
	} else if property.MinItems != nil && property.MaxItems != nil {
		return fmt.Sprintf("setvalidator.SizeBetween(%d,%d)", *property.MinItems, *property.MaxItems), nil
	}

	return "", nil
}

// attributeDefaultValue returns any default value for the specified Property.
func attributeDefaultValue(path []string, property *ccschema.Property) (Features, string, string, error) {
	var features Features

	switch property.Default.(type) {
	case nil:
		return features, "", "", nil
	}

	switch propertyType := property.Type.String(); propertyType {
	//
	// Primitive types.
	//
	case ccschema.PropertyTypeBoolean:
		switch v := property.Default.(type) {
		case bool:
			features.FrameworkDefaultsPackages = append(features.FrameworkDefaultsPackages, "booldefault")
			return features, fmt.Sprintf("booldefault.StaticBool(%t)", v), "", nil
		case string:
			if v, err := strconv.ParseBool(v); err != nil {
				return features, "", "", err
			} else {
				features.FrameworkDefaultsPackages = append(features.FrameworkDefaultsPackages, "booldefault")
				return features, fmt.Sprintf("booldefault.StaticBool(%t)", v), "", nil
			}
		default:
			return features, "", "", fmt.Errorf("%s (%s) has invalid default value type: %T", strings.Join(path, "/"), propertyType, v)
		}

	case ccschema.PropertyTypeInteger:
		switch v := property.Default.(type) {
		case float64:
			features.FrameworkDefaultsPackages = append(features.FrameworkDefaultsPackages, "int64default")
			return features, fmt.Sprintf("int64default.StaticInt64(%d)", int64(v)), "", nil
		default:
			return features, "", "", fmt.Errorf("%s (%s) has invalid default value type: %T", strings.Join(path, "/"), propertyType, v)
		}

	case ccschema.PropertyTypeNumber:
		switch v := property.Default.(type) {
		case float64:
			features.FrameworkDefaultsPackages = append(features.FrameworkDefaultsPackages, "float64default")
			return features, fmt.Sprintf("float64default.StaticFloat64(%f)", v), "", nil
		default:
			return features, "", "", fmt.Errorf("%s (%s) has invalid default value type: %T", strings.Join(path, "/"), propertyType, v)
		}

	case ccschema.PropertyTypeString:
		switch v := property.Default.(type) {
		case bool:
			features.FrameworkDefaultsPackages = append(features.FrameworkDefaultsPackages, "stringdefault")
			return features, fmt.Sprintf("stringdefault.StaticString(%q)", strconv.FormatBool(v)), "", nil
		case string:
			features.FrameworkDefaultsPackages = append(features.FrameworkDefaultsPackages, "stringdefault")
			return features, fmt.Sprintf("stringdefault.StaticString(%q)", v), "", nil
		default:
			return features, "", "", fmt.Errorf("%s (%s) has invalid default value type: %T", strings.Join(path, "/"), propertyType, v)
		}

	//
	// Complex types.
	//
	case ccschema.PropertyTypeArray:
		if arrayType := aggregateType(property); arrayType == aggregateSet {
			//
			// Set.
			//
			switch v := property.Default.(type) {
			case []interface{}:
				switch itemType := property.Items.Type.String(); itemType {
				case ccschema.PropertyTypeString:
					features.UsesInternalDefaultsPackage = true
					w := &strings.Builder{}
					fprintf(w, "defaults.StaticSetOfString(\n")
					for _, elem := range v {
						switch v := elem.(type) {
						case string:
							fprintf(w, "%q,\n", v)
						default:
							return features, "", "", fmt.Errorf("%s (%s/%s) has invalid default value element type: %T", strings.Join(path, "/"), propertyType, itemType, v)
						}
					}
					fprintf(w, ")")
					return features, w.String(), "", nil

				case ccschema.PropertyTypeObject:
					if len(v) == 0 {
						features.UsesInternalDefaultsPackage = true
						return features, "", "defaults.EmptySetNestedObject()", nil
					}
					return features, "", "", fmt.Errorf("%s (%s) has unsupported default value item type length (>0): %s", strings.Join(path, "/"), propertyType, itemType)

				default:
					return features, "", "", fmt.Errorf("%s (%s) has unsupported default value item type: %s", strings.Join(path, "/"), propertyType, itemType)
				}
			default:
				return features, "", "", fmt.Errorf("%s (%s) has invalid default value type: %T", strings.Join(path, "/"), propertyType, v)
			}
		} else {
			//
			// List.
			//
			switch v := property.Default.(type) {
			case []interface{}:
				switch itemType := property.Items.Type.String(); itemType {
				case ccschema.PropertyTypeString:
					features.UsesInternalDefaultsPackage = true
					w := &strings.Builder{}
					fprintf(w, "defaults.StaticListOfString(\n")
					for _, elem := range v {
						switch v := elem.(type) {
						case string:
							fprintf(w, "%q,\n", v)
						default:
							return features, "", "", fmt.Errorf("%s (%s/%s) has invalid default value element type: %T", strings.Join(path, "/"), propertyType, itemType, v)
						}
					}
					fprintf(w, ")")
					return features, w.String(), "", nil

				case ccschema.PropertyTypeObject:
					if len(v) == 0 {
						features.UsesInternalDefaultsPackage = true
						return features, "", "defaults.EmptyListNestedObject()", nil
					}
					return features, "", "", fmt.Errorf("%s (%s) has unsupported default value item type length (>0): %s", strings.Join(path, "/"), propertyType, itemType)

				default:
					return features, "", "", fmt.Errorf("%s (%s) has unsupported default value item type: %s", strings.Join(path, "/"), propertyType, itemType)
				}
			default:
				return features, "", "", fmt.Errorf("%s (%s) has invalid default value type: %T", strings.Join(path, "/"), propertyType, v)
			}
		}

	case ccschema.PropertyTypeObject:
		switch v := property.Default.(type) {
		case map[string]interface{}:
			if _, ok := v["properties"]; ok {
				// For example:
				//
				// "ReplicationSpecification": {
				// 	"type": "object",
				// 	"additionalProperties": false,
				// 	"properties": {
				// 	  "ReplicationStrategy": {
				// 		"type": "string",
				// 		"enum": [
				// 		  "SINGLE_REGION",
				// 		  "MULTI_REGION"
				// 		]
				// 	  },
				// 	  "RegionList": {
				// 		"$ref": "#/definitions/RegionList"
				// 	  }
				// 	},
				// 	"default": {
				// 	  "properties": {
				// 		"ReplicationStrategy": {
				// 		  "type": "string",
				// 		  "const": "SINGLE_REGION"
				// 		}
				// 	  }
				// 	},
				// 	"dependencies": {
				// 	  "RegionList": [
				// 		"ReplicationStrategy"
				// 	  ]
				// 	}
				// },
				//
				return features, "", "", nil
			}

			features.UsesInternalDefaultsPackage = true
			w := &strings.Builder{}
			w.WriteString("defaults.StaticPartialObject(")
			writeObjectGoLiteral(w, v)
			w.WriteString(")")
			return features, "", w.String(), nil
		default:
			return features, "", "", fmt.Errorf("%s (%s) has invalid default value type: %T", strings.Join(path, "/"), propertyType, v)
		}

	default:
		return features, "", "", fmt.Errorf("%s (%s) has unsupported default value type", strings.Join(path, "/"), propertyType)
	}
}

type primitiveValidatorsGenerator func([]string, *ccschema.Property) (Features, []string, error)

// integerValidators returns any validators for the specified integer Property.
func integerValidators(path []string, property *ccschema.Property) (Features, []string, error) {
	var features Features

	if propertyType := property.Type.String(); propertyType != ccschema.PropertyTypeInteger {
		return features, nil, fmt.Errorf("invalid property type: %s", propertyType)
	}

	var validators []string

	if property.Minimum != nil && property.Maximum == nil {
		min, err := (*property.Minimum).Int64()

		if err != nil {
			return features, nil, err
		}

		validators = append(validators, fmt.Sprintf("int64validator.AtLeast(%d)", min))
		features.FrameworkValidatorsPackages = append(features.FrameworkValidatorsPackages, "int64validator")
	} else if property.Minimum == nil && property.Maximum != nil {
		max, err := (*property.Maximum).Int64()

		if err != nil {
			return features, nil, err
		}

		validators = append(validators, fmt.Sprintf("int64validator.AtMost(%d)", max))
		features.FrameworkValidatorsPackages = append(features.FrameworkValidatorsPackages, "int64validator")
	} else if property.Minimum != nil && property.Maximum != nil {
		min, err := (*property.Minimum).Int64()

		if err != nil {
			return features, nil, err
		}

		max, err := (*property.Maximum).Int64()

		if err != nil {
			return features, nil, err
		}

		validators = append(validators, fmt.Sprintf("int64validator.Between(%d,%d)", min, max))
		features.FrameworkValidatorsPackages = append(features.FrameworkValidatorsPackages, "int64validator")
	}

	if property.Format != nil {
		if format := *property.Format; format != "int64" && format != "int32" {
			return features, nil, fmt.Errorf("%s has unsupported format: %s", strings.Join(path, "/"), format)
		}
	}

	if len(property.Enum) > 0 {
		w := &strings.Builder{}
		fprintf(w, "int64validator.OneOf(\n")
		for _, enum := range property.Enum {
			fprintf(w, "%d", int(enum.(float64)))
			fprintf(w, ",\n")
		}
		fprintf(w, ")")
		validators = append(validators, w.String())
		features.FrameworkValidatorsPackages = append(features.FrameworkValidatorsPackages, "int64validator")
	}

	return features, validators, nil
}

// numberValidators returns any validators for the specified number Property.
func numberValidators(path []string, property *ccschema.Property) (Features, []string, error) {
	var features Features

	if propertyType := property.Type.String(); propertyType != ccschema.PropertyTypeNumber {
		return features, nil, fmt.Errorf("invalid property type: %s", propertyType)
	}

	var validators []string

	if property.Minimum != nil && property.Maximum == nil {
		min, err := (*property.Minimum).Float64()

		if err != nil {
			return features, nil, err
		}

		validators = append(validators, fmt.Sprintf("float64validator.AtLeast(%f)", min))
		features.FrameworkValidatorsPackages = append(features.FrameworkValidatorsPackages, "float64validator")
	} else if property.Minimum == nil && property.Maximum != nil {
		max, err := (*property.Maximum).Float64()

		if err != nil {
			return features, nil, err
		}

		validators = append(validators, fmt.Sprintf("float64validator.AtMost(%f)", max))
		features.FrameworkValidatorsPackages = append(features.FrameworkValidatorsPackages, "float64validator")
	} else if property.Minimum != nil && property.Maximum != nil {
		min, err := (*property.Minimum).Float64()

		if err != nil {
			return features, nil, err
		}

		max, err := (*property.Maximum).Float64()

		if err != nil {
			return features, nil, err
		}

		validators = append(validators, fmt.Sprintf("float64validator.Between(%f,%f)", min, max))
		features.FrameworkValidatorsPackages = append(features.FrameworkValidatorsPackages, "float64validator")
	}

	if property.Format != nil {
		if format := *property.Format; format != "double" {
			return features, nil, fmt.Errorf("%s has unsupported format: %s", strings.Join(path, "/"), format)
		}
	}
	//if len(property.Enum) > 0 {
	//	return features, nil, fmt.Errorf("%s has enumerated values", strings.Join(path, "/"))
	//}

	return features, validators, nil
}

// stringCustomType returns any custom type for the specified string Property.
func stringCustomType(path []string, property *ccschema.Property) (Features, string, error) { //nolint:unparam
	var features Features
	var customType string

	if propertyType := property.Type.String(); propertyType != ccschema.PropertyTypeString {
		return features, customType, fmt.Errorf("invalid property type: %s", propertyType)
	}

	if property.Format != nil {
		switch format := *property.Format; format {
		case "date-time":
			features.UsesFrameworkTimeTypes = true
			customType = "timetypes.RFC3339Type{}"
		}
	}

	return features, customType, nil
}

// stringValidators returns any validators for the specified string Property.
func stringValidators(path []string, property *ccschema.Property) (Features, []string, error) {
	var features Features

	if propertyType := property.Type.String(); propertyType != ccschema.PropertyTypeString {
		return features, nil, fmt.Errorf("invalid property type: %s", propertyType)
	}

	var validators []string

	if property.MinLength != nil && property.MaxLength == nil {
		validators = append(validators, fmt.Sprintf("stringvalidator.LengthAtLeast(%d)", *property.MinLength))
		features.FrameworkValidatorsPackages = append(features.FrameworkValidatorsPackages, "stringvalidator")
	} else if property.MinLength == nil && property.MaxLength != nil {
		validators = append(validators, fmt.Sprintf("stringvalidator.LengthAtMost(%d)", *property.MaxLength))
		features.FrameworkValidatorsPackages = append(features.FrameworkValidatorsPackages, "stringvalidator")
	} else if property.MinLength != nil && property.MaxLength != nil {
		validators = append(validators, fmt.Sprintf("stringvalidator.LengthBetween(%d,%d)", *property.MinLength, *property.MaxLength))
		features.FrameworkValidatorsPackages = append(features.FrameworkValidatorsPackages, "stringvalidator")
	}

	if property.Pattern != nil && *property.Pattern != "" {
		features.UsesRegexpInValidation = true
		validators = append(validators, fmt.Sprintf("stringvalidator.RegexMatches(regexp.MustCompile(%q), \"\")", *property.Pattern))
		features.FrameworkValidatorsPackages = append(features.FrameworkValidatorsPackages, "stringvalidator")
	}

	if property.Format != nil {
		switch format := *property.Format; format {
		default:
			// TODO
			// return nil, fmt.Errorf("%s has unsupported format: %s", strings.Join(path, "/"), format)
		}
	}

	if len(property.Enum) > 0 {
		w := &strings.Builder{}
		fprintf(w, "stringvalidator.OneOf(\n")
		for _, enum := range property.Enum {
			fprintf(w, "\"")
			fprintf(w, "%s", enum.(string))
			fprintf(w, "\",\n")
		}
		fprintf(w, ")")
		validators = append(validators, w.String())
		features.FrameworkValidatorsPackages = append(features.FrameworkValidatorsPackages, "stringvalidator")
	}

	return features, validators, nil
}

func writeObjectGoLiteral(w io.Writer, obj map[string]interface{}) {
	if obj == nil {
		fprintf(w, "nil")
		return
	}

	// Sort the keys to reduce diffs.
	keys := tfmaps.Keys(obj)
	sort.Strings(keys)

	fprintf(w, "map[string]interface{}{\n")
	for _, key := range keys {
		fprintf(w, "%q:", naming.CloudControlPropertyToTerraformAttribute(key))
		switch value := obj[key]; v := value.(type) {
		case bool:
			fprintf(w, "%t", v)
		case string:
			fprintf(w, "%q", v)
		case map[string]interface{}:
			writeObjectGoLiteral(w, v)
		default:
			fprintf(w, "nil")
		}
		fprintf(w, ",\n")
	}
	fprintf(w, "}")
}
