// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package shared

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/hashicorp/cli"
	"github.com/volcengine/terraform-provider-volcenginecc/internal/ccschema"
	"github.com/volcengine/terraform-provider-volcenginecc/internal/naming"
	"github.com/volcengine/terraform-provider-volcenginecc/internal/provider/generators/shared/codegen"
	"github.com/volcengine/terraform-provider-volcenginecc/internal/provider/generators/shared/schemalint"
	tfslices "github.com/volcengine/terraform-provider-volcenginecc/internal/slices"
)

const (
	DataSourceType = "DataSource"
	ResourceType   = "Resource"
)

// GenerateTemplateData generates the templates body from the Resource
// constructed from a Cloud Control type's Schema file.
// This method can be applied to both singular data source and resource types.
func GenerateTemplateData(ui cli.Ui, cfTypeSchemaFile, resType, tfResourceType, packageName string) (*TemplateData, error) {
	resource, err := NewResource(tfResourceType, cfTypeSchemaFile)

	if err != nil {
		return nil, fmt.Errorf("reading Cloud Control resource schema for %s: %w", tfResourceType, err)
	}

	cfTypeName := *resource.CfResource.TypeName
	org, svc, res, err := naming.ParseCloudControlTypeName(cfTypeName)

	if err != nil {
		return nil, fmt.Errorf("incorrect format for Cloud Control Resource Provider Schema type name: %s", cfTypeName)
	}

	// e.g. "logGroupResource" or "logGroupDataSource"
	factoryFunctionName := string(bytes.ToLower([]byte(res[:1]))) + res[1:] + resType

	// e.g. "TestAccLogsLogGroup"
	acceptanceTestFunctionPrefix := fmt.Sprintf("TestAcc%[1]s%[2]s%[3]s", org, svc, res)

	sb := strings.Builder{}
	codeEmitter := codegen.Emitter{
		CfResource:   resource.CfResource,
		IsDataSource: resType == DataSourceType,
		Ui:           ui,
		Writer:       &sb,
	}

	// Generate code for the Cloud Control root properties schema.
	attributeNameMap := make(map[string]string) // Terraform attribute name to Cloud Control property name.
	codeFeatures, err := codeEmitter.EmitRootPropertiesSchema(resource.TfType, attributeNameMap)

	if err != nil {
		return nil, fmt.Errorf("emitting schema code: %w", err)
	}

	rootPropertiesSchema := sb.String()
	sb.Reset()

	templateData := &TemplateData{
		AcceptanceTestFunctionPrefix: acceptanceTestFunctionPrefix,
		AttributeNameMap:             attributeNameMap,
		CloudControlTypeName:         cfTypeName,
		FactoryFunctionName:          factoryFunctionName,
		HasRequiredAttribute:         true,
		PackageName:                  packageName,
		RootPropertiesSchema:         rootPropertiesSchema,
		SchemaVersion:                1,
		TerraformTypeName:            resource.TfType,
	}

	if !codeFeatures.HasRequiredRootProperty {
		templateData.HasRequiredAttribute = false
	}
	if codeFeatures.UsesFrameworkTypes {
		templateData.ImportFrameworkTypes = true
	}
	if codeFeatures.UsesFrameworkJSONTypes {
		templateData.ImportFrameworkJSONTypes = true
	}
	if codeFeatures.UsesFrameworkTimeTypes {
		templateData.ImportFrameworkTimeTypes = true
	}
	if codeFeatures.UsesInternalDefaultsPackage {
		templateData.ImportInternalDefaults = true
	}
	if codeFeatures.HasValidator {
		templateData.ImportFrameworkValidator = true
	}
	if codeFeatures.UsesInternalValidatorsPackage {
		templateData.ImportInternalValidators = true
	}

	if resType == DataSourceType {
		templateData.SchemaDescription = fmt.Sprintf("Data Source schema for %s", cfTypeName)

		return templateData, nil
	}

	templateData.HasUpdateMethod = true

	if !codeFeatures.HasUpdatableProperty {
		templateData.HasUpdateMethod = false
	}
	if codeFeatures.UsesRegexpInValidation {
		templateData.ImportRegexp = true
	}

	if description := resource.CfResource.Description; description != nil {
		templateData.SchemaDescription = *description
	}

	for _, path := range resource.CfResource.WriteOnlyProperties {
		templateData.WriteOnlyPropertyPaths = append(templateData.WriteOnlyPropertyPaths, string(path))
	}

	for _, path := range resource.CfResource.ReadOnlyProperties {
		templateData.ReadOnlyPropertyPaths = append(templateData.ReadOnlyPropertyPaths, string(path))
	}

	for _, path := range resource.CfResource.CreateOnlyProperties {
		templateData.CreateOnlyPropertyPaths = append(templateData.CreateOnlyPropertyPaths, string(path))
	}

	for _, path := range resource.CfResource.PrimaryIdentifier {
		templateData.PrimaryIdentifier = append(templateData.PrimaryIdentifier, string(path))
	}
	templateData.CollectionIdentities = collectCollectionIdentities(resource.CfResource.Properties, nil)

	if v, ok := resource.CfResource.Handlers[ccschema.HandlerTypeCreate]; ok {
		templateData.CreateTimeoutInMinutes = v.TimeoutInMinutes
	}
	if v, ok := resource.CfResource.Handlers[ccschema.HandlerTypeUpdate]; ok {
		templateData.UpdateTimeoutInMinutes = v.TimeoutInMinutes
	}
	if v, ok := resource.CfResource.Handlers[ccschema.HandlerTypeDelete]; ok {
		templateData.DeleteTimeoutInMinutes = v.TimeoutInMinutes
	}

	templateData.FrameworkDefaultsPackages = tfslices.AppendUnique(templateData.FrameworkDefaultsPackages, codeFeatures.FrameworkDefaultsPackages...)
	templateData.FrameworkPlanModifierPackages = []string{"stringplanmodifier"} // For the 'id' attribute.
	templateData.FrameworkPlanModifierPackages = tfslices.AppendUnique(templateData.FrameworkPlanModifierPackages, codeFeatures.FrameworkPlanModifierPackages...)
	templateData.FrameworkValidatorsPackages = tfslices.AppendUnique(templateData.FrameworkValidatorsPackages, codeFeatures.FrameworkValidatorsPackages...)

	return templateData, nil
}

