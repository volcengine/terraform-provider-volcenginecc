package generic

import (
	"encoding/json"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func TestMergeLocalWithRemoteForSets_SetUsesRemote(t *testing.T) {
	attributes := map[string]schema.Attribute{
		"name": schema.StringAttribute{
			Required: true,
		},
		"tags": schema.SetNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"key": schema.StringAttribute{
						Required: true,
					},
					"value": schema.StringAttribute{
						Optional: true,
					},
				},
			},
			Optional: true,
		},
	}
	ccToTfNameMap := map[string]string{
		"Name": "name",
		"Tags": "tags",
	}

	local := `{"Name":"test","Tags":[{"Key":"b","Value":"2"},{"Key":"a","Value":"1"}]}`
	remote := `{"Name":"test-changed","Tags":[{"Key":"a","Value":"1"},{"Key":"b","Value":"2"}]}`

	result, err := mergeLocalWithRemoteForSets(local, remote, attributes, ccToTfNameMap)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var resultMap map[string]interface{}
	json.Unmarshal([]byte(result), &resultMap)

	// Name should come from local
	if resultMap["Name"] != "test" {
		t.Errorf("expected Name='test' (from local), got '%v'", resultMap["Name"])
	}

	// Tags (Set) should come from remote - order should be a,b (remote order)
	tags := resultMap["Tags"].([]interface{})
	firstTag := tags[0].(map[string]interface{})
	if firstTag["Key"] != "a" {
		t.Errorf("expected first tag Key='a' (from remote), got '%v'", firstTag["Key"])
	}
}

func TestMergeLocalWithRemoteForSets_NonSetUsesLocal(t *testing.T) {
	attributes := map[string]schema.Attribute{
		"name": schema.StringAttribute{
			Required: true,
		},
		"status": schema.StringAttribute{
			Computed: true,
		},
	}
	ccToTfNameMap := map[string]string{
		"Name":   "name",
		"Status": "status",
	}

	local := `{"Name":"test","Status":"running"}`
	remote := `{"Name":"test","Status":"stopped"}`

	result, err := mergeLocalWithRemoteForSets(local, remote, attributes, ccToTfNameMap)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var resultMap map[string]interface{}
	json.Unmarshal([]byte(result), &resultMap)

	// Status should come from local, not remote
	if resultMap["Status"] != "running" {
		t.Errorf("expected Status='running' (from local), got '%v'", resultMap["Status"])
	}
}

func TestMergeLocalWithRemoteForSets_NestedSetInObject(t *testing.T) {
	attributes := map[string]schema.Attribute{
		"config": schema.SingleNestedAttribute{
			Attributes: map[string]schema.Attribute{
				"name": schema.StringAttribute{
					Required: true,
				},
				"items": schema.SetNestedAttribute{
					NestedObject: schema.NestedAttributeObject{
						Attributes: map[string]schema.Attribute{
							"id": schema.StringAttribute{
								Required: true,
							},
						},
					},
					Optional: true,
				},
			},
			Optional: true,
		},
	}
	ccToTfNameMap := map[string]string{
		"Config": "config",
		"Name":   "name",
		"Items":  "items",
		"Id":     "id",
	}

	local := `{"Config":{"Name":"local-name","Items":[{"Id":"2"},{"Id":"1"}]}}`
	remote := `{"Config":{"Name":"remote-name","Items":[{"Id":"1"},{"Id":"2"}]}}`

	result, err := mergeLocalWithRemoteForSets(local, remote, attributes, ccToTfNameMap)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var resultMap map[string]interface{}
	json.Unmarshal([]byte(result), &resultMap)

	config := resultMap["Config"].(map[string]interface{})
	// Name inside object should use local
	if config["Name"] != "local-name" {
		t.Errorf("expected Config.Name='local-name', got '%v'", config["Name"])
	}
	// Items (Set) inside object should use remote order
	items := config["Items"].([]interface{})
	firstItem := items[0].(map[string]interface{})
	if firstItem["Id"] != "1" {
		t.Errorf("expected first item Id='1' (from remote), got '%v'", firstItem["Id"])
	}
}

