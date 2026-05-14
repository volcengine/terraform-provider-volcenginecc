// Copyright (c) 2025 Beijing Volcano Engine Technology Co., Ltd.
// SPDX-License-Identifier: MPL-2.0

// Hand-written customization for `volcenginecc_iam_accesskey`.
//
// The Volcengine Cloud Control API only returns `SecretAccessKey` from
// CreateResource (and even that is currently subject to a server-side fix —
// see commit history). GetResource never returns it. The default flow in
// `internal/generic/resource.go` therefore loses the secret immediately after
// creation.
//
// This file registers a custom override that:
//   - During Create, captures `SecretAccessKey` from the CreateResource event's
//     ResourceModel and writes it into state.
//   - During Read, restores `secret_access_key` from prior state because the
//     auto-generated Read overwrites it with null from GetResource.
//
// Update/Delete/Metadata/Schema/Configure/ImportState/ConfigValidators all
// delegate to the auto-generated inner resource. The override is registered via
// `customresources.Register` so the auto-generated `accesskey_resource_gen.go`
// stays untouched and survives `make resources`.

package iam

import (
	"context"
	"encoding/json"
	"fmt"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/volcengine/terraform-provider-volcenginecc/internal/base"
	"github.com/volcengine/terraform-provider-volcenginecc/internal/cloudcontrol"
	"github.com/volcengine/terraform-provider-volcenginecc/internal/customresources"
	"github.com/volcengine/terraform-provider-volcenginecc/internal/generic"
	tfcloudcontrol "github.com/volcengine/terraform-provider-volcenginecc/internal/service/cloudcontrol"
	"github.com/volcengine/terraform-provider-volcenginecc/internal/util"
)

const (
	accesskeyTerraformTypeName    = "volcenginecc_iam_accesskey"
	accesskeyCloudControlTypeName = "Volcengine::IAM::Accesskey"
	accesskeySecretTerraformField = "secret_access_key"
)

// Mirror of the WithAttributeNameMap call in the auto-generated
// accesskey_resource_gen.go. Duplicated here so this override stays
// self-contained; if the upstream schema changes, this mapping must be
// updated alongside. The "id" → "ID" entry is intentionally omitted because
// the generic framework synthesizes it for the reserved top-level identifier
// (see resourceWithAttributeNameMap in internal/generic/resource.go); the
// translator never sees a known "id" value during Create either way.
var accesskeyTfToCcNameMap = map[string]string{
	"access_key_id":     "AccessKeyId",
	"created_time":      "CreatedTime",
	"last_login_date":   "LastLoginDate",
	"region":            "Region",
	"request_time":      "RequestTime",
	"secret_access_key": "SecretAccessKey",
	"service":           "Service",
	"status":            "Status",
	"updated_time":      "UpdatedTime",
	"user_name":         "UserName",
}

var accesskeyCcToTfNameMap = func() map[string]string {
	m := make(map[string]string, len(accesskeyTfToCcNameMap))
	for tf, cc := range accesskeyTfToCcNameMap {
		m[cc] = tf
	}
	return m
}()

func init() {
	customresources.Register(accesskeyTerraformTypeName, func(_ context.Context, inner resource.Resource) (resource.Resource, error) {
		return &accesskeyResourceWithSecret{Resource: inner}, nil
	})
}

// accesskeyResourceWithSecret embeds the auto-generated inner resource so that
// Metadata/Schema/Update/Delete (declared on resource.Resource) forward
// automatically. Configure/Read/Create are overridden; ImportState and
// ConfigValidators are explicitly forwarded because they belong to optional
// interfaces the framework probes via type assertion on the outer type.
type accesskeyResourceWithSecret struct {
	resource.Resource
	provider tfcloudcontrol.Provider
}

func (r *accesskeyResourceWithSecret) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if c, ok := r.Resource.(resource.ResourceWithConfigure); ok {
		c.Configure(ctx, req, resp)
	}
	if v := req.ProviderData; v != nil {
		if p, ok := v.(tfcloudcontrol.Provider); ok {
			r.provider = p
		}
	}
}

// Read delegates to the inner generic Read, then restores secret_access_key
// from prior state. The Volcengine IAM API's GetResource never returns
// SecretAccessKey, so without this post-processing the value would be
// clobbered to null on every refresh.
func (r *accesskeyResourceWithSecret) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var priorSecret types.String
	priorDiags := req.State.GetAttribute(ctx, path.Root(accesskeySecretTerraformField), &priorSecret)

	r.Resource.Read(ctx, req, resp)
	if resp.Diagnostics.HasError() {
		return
	}

	if !priorDiags.HasError() && !priorSecret.IsNull() && !priorSecret.IsUnknown() && priorSecret.ValueString() != "" {
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root(accesskeySecretTerraformField), priorSecret)...)
	}
}

func (r *accesskeyResourceWithSecret) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	if i, ok := r.Resource.(resource.ResourceWithImportState); ok {
		i.ImportState(ctx, req, resp)
	}
}

