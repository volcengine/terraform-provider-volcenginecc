// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package generic

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	ccdiag "github.com/volcengine/terraform-provider-volcenginecc/internal/errs/diag"
)

func TestMultiset(t *testing.T) {
	t.Parallel()

	tagAttrTypes := map[string]attr.Type{
		"key":   types.StringType,
		"value": types.StringType,
	}
	tagElemType := types.ObjectType{AttrTypes: tagAttrTypes}

	type testCase struct {
		plannedValue  basetypes.ListValue
		currentValue  basetypes.ListValue
		expectedValue basetypes.ListValue
		expectError   bool
	}
	tests := map[string]testCase{
		"null lists": {
			plannedValue:  types.ListNull(types.StringType),
			currentValue:  types.ListNull(types.StringType),
			expectedValue: types.ListNull(types.StringType),
		},
		"single item": {
			plannedValue: types.ListValueMust(types.StringType, []attr.Value{
				types.StringValue("gamma"),
			}),
			currentValue: types.ListValueMust(types.StringType, []attr.Value{
				types.StringValue("beta"),
			}),
			expectedValue: types.ListValueMust(types.StringType, []attr.Value{
				types.StringValue("gamma"),
			}),
		},
		"different lengths": {
			plannedValue: types.ListValueMust(types.StringType, []attr.Value{
				types.StringValue("gamma"),
				types.StringValue("gamma"),
			}),
			currentValue: types.ListValueMust(types.StringType, []attr.Value{
				types.StringValue("gamma"),
			}),
			expectedValue: types.ListValueMust(types.StringType, []attr.Value{
				types.StringValue("gamma"),
				types.StringValue("gamma"),
			}),
		},
		"equivalent": {
			plannedValue: types.ListValueMust(types.StringType, []attr.Value{
				types.StringValue("gamma"),
				types.StringValue("gamma"),
				types.StringValue("beta"),
			}),
			currentValue: types.ListValueMust(types.StringType, []attr.Value{
				types.StringValue("gamma"),
				types.StringValue("beta"),
				types.StringValue("gamma"),
			}),
			expectedValue: types.ListValueMust(types.StringType, []attr.Value{
				types.StringValue("gamma"),
				types.StringValue("beta"),
				types.StringValue("gamma"),
			}),
		},
		"list of objects": {
			plannedValue: types.ListValueMust(tagElemType, []attr.Value{
				types.ObjectValueMust(tagAttrTypes, map[string]attr.Value{
					"key":   types.StringValue("k2"),
					"value": types.StringValue("v2"),
				}),
				types.ObjectValueMust(tagAttrTypes, map[string]attr.Value{
					"key":   types.StringValue("k1"),
					"value": types.StringValue("v1"),
				}),
			}),
			currentValue: types.ListValueMust(tagElemType, []attr.Value{
				types.ObjectValueMust(tagAttrTypes, map[string]attr.Value{
					"key":   types.StringValue("k1"),
					"value": types.StringValue("v1"),
				}),
				types.ObjectValueMust(tagAttrTypes, map[string]attr.Value{
					"key":   types.StringValue("k2"),
					"value": types.StringValue("v2"),
				}),
			}),
			expectedValue: types.ListValueMust(tagElemType, []attr.Value{
				types.ObjectValueMust(tagAttrTypes, map[string]attr.Value{
					"key":   types.StringValue("k1"),
					"value": types.StringValue("v1"),
				}),
				types.ObjectValueMust(tagAttrTypes, map[string]attr.Value{
					"key":   types.StringValue("k2"),
					"value": types.StringValue("v2"),
				}),
			}),
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			ctx := context.TODO()
			request := planmodifier.ListRequest{
				Path:       path.Root("test"),
				PlanValue:  test.plannedValue,
				StateValue: test.currentValue,
			}
			response := planmodifier.ListResponse{}
			Multiset().PlanModifyList(ctx, request, &response)

			if !response.Diagnostics.HasError() && test.expectError {
				t.Fatal("expected error, got no error")
			}

			if response.Diagnostics.HasError() && !test.expectError {
				t.Fatalf("got unexpected error: %s", ccdiag.DiagnosticsError(response.Diagnostics))
			}

			if diff := cmp.Diff(response.PlanValue, test.expectedValue); diff != "" {
				t.Errorf("unexpected diff (+wanted, -got): %s", diff)
			}
		})
	}
}