func TestMergeLocalWithRemoteForSets_ListWithNestedSet(t *testing.T) {
	attributes := map[string]schema.Attribute{
		"roles": schema.ListNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						Required: true,
					},
					"permissions": schema.SetNestedAttribute{
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"action": schema.StringAttribute{
									Required: true,
								},
							},
						},
						Optional: true,
					},
				},
			},
			Optional: true,
		},
	}
	ccToTfNameMap := map[string]string{
		"Roles":       "roles",
		"Name":        "name",
		"Permissions": "permissions",
		"Action":      "action",
	}

	local := `{"Roles":[{"Name":"admin","Permissions":[{"Action":"write"},{"Action":"read"}]}]}`
	remote := `{"Roles":[{"Name":"admin-changed","Permissions":[{"Action":"read"},{"Action":"write"}]}]}`

	result, err := mergeLocalWithRemoteForSets(local, remote, attributes, ccToTfNameMap)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var resultMap map[string]interface{}
	json.Unmarshal([]byte(result), &resultMap)

	roles := resultMap["Roles"].([]interface{})
	role := roles[0].(map[string]interface{})

	// Name in list element should use local
	if role["Name"] != "admin" {
		t.Errorf("expected role Name='admin' (from local), got '%v'", role["Name"])
	}

	// Permissions (Set) should use remote order
	perms := role["Permissions"].([]interface{})
	firstPerm := perms[0].(map[string]interface{})
	if firstPerm["Action"] != "read" {
		t.Errorf("expected first permission Action='read' (from remote), got '%v'", firstPerm["Action"])
	}
}

func TestMergeLocalWithRemoteForSets_NoSetAttributes(t *testing.T) {
	attributes := map[string]schema.Attribute{
		"name": schema.StringAttribute{
			Required: true,
		},
		"count": schema.Int64Attribute{
			Computed: true,
		},
	}
	ccToTfNameMap := map[string]string{
		"Name":  "name",
		"Count": "count",
	}

	local := `{"Name":"test","Count":10}`
	remote := `{"Name":"test-remote","Count":20}`

	result, err := mergeLocalWithRemoteForSets(local, remote, attributes, ccToTfNameMap)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var resultMap map[string]interface{}
	json.Unmarshal([]byte(result), &resultMap)

	if resultMap["Name"] != "test" {
		t.Errorf("expected Name='test' (from local), got '%v'", resultMap["Name"])
	}
	if resultMap["Count"] != float64(10) {
		t.Errorf("expected Count=10 (from local), got '%v'", resultMap["Count"])
	}
}

func TestMergeLocalWithRemoteForSets_ReadonlyRemoteChangeIgnored(t *testing.T) {
	attributes := map[string]schema.Attribute{
		"name": schema.StringAttribute{
			Required: true,
		},
		"updated_at": schema.StringAttribute{
			Computed: true,
		},
		"tags": schema.SetNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"key": schema.StringAttribute{
						Required: true,
					},
				},
			},
			Optional: true,
		},
	}
	ccToTfNameMap := map[string]string{
		"Name":      "name",
		"UpdatedAt": "updated_at",
		"Tags":      "tags",
	}

	local := `{"Name":"test","UpdatedAt":"2024-01-01","Tags":[{"Key":"b"},{"Key":"a"}]}`
	remote := `{"Name":"test","UpdatedAt":"2024-06-15","Tags":[{"Key":"a"},{"Key":"b"}]}`

	result, err := mergeLocalWithRemoteForSets(local, remote, attributes, ccToTfNameMap)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var resultMap map[string]interface{}
	json.Unmarshal([]byte(result), &resultMap)

	// Readonly UpdatedAt should use local value, not the changed remote
	if resultMap["UpdatedAt"] != "2024-01-01" {
		t.Errorf("expected UpdatedAt='2024-01-01' (from local), got '%v'", resultMap["UpdatedAt"])
	}

	// Tags (Set) should use remote order
	tags := resultMap["Tags"].([]interface{})
	firstTag := tags[0].(map[string]interface{})
	if firstTag["Key"] != "a" {
		t.Errorf("expected first tag Key='a' (from remote), got '%v'", firstTag["Key"])
	}
}