func (r *accesskeyResourceWithSecret) ConfigValidators(ctx context.Context) []resource.ConfigValidator {
	if c, ok := r.Resource.(resource.ResourceWithConfigValidators); ok {
		return c.ConfigValidators(ctx)
	}
	return nil
}

// Create reproduces the generic Create flow but captures the CreateResource
// event's ResourceModel so it can extract SecretAccessKey, which the Volcengine
// IAM API only returns at creation time.
func (r *accesskeyResourceWithSecret) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	cloudControlClient := r.provider.CloudControlAPIClient(ctx)

	desiredState, err := generic.ToCloudControlString(ctx, request.Plan.Schema, request.Plan.Raw, accesskeyTfToCcNameMap)
	if err != nil {
		response.Diagnostics.Append(generic.DesiredStateErrorDiag("Plan", err))
		return
	}

	targetState := make(map[string]any)
	if err := json.Unmarshal([]byte(desiredState), &targetState); err != nil {
		response.Diagnostics.Append(generic.DesiredStateErrorDiag("Plan", err))
		return
	}

	output, err := cloudControlClient.CreateResourceWithContext(ctx, &cloudcontrol.CreateResourceInput{
		TypeName:    util.StringPtr(accesskeyCloudControlTypeName),
		RegionID:    r.provider.Region(ctx),
		ClientToken: util.StringPtr(util.GenerateToken(32)),
		TargetState: &targetState,
	})
	if err != nil {
		response.Diagnostics.Append(generic.ServiceOperationErrorDiag("Cloud Control API", "CreateResource", err))
		return
	}
	if output == nil || output.OperationStatus == nil {
		response.Diagnostics.Append(generic.ServiceOperationEmptyResultDiag("Cloud Control API", "CreateResource"))
		return
	}

	var event *cloudcontrol.ProgressEvent
	switch *output.OperationStatus {
	case base.SUCCESS:
		// Synchronous success: the create response itself carries the ProgressEvent.
		e := output.ProgressEvent
		event = &e
	case base.IN_PROGRESS, base.PENDING:
		taskId := ""
		if output.TaskID != nil {
			taskId = *output.TaskID
		}
		tflog.Info(ctx, "Cloud Control API CreateResource waiting task ......  ", map[string]interface{}{
			"TaskID":    hclog.Fmt("%v", taskId),
			"RequestID": hclog.Fmt("%v", output.GetRequestId()),
		})
		event, _, err = tfcloudcontrol.AwaitTask(ctx, cloudControlClient, taskId)
		if err != nil {
			response.Diagnostics.Append(generic.ServiceOperationErrorDiag("Cloud Control API Failed", "GetTask", err))
			return
		}
	case base.FAILED:
		response.Diagnostics.Append(generic.ServiceOperationErrorDiag("Cloud Control API Failed", "CreateResource",
			fmt.Errorf("invoke create handler failed status,resp:%s ", util.JsonString(output))))
		return
	default:
		response.Diagnostics.Append(generic.ServiceOperationErrorDiag("Cloud Control API Failed", "CreateResource",
			fmt.Errorf("invoke create handler other status,resp:%s ", util.JsonString(output))))
		return
	}

	if event == nil || event.Identifier == nil {
		response.Diagnostics.Append(generic.ServiceOperationEmptyResultDiag("Cloud Control API", "CreateResource"))
		return
	}
	id := *event.Identifier

	response.State.Raw = request.Plan.Raw
	if diags := response.State.SetAttribute(ctx, path.Root("id"), id); diags.HasError() {
		response.Diagnostics.Append(diags...)
		return
	}

	// Fill all unknown attributes (computed timestamps, AccessKeyId, and crucially
	// SecretAccessKey) from a single resource-model JSON. Prefer the CreateResource
	// event's ResourceModel because it's the only response that ever carries
	// SecretAccessKey; fall back to GetResource if the event didn't include a model.
	unknowns, err := generic.UnknownValuePaths(ctx, response.State.Raw)
	if err != nil {
		response.Diagnostics.AddError("Creation Of Terraform State Unsuccessful", err.Error())
		return
	}
	if len(unknowns) > 0 {
		var resourceModel string
		if event.ResourceModel != nil && *event.ResourceModel != "" {
			resourceModel = *event.ResourceModel
		} else {
			description, err := tfcloudcontrol.FindResourceByTypeNameAndID(ctx, cloudControlClient, r.provider.Region(ctx), accesskeyCloudControlTypeName, id)
			if err != nil {
				response.Diagnostics.Append(generic.ServiceOperationErrorDiag("Cloud Control API", "GetResource", err))
				return
			}
			if description == nil {
				response.Diagnostics.Append(generic.ServiceOperationEmptyResultDiag("Cloud Control API", "GetResource"))
				return
			}
			resourceModel = util.ToString(description.ResourceDescription.Properties)
		}
		if err := generic.SetUnknownValuesFromResourceModel(ctx, &response.State, unknowns, resourceModel, accesskeyCcToTfNameMap); err != nil {
			response.Diagnostics.AddError("Creation Of Terraform State Unsuccessful", err.Error())
			return
		}
	}
}
