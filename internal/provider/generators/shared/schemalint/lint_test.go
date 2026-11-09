package schemalint

import (
	"errors"
	"strings"
	"testing"

	"github.com/volcengine/terraform-provider-volcenginecc/internal/ccschema"
)

func TestValidate(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		schema   string
		baseline []string
		want     []string // "<path> [<rule>]"
	}{
		{
			name: "compliant",
			schema: `{
				"properties":{
					"InstanceId":{"type":"string"},
					"SecurityGroupIds":{"type":"array","items":{"type":"string"}},
					"Tags":{"type":"array","items":{"type":"object","properties":{"Key":{"type":"string"},"Value":{"type":"string"}}}},
					"Rules":{"type":"array","items":{"$ref":"#/definitions/Rule"}}
				},
				"definitions":{"Rule":{"type":"object","required":["Ipv6Address"],"properties":{"Ipv6Address":{"type":"string"}}}},
				"required":["InstanceId"],
				"primaryIdentifier":["/properties/InstanceId"],
				"readOnlyProperties":["/properties/Rules/*/Ipv6Address"],
				"filterProperties":{"filterableProperties":["/properties/Tags/*/Key"],"required":["/properties/InstanceId"]},
				"tagging":{"taggable":true,"tagProperty":"/properties/Tags"}
			}`,
		},
		{
			name: "nested term spelling and order",
			schema: `{
				"properties":{
					"Rules":{"type":"array","items":{"$ref":"#/definitions/Rule"}}
				},
				"definitions":{"Rule":{"type":"object","properties":{
					"SecurityGroupIDs":{"type":"array","items":{"type":"string"}},
					"AddressIpVersion":{"type":"string"},
					"key":{"type":"string"}
				}}}
			}`,
			want: []string{
				"/properties/Rules/*/AddressIpVersion [term-order]",
				"/properties/Rules/*/SecurityGroupIDs [term-spelling]",
				"/properties/Rules/*/key [pascal-case]",
			},
		},
		{
			name: "out-of-table term spelled two ways across levels",
			schema: `{
				"properties":{
					"DBName":{"type":"string"},
					"Account":{"type":"object","properties":{"DbUser":{"type":"string"}}}
				}
			}`,
			want: []string{
				"/properties/Account/DbUser [term-consistency]",
				"/properties/DBName [term-consistency]",
			},
		},
		{
			name: "new field must reuse the published spelling",
			schema: `{
				"properties":{
					"DBInstanceID":{"type":"string"},
					"DbName":{"type":"string"},
					"DBUser":{"type":"string"}
				}
			}`,
			baseline: []string{"/properties/DBInstanceID"},
			want: []string{
				"/properties/DbName [term-consistency]",
			},
		},
		{
			name: "spelling variants declared side by side",
			schema: `{
				"properties":{
					"IPAddress":{"type":"string"},
					"IpAddress":{"type":"string"}
				}
			}`,
			baseline: []string{"/properties/IPAddress"},
			want: []string{
				"/properties/IpAddress [duplicate-field]",
			},
		},
		{
			name: "json pointers",
			schema: `{
				"properties":{
					"TopicId":{"type":"string"},
					"Rules":{"type":"array","items":{"type":"object","required":["name"],"properties":{"Name":{"type":"string"}}}},
					"Config":{"type":"object","required":["Missing"],"properties":{"Mode":{"type":"string"}}}
				},
				"required":["TopicID"],
				"primaryIdentifier":["/properties/TopicID"],
				"createOnlyProperties":["/properties/Rules/0/Name"],
				"writeOnlyProperties":["/properties/Rules/Name"],
				"readOnlyProperties":["TopicId"],
				"filterProperties":{"filterableProperties":["/properties/Config/Unknown"]}
			}`,
			want: []string{
				"/properties/Config/Missing [json-pointer]",
				"/properties/Config/Unknown [json-pointer]",
				"/properties/Rules/*/name [json-pointer]",
				"/properties/Rules/0/Name [json-pointer]",
				"/properties/Rules/Name [json-pointer]",
				"/properties/TopicID [json-pointer]",
				"/properties/TopicID [json-pointer]",
				"TopicId [json-pointer]",
			},
		},
		{
			name: "taggable resource without Tags",
			schema: `{
				"properties":{"ResourceTags":{"type":"array","items":{"type":"object"}}},
				"tagging":{"taggable":true,"tagProperty":"/properties/ResourceTags"}
			}`,
			want: []string{
				"/tagging [tag-property]",
				"/tagging [tag-property]",
			},
		},
		{
			name: "baseline exempts published fields and pointers",
			schema: `{
				"properties":{
					"VPC":{"type":"string"},
					"IpList":{"type":"array","items":{"type":"string"}}
				},
				"tagging":{"taggable":true},
				"filterProperties":{"filterableProperties":["/properties/IpList/0"]}
			}`,
			baseline: []string{"/properties/VPC", "/properties/IpList/0", "/tagging"},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			resource := mustExpandedResource(t, testCase.schema)
			baseline := Baseline{"Volcengine::Test::Resource": testCase.baseline}

			var got []string
			for _, v := range Validate(resource, baseline) {
				got = append(got, v.Path+" ["+string(v.Rule)+"]")
			}

			if strings.Join(got, "\n") != strings.Join(testCase.want, "\n") {
				t.Errorf("Validate() =\n%s\nwant\n%s", strings.Join(got, "\n"), strings.Join(testCase.want, "\n"))
			}
		})
	}
}

func TestValidateMessages(t *testing.T) {
	t.Parallel()

	resource := mustExpandedResource(t, `{
		"properties":{
			"SecurityGroupIDs":{"type":"array","items":{"type":"string"}},
			"Topics":{"type":"array","items":{"type":"object","properties":{"TopicId":{"type":"string"}}}}
		},
		"readOnlyProperties":["/properties/Topics/*/TopicID"]
	}`)

	err := Check(resource, nil)

	var lintErr *Error
	if !errors.As(err, &lintErr) {
		t.Fatalf("Check() error = %v, want *Error", err)
	}
	for _, want := range []string{
		"Volcengine::Test::Resource violates 2 schema naming rule(s):",
		`/properties/SecurityGroupIDs: use the fixed term spelling "IDs" -> "Ids"; suggested name "SecurityGroupIds" [term-spelling]`,
		`/properties/Topics/*/TopicID: readOnlyProperties pointer "/properties/Topics/*/TopicID" segment "TopicID" does not match the field name "TopicId" [json-pointer]`,
	} {
		if !strings.Contains(err.Error(), want) {
			t.Errorf("Check() error =\n%s\nwant it to contain\n%s", err, want)
		}
	}
}

func mustExpandedResource(t *testing.T, properties string) *ccschema.Resource {
	t.Helper()

	document := `{"typeName":"Volcengine::Test::Resource",` + strings.TrimPrefix(strings.TrimSpace(properties), "{")
	schema, err := ccschema.NewResourceJsonSchemaDocument(document)
	if err != nil {
		t.Fatalf("NewResourceJsonSchemaDocument() error = %v", err)
	}
	resource, err := schema.Resource()
	if err != nil {
		t.Fatalf("Resource() error = %v", err)
	}
	if err := resource.Expand(); err != nil {
		t.Fatalf("Expand() error = %v", err)
	}

	return resource
}
