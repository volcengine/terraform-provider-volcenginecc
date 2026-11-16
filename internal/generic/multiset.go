// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package generic

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

type multisetAttributePlanModifier struct{}

// A multiset is an array allowing non-unique items with insertion order not significant.
// Multisets do not correspond directly with either Terraform Lists (insertion order is significant) or Sets (unique items).
// Multiset Attributes are declared as Lists with a plan modifier that suppresses semantically insignificant differences.
func Multiset() planmodifier.List {
	return multisetAttributePlanModifier{}
}

func (attributePlanModifier multisetAttributePlanModifier) Description(_ context.Context) string {
	return "Suppresses semantically insignificant differences."
}

func (attributePlanModifier multisetAttributePlanModifier) MarkdownDescription(ctx context.Context) string {
	return attributePlanModifier.Description(ctx)
}

// matchesConfiguredValues compares the known values supplied by configuration,
// ignoring computed fields that are null in configuration. The comparable
// result is false when the configuration does not provide a usable match key.
func matchesConfiguredValues(configValue, plannedValue, currentValue attr.Value) (bool, bool) {
	if configValue.IsUnknown() {
		return false, true
	}
	if configValue.IsNull() {
		return false, false
	}

	configObject, isObject := configValue.(interface{ Attributes() map[string]attr.Value })
	if !isObject {
		return configValue.Equal(plannedValue) && configValue.Equal(currentValue), true
	}

	plannedObject, plannedIsObject := plannedValue.(interface{ Attributes() map[string]attr.Value })
	currentObject, currentIsObject := currentValue.(interface{ Attributes() map[string]attr.Value })
	if !plannedIsObject || !currentIsObject {
		return false, true
	}

	plannedAttributes := plannedObject.Attributes()
	currentAttributes := currentObject.Attributes()
	hasConfiguredValue := false

	for name, configAttribute := range configObject.Attributes() {
		if configAttribute.IsUnknown() {
			return false, true
		}
		if configAttribute.IsNull() {
			continue
		}

		plannedAttribute, exists := plannedAttributes[name]
		if !exists {
			return false, true
		}
		currentAttribute, exists := currentAttributes[name]
		if !exists {
			return false, true
		}

		matches, comparable := matchesConfiguredValues(configAttribute, plannedAttribute, currentAttribute)
		if !comparable {
			continue
		}
		hasConfiguredValue = true
		if !matches {
			return false, true
		}
	}

	return hasConfiguredValue, hasConfiguredValue
}

func (attributePlanModifier multisetAttributePlanModifier) PlanModifyList(ctx context.Context, request planmodifier.ListRequest, response *planmodifier.ListResponse) {
	if request.StateValue.IsNull() {
		response.PlanValue = request.PlanValue

		return
	}

	// If the current value is semantically equivalent to the planned value
	// then return the current value, else return the planned value.

	planned, diags := request.PlanValue.ToListValue(ctx)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	current, diags := request.StateValue.ToListValue(ctx)
	response.Diagnostics.Append(diags...)
	if response.Diagnostics.HasError() {
		return
	}

	if len(planned.Elements()) != len(current.Elements()) {
		response.PlanValue = request.PlanValue

		return
	}

	currentVals := make([]attr.Value, len(current.Elements()))
	copy(currentVals, current.Elements())

	configValues := request.ConfigValue.Elements()
	configMatchesPlan := !request.ConfigValue.IsNull() && !request.ConfigValue.IsUnknown() && len(configValues) == len(planned.Elements())

	for plannedIndex, plannedVal := range planned.Elements() {
		for i, currentVal := range currentVals {
			equivalent := currentVal.Equal(plannedVal)
			if configMatchesPlan {
				matches, comparable := matchesConfiguredValues(configValues[plannedIndex], plannedVal, currentVal)
				if comparable {
					equivalent = matches
				}
			}

			if equivalent {
				// Remove the match so duplicate elements are compared as a multiset.
				currentVals = append(currentVals[:i], currentVals[i+1:]...)

				break
			}
		}
	}

	if len(currentVals) == 0 {
		// Every planned value is equal to a current value.
		response.PlanValue = request.StateValue
	} else {
		response.PlanValue = request.PlanValue
	}
}