func TestMergeLocalWithRemoteForSets_ReadOnlySetUsesLocal(t *testing.T) {
	// Read-only (Computed-only) Sets must keep local (prior state) ordering so the
	// Update side stays in sync with plannedDesiredState (which inherits prior state
	// ordering for read-only fields). Otherwise jsonpatch.CreatePatch would emit a
	// patch op for the read-only field, which CCAPI rejects.
	attributes := map[string]schema.Attribute{
		"name": schema.StringAttribute{
			Required: true,
		},
		// Writable Set
		"writable_set": schema.SetNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Required: true,
					},
				},
			},
			Optional: true,
		},
		// Read-only Set (Computed only)
		"readonly_set": schema.SetNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
				},
			},
			Computed: true,
		},
		// ListNested containing a nested read-only Set
		"items": schema.ListNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						Required: true,
					},
					"readonly_inner_set": schema.SetNestedAttribute{
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Computed: true,
								},
							},
						},
						Computed: true,
					},
				},
			},
			Optional: true,
		},
	}
	ccToTfNameMap := map[string]string{
		"Name":             "name",
		"WritableSet":      "writable_set",
		"ReadonlySet":      "readonly_set",
		"Items":            "items",
		"ReadonlyInnerSet": "readonly_inner_set",
		"Id":               "id",
	}

	local := `{"Name":"test","WritableSet":[{"Id":"b"},{"Id":"a"}],"ReadonlySet":[{"Id":"y"},{"Id":"x"}],"Items":[{"Name":"i1","ReadonlyInnerSet":[{"Id":"q"},{"Id":"p"}]}]}`
	remote := `{"Name":"test","WritableSet":[{"Id":"a"},{"Id":"b"}],"ReadonlySet":[{"Id":"x"},{"Id":"y"}],"Items":[{"Name":"i1","ReadonlyInnerSet":[{"Id":"p"},{"Id":"q"}]}]}`

	result, err := mergeLocalWithRemoteForSets(local, remote, attributes, ccToTfNameMap)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var resultMap map[string]interface{}
	json.Unmarshal([]byte(result), &resultMap)

	// Writable Set → remote order (a, b)
	ws := resultMap["WritableSet"].([]interface{})
	if ws[0].(map[string]interface{})["Id"] != "a" {
		t.Errorf("expected WritableSet[0].Id='a' (remote), got '%v'", ws[0].(map[string]interface{})["Id"])
	}

	// Read-only Set → local order (y, x)
	rs := resultMap["ReadonlySet"].([]interface{})
	if rs[0].(map[string]interface{})["Id"] != "y" {
		t.Errorf("expected ReadonlySet[0].Id='y' (local), got '%v'", rs[0].(map[string]interface{})["Id"])
	}

	// Nested read-only Set inside ListNested → local order (q, p)
	items := resultMap["Items"].([]interface{})
	inner := items[0].(map[string]interface{})["ReadonlyInnerSet"].([]interface{})
	if inner[0].(map[string]interface{})["Id"] != "q" {
		t.Errorf("expected Items[0].ReadonlyInnerSet[0].Id='q' (local), got '%v'", inner[0].(map[string]interface{})["Id"])
	}
}

func TestMergeLocalWithRemoteForSets_TagsAlwaysUseRemote(t *testing.T) {
	// Tags must always use remote data regardless of schema type,
	// because MergeSystemTags relies on remote Tags for sys:/volc: tag filtering.
	attributes := map[string]schema.Attribute{
		"name": schema.StringAttribute{
			Required: true,
		},
		"tags": schema.ListNestedAttribute{
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"key": schema.StringAttribute{
						Required: true,
					},
					"value": schema.StringAttribute{
						Optional: true,
					},
				},
			},
			Optional: true,
		},
	}
	ccToTfNameMap := map[string]string{
		"Name": "name",
		"Tags": "tags",
	}

	// Local Tags are stale and missing sys: tags
	local := `{"Name":"test","Tags":[{"Key":"env","Value":"prod"}]}`
	// Remote Tags include sys: tags added by the platform
	remote := `{"Name":"test-changed","Tags":[{"Key":"env","Value":"prod"},{"Key":"sys:created-by","Value":"system"}]}`

	result, err := mergeLocalWithRemoteForSets(local, remote, attributes, ccToTfNameMap)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var resultMap map[string]interface{}
	json.Unmarshal([]byte(result), &resultMap)

	// Name should use local
	if resultMap["Name"] != "test" {
		t.Errorf("expected Name='test' (from local), got '%v'", resultMap["Name"])
	}

	// Tags should use remote (even though it's a List type, "Tags" key is special-cased)
	tags := resultMap["Tags"].([]interface{})
	if len(tags) != 2 {
		t.Fatalf("expected 2 tags (from remote), got %d", len(tags))
	}
	secondTag := tags[1].(map[string]interface{})
	if secondTag["Key"] != "sys:created-by" {
		t.Errorf("expected second tag Key='sys:created-by' (from remote), got '%v'", secondTag["Key"])
	}
}