type TemplateData struct {
	AcceptanceTestFunctionPrefix  string
	AttributeNameMap              map[string]string
	CloudControlTypeName          string
	CreateTimeoutInMinutes        int
	DeleteTimeoutInMinutes        int
	FactoryFunctionName           string
	FrameworkDefaultsPackages     []string
	FrameworkPlanModifierPackages []string
	FrameworkValidatorsPackages   []string
	HasRequiredAttribute          bool
	HasUpdateMethod               bool
	ImportFrameworkTypes          bool
	ImportFrameworkJSONTypes      bool
	ImportFrameworkTimeTypes      bool
	ImportFrameworkValidator      bool
	ImportInternalDefaults        bool
	ImportInternalValidators      bool
	ImportRegexp                  bool
	PackageName                   string
	PrimaryIdentifier             []string
	RootPropertiesSchema          string
	SchemaDescription             string
	SchemaVersion                 int64
	TerraformTypeName             string
	UpdateTimeoutInMinutes        int
	WriteOnlyPropertyPaths        []string
	ReadOnlyPropertyPaths         []string
	CreateOnlyPropertyPaths       []string
	CollectionIdentities          []CollectionIdentity
}

// CollectionIdentity contains schema-derived metadata emitted into ResourceOptions.
type CollectionIdentity struct {
	PropertyPath    string
	IdentifierPaths []string
	UniqueItems     bool
}

// collectCollectionIdentities walks expanded properties in stable name order,
// retains wildcard segments below object arrays, and emits children before
// parents so every consumer observes the required inner-first sort order.
func collectCollectionIdentities(properties map[string]*ccschema.Property, parentPath []string) []CollectionIdentity {
	var result []CollectionIdentity
	names := make([]string, 0, len(properties))
	for name := range properties {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		property := properties[name]
		if property == nil {
			continue
		}

		propertyPath := append(append([]string(nil), parentPath...), name)
		switch property.Type.String() {
		case ccschema.PropertyTypeObject, "":
			result = append(result, collectCollectionIdentities(property.Properties, propertyPath)...)
		case ccschema.PropertyTypeArray:
			if property.Items != nil && property.Items.Type.String() == ccschema.PropertyTypeObject {
				itemPath := append(append([]string(nil), propertyPath...), "*")
				result = append(result, collectCollectionIdentities(property.Items.Properties, itemPath)...)
			}
		}

		if len(property.ElementIdentifier) > 0 {
			uniqueItems := property.UniqueItems != nil && *property.UniqueItems
			result = append(result, CollectionIdentity{
				PropertyPath:    "/" + strings.Join(propertyPath, "/"),
				IdentifierPaths: append([]string(nil), property.ElementIdentifier...),
				UniqueItems:     uniqueItems,
			})
		}
	}

	return result
}

// namingBaseline loads the embedded schema naming baseline once per generator run.
var namingBaseline = sync.OnceValues(schemalint.DefaultBaseline)

type Resource struct {
	CfResource *ccschema.Resource
	TfType     string
}

// NewResource creates a Resource type
// from the corresponding resource's Cloud Control Schema file
func NewResource(resourceType, cfTypeSchemaFile string) (*Resource, error) {
	resourceSchema, err := ccschema.NewResourceJsonSchemaPath(cfTypeSchemaFile)

	if err != nil {
		return nil, fmt.Errorf("reading Cloud Control Resource Type Schema: %w", err)
	}

	resource, err := resourceSchema.Resource()

	if err != nil {
		return nil, fmt.Errorf("parsing Cloud Control Resource Type Schema: %w", err)
	}

	if err := resource.Expand(); err != nil {
		return nil, fmt.Errorf("expanding JSON Pointer references: %w", err)
	}
	if err := resource.ValidateCollectionIdentities(); err != nil {
		return nil, fmt.Errorf("validating collection identities: %w", err)
	}

	baseline, err := namingBaseline()
	if err != nil {
		return nil, err
	}
	if err := schemalint.Check(resource, baseline); err != nil {
		return nil, fmt.Errorf("validating schema naming: %w", err)
	}

	return &Resource{
		CfResource: resource,
		TfType:     resourceType,
	}, nil
}
