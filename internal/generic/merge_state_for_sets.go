package generic

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// mergeLocalWithRemoteForSets merges local state JSON with remote state JSON.
// For Set type attributes, the remote value is used (to preserve correct ordering).
// For all other types, the local value is used (to avoid readonly attribute changes causing unexpected diffs).
func mergeLocalWithRemoteForSets(localState, remoteState string, attributes map[string]schema.Attribute, ccToTfNameMap map[string]string) (string, error) {
	var localMap, remoteMap map[string]interface{}
	if err := json.Unmarshal([]byte(localState), &localMap); err != nil {
		return "", err
	}
	if err := json.Unmarshal([]byte(remoteState), &remoteMap); err != nil {
		return "", err
	}

	result := mergeMapForSets(localMap, remoteMap, attributes, ccToTfNameMap)

	// Top-level Tags must always use remote data because MergeSystemTags relies on
	// remote Tags to correctly handle sys:/volc: system tag filtering.
	if remoteTags, ok := remoteMap["Tags"]; ok {
		result["Tags"] = remoteTags
	}

	merged, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(merged), nil
}

// mergeMapForSets recursively merges two JSON object maps.
// Set type attributes use the remote value; all other types use the local value.
// For nested structures (List/Object/Map), it recurses to handle deeply nested Set attributes.
func mergeMapForSets(local, remote map[string]interface{}, attributes map[string]schema.Attribute, ccToTfNameMap map[string]string) map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range local {
		result[k] = v
	}

	for key, remoteVal := range remote {
		tfKey, ok := ccToTfNameMap[key]
		if !ok {
			continue
		}
		attrSchema, ok := attributes[tfKey]
		if !ok {
			continue
		}

		switch attrSchema.GetType().(type) {
		case types.SetType:
			// Use remote value for Set types to preserve correct ordering
			result[key] = remoteVal

		case types.ListType:
			if nested, ok := attrSchema.(schema.ListNestedAttribute); ok {
				localArr, lok := local[key].([]interface{})
				remoteArr, rok := remoteVal.([]interface{})
				if lok && rok {
					result[key] = mergeArrayForSets(localArr, remoteArr, nested.NestedObject.Attributes, ccToTfNameMap)
				}
			}

		case types.ObjectType:
			if nested, ok := attrSchema.(schema.SingleNestedAttribute); ok {
				localObj, lok := local[key].(map[string]interface{})
				remoteObj, rok := remoteVal.(map[string]interface{})
				if lok && rok {
					result[key] = mergeMapForSets(localObj, remoteObj, nested.Attributes, ccToTfNameMap)
				}
			}

		case types.MapType:
			if nested, ok := attrSchema.(schema.MapNestedAttribute); ok {
				localMap, lok := local[key].(map[string]interface{})
				remoteMap, rok := remoteVal.(map[string]interface{})
				if lok && rok {
					mergedMap := make(map[string]interface{})
					for mk, mv := range localMap {
						mergedMap[mk] = mv
					}
					for mk, mv := range remoteMap {
						if localItem, exists := localMap[mk]; exists {
							li, lok2 := localItem.(map[string]interface{})
							ri, rok2 := mv.(map[string]interface{})
							if lok2 && rok2 {
								mergedMap[mk] = mergeMapForSets(li, ri, nested.NestedObject.Attributes, ccToTfNameMap)
							}
						}
					}
					result[key] = mergedMap
				}
			}

		default:
			// Scalar/other types: keep local value (already in result)
		}
	}

	return result
}

// mergeArrayForSets merges two JSON arrays element-by-element.
// Uses the local array as the base, recursively processing each element for nested Set attributes.
func mergeArrayForSets(local, remote []interface{}, attributes map[string]schema.Attribute, ccToTfNameMap map[string]string) []interface{} {
	result := make([]interface{}, len(local))
	for i, lv := range local {
		localObj, lok := lv.(map[string]interface{})
		if i < len(remote) {
			remoteObj, rok := remote[i].(map[string]interface{})
			if lok && rok {
				result[i] = mergeMapForSets(localObj, remoteObj, attributes, ccToTfNameMap)
			} else {
				result[i] = lv
			}
		} else {
			result[i] = lv
		}
	}
	return result
}