func TestMultisetUsesConfiguredObjectAttributes(t *testing.T) {
	t.Parallel()

	nodeAttributeTypes := map[string]attr.Type{
		"node_id":   types.StringType,
		"node_type": types.StringType,
		"node_spec": types.StringType,
		"zone_id":   types.StringType,
	}
	nodeType := types.ObjectType{AttrTypes: nodeAttributeTypes}
	node := func(id attr.Value, nodeType, nodeSpec, zoneID string) attr.Value {
		return types.ObjectValueMust(nodeAttributeTypes, map[string]attr.Value{
			"node_id":   id,
			"node_type": types.StringValue(nodeType),
			"node_spec": types.StringValue(nodeSpec),
			"zone_id":   types.StringValue(zoneID),
		})
	}
	configuredNode := func(nodeType, nodeSpec, zoneID string) attr.Value {
		return node(types.StringNull(), nodeType, nodeSpec, zoneID)
	}

	stateValue := types.ListValueMust(nodeType, []attr.Value{
		node(types.StringValue("node-a"), "Primary", "spec-a", "zone-a"),
		node(types.StringValue("node-b"), "Secondary", "spec-b", "zone-b"),
	})

	tests := map[string]struct {
		plan     basetypes.ListValue
		config   basetypes.ListValue
		expected basetypes.ListValue
	}{
		"reordered elements with stale computed identities": {
			plan: types.ListValueMust(nodeType, []attr.Value{
				node(types.StringValue("node-a"), "Secondary", "spec-b", "zone-b"),
				node(types.StringValue("node-b"), "Primary", "spec-a", "zone-a"),
			}),
			config: types.ListValueMust(nodeType, []attr.Value{
				configuredNode("Secondary", "spec-b", "zone-b"),
				configuredNode("Primary", "spec-a", "zone-a"),
			}),
			expected: stateValue,
		},
		"reordered elements with unknown computed identities": {
			plan: types.ListValueMust(nodeType, []attr.Value{
				node(types.StringUnknown(), "Secondary", "spec-b", "zone-b"),
				node(types.StringUnknown(), "Primary", "spec-a", "zone-a"),
			}),
			config: types.ListValueMust(nodeType, []attr.Value{
				configuredNode("Secondary", "spec-b", "zone-b"),
				configuredNode("Primary", "spec-a", "zone-a"),
			}),
			expected: stateValue,
		},
		"configured change is not suppressed": {
			plan: types.ListValueMust(nodeType, []attr.Value{
				node(types.StringUnknown(), "Primary", "spec-updated", "zone-a"),
				node(types.StringUnknown(), "Secondary", "spec-b", "zone-b"),
			}),
			config: types.ListValueMust(nodeType, []attr.Value{
				configuredNode("Primary", "spec-updated", "zone-a"),
				configuredNode("Secondary", "spec-b", "zone-b"),
			}),
			expected: types.ListValueMust(nodeType, []attr.Value{
				node(types.StringUnknown(), "Primary", "spec-updated", "zone-a"),
				node(types.StringUnknown(), "Secondary", "spec-b", "zone-b"),
			}),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			response := planmodifier.ListResponse{}
			Multiset().PlanModifyList(context.Background(), planmodifier.ListRequest{
				Path:        path.Root("node_info"),
				ConfigValue: test.config,
				PlanValue:   test.plan,
				StateValue:  stateValue,
			}, &response)

			if response.Diagnostics.HasError() {
				t.Fatalf("unexpected diagnostics: %s", ccdiag.DiagnosticsError(response.Diagnostics))
			}
			if diff := cmp.Diff(response.PlanValue, test.expected); diff != "" {
				t.Errorf("unexpected plan value (+wanted, -got): %s", diff)
			}
		})
	}
}
