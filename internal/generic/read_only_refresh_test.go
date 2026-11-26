// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package generic

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/volcengine/terraform-provider-volcenginecc/internal/cloudcontrol"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

type refreshTestProvider struct {
	client *cloudcontrol.CloudControl
	region string
}

func (p refreshTestProvider) CloudControlAPIClient(context.Context) *cloudcontrol.CloudControl {
	return p.client
}

func (p refreshTestProvider) Region(context.Context) string {
	return p.region
}

func (p refreshTestProvider) RegisterLogger(ctx context.Context) context.Context {
	return ctx
}

func TestPopulateUnknownValuesRefreshesKnownReadOnlySet(t *testing.T) {
	ctx := context.Background()
	var getResourceCalls atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		if action := request.URL.Query().Get("Action"); action != "GetResource" {
			http.Error(w, "unexpected action", http.StatusBadRequest)
			return
		}
		getResourceCalls.Add(1)
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"ResponseMetadata": map[string]interface{}{"RequestId": "local-test"},
			"Result": map[string]interface{}{
				"TypeName": "Test::Instance",
				"ResourceDescription": map[string]interface{}{
					"Identifier": "instance-1",
					"Properties": `{"InstanceId":"instance-1","Endpoints":[{"Connection":"fresh"}]}`,
				},
			},
		})
	}))
	defer server.Close()

	sdkConfig := volcengine.NewConfig().
		WithRegion("cn-test").
		WithEndpoint(server.URL).
		WithDisableSSL(true).
		WithCredentials(credentials.NewStaticCredentials("test-access-key", "test-secret-key", "")).
		WithHTTPClient(server.Client())
	sdkSession, err := session.NewSession(sdkConfig)
	if err != nil {
		t.Fatalf("creating local test session: %v", err)
	}

	provider := refreshTestProvider{
		client: cloudcontrol.New(sdkSession),
		region: "cn-test",
	}
	endpointObjectTypes := map[string]attr.Type{"connection": types.StringType}
	endpointObjectType := types.ObjectType{AttrTypes: endpointObjectTypes}
	stateSchema := schema.Schema{Attributes: map[string]schema.Attribute{
		"id": schema.StringAttribute{Computed: true},
		"endpoints": schema.SetNestedAttribute{
			Computed: true,
			NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{
				"connection": schema.StringAttribute{Computed: true},
			}},
		},
	}}
	initialEndpoints := types.SetValueMust(endpointObjectType, []attr.Value{
		types.ObjectValueMust(endpointObjectTypes, map[string]attr.Value{"connection": types.StringValue("stale")}),
	})
	rootTypes := map[string]attr.Type{
		"id":        types.StringType,
		"endpoints": types.SetType{ElemType: endpointObjectType},
	}
	initialStateValue := types.ObjectValueMust(rootTypes, map[string]attr.Value{
		"id":        types.StringValue("instance-1"),
		"endpoints": initialEndpoints,
	})
	initialStateRaw, err := initialStateValue.ToTerraformValue(ctx)
	if err != nil {
		t.Fatalf("converting initial state: %v", err)
	}
	state := tfsdk.State{Schema: stateSchema, Raw: initialStateRaw}
	endpointsPath := path.Root("endpoints")

	resource := &genericResource{
		ccTypeName: "Test::Instance",
		ccToTfNameMap: map[string]string{
			"InstanceId": "id",
			"Endpoints":  "endpoints",
			"Connection": "connection",
		},
		readOnlyAttributePaths: []*path.Path{&endpointsPath},
		provider:               provider,
	}

	withoutRefresh := tfsdk.State{Schema: stateSchema, Raw: initialStateRaw}
	diags := resource.populateUnknownValues(ctx, "instance-1", &withoutRefresh, false)
	if diags.HasError() {
		t.Fatalf("populateUnknownValues without refresh returned diagnostics: %s", diags)
	}
	if got := getResourceCalls.Load(); got != 0 {
		t.Fatalf("GetResource calls without refresh = %d, want 0", got)
	}

	diags = resource.populateUnknownValues(ctx, "instance-1", &state, true)
	if diags.HasError() {
		t.Fatalf("populateUnknownValues with refresh returned diagnostics: %s", diags)
	}
	if got := getResourceCalls.Load(); got != 1 {
		t.Fatalf("GetResource calls = %d, want 1", got)
	}

	var endpoints types.Set
	if diags := state.GetAttribute(ctx, endpointsPath, &endpoints); diags.HasError() {
		t.Fatalf("reading refreshed endpoints: %s", diags)
	}
	values := endpoints.Elements()
	if len(values) != 1 {
		t.Fatalf("endpoint count = %d, want 1", len(values))
	}
	endpoint := values[0].(types.Object)
	connection := endpoint.Attributes()["connection"].(types.String)
	if got := connection.ValueString(); got != "fresh" {
		t.Fatalf("endpoint connection = %q, want fresh", got)
	}
}
