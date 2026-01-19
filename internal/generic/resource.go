// Copyright (c) HashiCorp, Inc.
// Copyright (c) 2025 Beijing Volcano Engine Technology Co., Ltd.
// SPDX-License-Identifier: MPL-2.0
// This file has been modified by Beijing Volcano Engine Technology Co., Ltd. on 2025

package generic

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/mattbaird/jsonpatch"
	"github.com/volcengine/terraform-provider-volcenginecc/internal/base"
	"github.com/volcengine/terraform-provider-volcenginecc/internal/cloudcontrol"
	ccdiag "github.com/volcengine/terraform-provider-volcenginecc/internal/errs/diag"
	tfcloudcontrol "github.com/volcengine/terraform-provider-volcenginecc/internal/service/cloudcontrol"
	"github.com/volcengine/terraform-provider-volcenginecc/internal/tfresource"
	"github.com/volcengine/terraform-provider-volcenginecc/internal/util"
)

// ResourceOptionsFunc is a type alias for a resource type functional option.
type ResourceOptionsFunc func(*genericResource) error

// resourceWithAttributeNameMap is a helper function to construct functional options
// that set a resource type's attribute name maps.
// If multiple resourceWithAttributeNameMap calls are made, the last call overrides
// the previous calls' values.
func resourceWithAttributeNameMap(v map[string]string) ResourceOptionsFunc {
	return func(o *genericResource) error {
		if _, ok := v["id"]; !ok {
			// Synthesize a mapping for the reserved top-level "id" attribute.
			v["id"] = "ID"
		}

		ccToTfNameMapping := make(map[string]string, len(v))

		for tfName, ccName := range v {
			_, ok := ccToTfNameMapping[ccName]
			if ok {
				return fmt.Errorf("duplicate attribute name mapping for CloudControl property %s", ccName)
			}
			ccToTfNameMapping[ccName] = tfName
		}

		o.tfToCcNameMap = v
		o.ccToTfNameMap = ccToTfNameMapping

		return nil
	}
}

// resourceWithCloudControlTypeName is a helper function to construct functional options
// that set a resource type's Cloud Control type name.
// If multiple resourceWithCloudControlTypeName calls are made, the last call overrides
// the previous calls' values.
func resourceWithCloudControlTypeName(v string) ResourceOptionsFunc {
	return func(o *genericResource) error {
		o.ccTypeName = v

		return nil
	}
}

// resourceWithTerraformSchema is a helper function to construct functional options
// that set a resource type's Terraform schema.
// If multiple resourceWithTerraformSchema calls are made, the last call overrides
// the previous calls' values.
func resourceWithTerraformSchema(v schema.Schema) ResourceOptionsFunc {
	return func(o *genericResource) error {
		o.tfSchema = v

		return nil
	}
}

// resourceWithTerraformTypeName is a helper function to construct functional options
// that set a resource type's Terraform type name.
// If multiple resourceWithTerraformTypeName calls are made, the last call overrides
// the previous calls' values.
func resourceWithTerraformTypeName(v string) ResourceOptionsFunc {
	return func(o *genericResource) error {
		o.tfTypeName = v

		return nil
	}
}

// resourceIsImmutableType is a helper function to construct functional options
// that set a resource type's immutability flag.
// If multiple resourceIsImmutableType calls are made, the last call overrides
// the previous calls' values.
func resourceIsImmutableType(v bool) ResourceOptionsFunc {
	return func(o *genericResource) error {
		o.isImmutableType = v

		return nil
	}
}

// resourceWithWriteOnlyPropertyPaths is a helper function to construct functional options
// that set a resource type's write-only property paths (JSON Pointer).
// If multiple resourceWithWriteOnlyPropertyPaths calls are made, the last call overrides
// the previous calls' values.
func resourceWithWriteOnlyPropertyPaths(v []string) ResourceOptionsFunc {
	return func(o *genericResource) error {
		writeOnlyAttributePaths := make([]*path.Path, 0)

		for _, writeOnlyPropertyPath := range v {
			writeOnlyPropertyPath = strings.ReplaceAll(writeOnlyPropertyPath, "/*/", "/")
			writeOnlyPropertyPath = strings.TrimSuffix(writeOnlyPropertyPath, "/*")
			writeOnlyAttributePath, err := o.propertyPathToAttributePath(writeOnlyPropertyPath)

			if err != nil {
				// return fmt.Errorf("creating write-only attribute path (%s): %w", writeOnlyPropertyPath, err)
				continue
			}

			writeOnlyAttributePaths = append(writeOnlyAttributePaths, writeOnlyAttributePath)
		}

		o.writeOnlyAttributePaths = writeOnlyAttributePaths

		return nil
	}
}

// resourceWithReadOnlyPropertyPaths is a helper function to construct functional options
// that set a resource type's read-only property paths (JSON Pointer).
// If multiple resourceWithWriteOnlyPropertyPaths calls are made, the last call overrides
// the previous calls' values.
func resourceWithReadOnlyPropertyPaths(v []string) ResourceOptionsFunc {
	return func(o *genericResource) error {
		readOnlyAttributePaths := make([]*path.Path, 0)

		for _, readOnlyPath := range v {
			readOnlyPath = strings.ReplaceAll(readOnlyPath, "/*/", "/")
			readOnlyPath = strings.TrimSuffix(readOnlyPath, "/*")

			readOnlyAttributePath, err := o.propertyPathToAttributePath(readOnlyPath)

			if err != nil {
				// return fmt.Errorf("creating write-only attribute path (%s): %w", writeOnlyPropertyPath, err)
				continue
			}

			readOnlyAttributePaths = append(readOnlyAttributePaths, readOnlyAttributePath)
		}

		o.readOnlyAttributePaths = readOnlyAttributePaths

		return nil
	}
}

// resourceWithCreateOnlyPropertyPaths is a helper function to construct functional options
// that set a resource type's create-only property paths (JSON Pointer).
// If multiple resourceWithWriteOnlyPropertyPaths calls are made, the last call overrides
// the previous calls' values.
func resourceWithCreateOnlyPropertyPaths(v []string) ResourceOptionsFunc {
	return func(o *genericResource) error {
		createOnlyAttributePaths := make([]*path.Path, 0)

		for _, createOnlyPath := range v {
			createOnlyPath = strings.ReplaceAll(createOnlyPath, "/*/", "/")
			createOnlyPath = strings.TrimSuffix(createOnlyPath, "/*")

			createOnlyAttributePath, err := o.propertyPathToAttributePath(createOnlyPath)

			if err != nil {
				// return fmt.Errorf("creating write-only attribute path (%s): %w", writeOnlyPropertyPath, err)
				continue
			}

			createOnlyAttributePaths = append(createOnlyAttributePaths, createOnlyAttributePath)
		}

		o.createOnlyAttributePaths = createOnlyAttributePaths

		return nil
	}
}

const (
	resourceMaxWaitTimeCreate = 120 * time.Minute
	resourceMaxWaitTimeUpdate = 120 * time.Minute
	resourceMaxWaitTimeDelete = 120 * time.Minute
)

// resourceWithCreateTimeoutInMinutes is a helper function to construct functional options
// that set a resource type's create timeout (in minutes).
// If multiple resourceWithCreateTimeoutInMinutes calls are made, the last call overrides
// the previous calls' values.
func resourceWithCreateTimeoutInMinutes(v int) ResourceOptionsFunc {
	return func(o *genericResource) error {
		if v > 0 {
			o.createTimeout = time.Duration(v) * time.Minute
		} else {
			o.createTimeout = resourceMaxWaitTimeCreate
		}

		return nil
	}
}

// resourceWithUpdateTimeoutInMinutes is a helper function to construct functional options
// that set a resource type's update timeout (in minutes).
// If multiple resourceWithUpdateTimeoutInMinutes calls are made, the last call overrides
// the previous calls' values.
func resourceWithUpdateTimeoutInMinutes(v int) ResourceOptionsFunc {
	return func(o *genericResource) error {
		if v > 0 {
			o.updateTimeout = time.Duration(v) * time.Minute
		} else {
			o.updateTimeout = resourceMaxWaitTimeUpdate
		}

		return nil
	}
}

// resourceWithDeleteTimeoutInMinutes is a helper function to construct functional options
// that set a resource type's delete timeout (in minutes).
// If multiple resourceWithDeleteTimeoutInMinutes calls are made, the last call overrides
// the previous calls' values.
func resourceWithDeleteTimeoutInMinutes(v int) ResourceOptionsFunc {
	return func(o *genericResource) error {
		if v > 0 {
			o.deleteTimeout = time.Duration(v) * time.Minute
		} else {
			o.deleteTimeout = resourceMaxWaitTimeDelete
		}

		return nil
	}
}

// resourceWithConfigValidators is a helper function to construct functional options
// that set a resource type's config validators.
// If multiple resourceWithConfigValidators calls are made, the last call overrides
// the previous calls' values.
func resourceWithConfigValidators(vs ...resource.ConfigValidator) ResourceOptionsFunc {
	return func(o *genericResource) error {
		o.configValidators = vs

		return nil
	}
}

// ResourceOptions is a type alias for a slice of resource type functional options.
type ResourceOptions []ResourceOptionsFunc

// WithAttributeNameMap is a helper function to construct functional options
// that set a resource type's attribute name map, append that function to the
// current slice of functional options and return the new slice of options.
// It is intended to be chained with other similar helper functions in a builder pattern.
func (opts ResourceOptions) WithAttributeNameMap(v map[string]string) ResourceOptions {
	return append(opts, resourceWithAttributeNameMap(v))
}

// WithCloudControlTypeName is a helper function to construct functional options
// that set a resource type's Cloud Control type name, append that function to the
// current slice of functional options and return the new slice of options.
// It is intended to be chained with other similar helper functions in a builder pattern.
func (opts ResourceOptions) WithCloudControlTypeName(v string) ResourceOptions {
	return append(opts, resourceWithCloudControlTypeName(v))
}

// WithTerraformSchema is a helper function to construct functional options
// that set a resource type's Terraform schema, append that function to the
// current slice of functional options and return the new slice of options.
// It is intended to be chained with other similar helper functions in a builder pattern.
func (opts ResourceOptions) WithTerraformSchema(v schema.Schema) ResourceOptions {
	return append(opts, resourceWithTerraformSchema(v))
}

// WithTerraformTypeName is a helper function to construct functional options
// that set a resource type's Terraform type name, append that function to the
// current slice of functional options and return the new slice of options.
// It is intended to be chained with other similar helper functions in a builder pattern.
func (opts ResourceOptions) WithTerraformTypeName(v string) ResourceOptions {
	return append(opts, resourceWithTerraformTypeName(v))
}

// IsImmutableType is a helper function to construct functional options
// that set a resource type's Terraform immutability flag, append that function to the
// current slice of functional options and return the new slice of options.
// It is intended to be chained with other similar helper functions in a builder pattern.
func (opts ResourceOptions) IsImmutableType(v bool) ResourceOptions {
	return append(opts, resourceIsImmutableType(v))
}

// WithWriteOnlyPropertyPaths is a helper function to construct functional options
// that set a resource type's write-only property paths, append that function to the
// current slice of functional options and return the new slice of options.
// It is intended to be chained with other similar helper functions in a builder pattern.
func (opts ResourceOptions) WithWriteOnlyPropertyPaths(v []string) ResourceOptions {
	return append(opts, resourceWithWriteOnlyPropertyPaths(v))
}

// WithReadOnlyPropertyPaths is a helper function to construct functional options
// that set a resource type's read-only property paths, append that function to the
// current slice of functional options and return the new slice of options.
// It is intended to be chained with other similar helper functions in a builder pattern.
func (opts ResourceOptions) WithReadOnlyPropertyPaths(v []string) ResourceOptions {
	return append(opts, resourceWithReadOnlyPropertyPaths(v))
}

// WithCreateOnlyPropertyPaths is a helper function to construct functional options
// that set a resource type's read-only property paths, append that function to the
// current slice of functional options and return the new slice of options.
// It is intended to be chained with other similar helper functions in a builder pattern.
func (opts ResourceOptions) WithCreateOnlyPropertyPaths(v []string) ResourceOptions {
	return append(opts, resourceWithCreateOnlyPropertyPaths(v))
}

// WithCreateTimeoutInMinutes is a helper function to construct functional options
// that set a resource type's create timeout, append that function to the
// current slice of functional options and return the new slice of options.
// It is intended to be chained with other similar helper functions in a builder pattern.
func (opts ResourceOptions) WithCreateTimeoutInMinutes(v int) ResourceOptions {
	return append(opts, resourceWithCreateTimeoutInMinutes(v))
}

// WithUpdateTimeoutInMinutes is a helper function to construct functional options
// that set a resource type's update timeout, append that function to the
// current slice of functional options and return the new slice of options.
// It is intended to be chained with other similar helper functions in a builder pattern.
func (opts ResourceOptions) WithUpdateTimeoutInMinutes(v int) ResourceOptions {
	return append(opts, resourceWithUpdateTimeoutInMinutes(v))
}

// WithDeleteTimeoutInMinutes is a helper function to construct functional options
// that set a resource type's delete timeout, append that function to the
// current slice of functional options and return the new slice of options.
// It is intended to be chained with other similar helper functions in a builder pattern.
func (opts ResourceOptions) WithDeleteTimeoutInMinutes(v int) ResourceOptions {
	return append(opts, resourceWithDeleteTimeoutInMinutes(v))
}

// WithConfigValidators is a helper function to construct functional options
// that set a resource type's config validators, append that function to the
// current slice of functional options and return the new slice of options.
// It is intended to be chained with other similar helper functions in a builder pattern.
func (opts ResourceOptions) WithConfigValidators(vs ...resource.ConfigValidator) ResourceOptions {
	return append(opts, resourceWithConfigValidators(vs...))
}

// NewResource returns a new Resource from the specified varidaic list of functional options.
// It's public as it's called from generated code.
func NewResource(_ context.Context, optFns ...ResourceOptionsFunc) (resource.Resource, error) {
	v := &genericResource{}

	for _, optFn := range optFns {
		err := optFn(v)

		if err != nil {
			return nil, err
		}
	}

	if v.ccTypeName == "" {
		return nil, fmt.Errorf("no Cloud Control type name specified")
	}
	if v.tfTypeName == "" {
		return nil, fmt.Errorf("no Terraform type name specified")
	}

	return v, nil
}

// Implements resource.Resource.
type genericResource struct {
	ccTypeName               string            // Cloud Control type name for the resource type
	tfSchema                 schema.Schema     // Terraform schema for the resource type
	tfTypeName               string            // Terraform type name for resource type
	tfToCcNameMap            map[string]string // Map of Terraform attribute name to Cloud Control property name
	ccToTfNameMap            map[string]string // Map of Cloud Control property name to Terraform attribute name
	isImmutableType          bool              // Resources cannot be updated and must be recreated
	writeOnlyAttributePaths  []*path.Path      // Paths to any write-only attributes
	readOnlyAttributePaths   []*path.Path      // Paths to any read-only attributes
	createOnlyAttributePaths []*path.Path      // Paths to any create-only attributes

	createTimeout    time.Duration              // Maximum wait time for resource creation
	updateTimeout    time.Duration              // Maximum wait time for resource update
	deleteTimeout    time.Duration              // Maximum wait time for resource deletion
	configValidators []resource.ConfigValidator // Required attributes validators
	provider         tfcloudcontrol.Provider
}

var (
	// Path to the "id" attribute which uniquely (for a specific resource type) identifies the resource.
	// This attribute is required for acceptance testing.
	idAttributePath = path.Root("id")
)

func (r *genericResource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = r.tfTypeName
}

func (r *genericResource) Schema(_ context.Context, request resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = r.tfSchema
}

func (r *genericResource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) { //nolint:unparam
	if v := request.ProviderData; v != nil {
		r.provider = v.(tfcloudcontrol.Provider)
	}
}

func (r *genericResource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	ctx = r.bootstrapContext(ctx)

	traceEntry(ctx, "Resource.Create")

	cloudControlClient := r.provider.CloudControlAPIClient(ctx)

	tflog.Debug(ctx, "Request.Plan.Raw", map[string]interface{}{
		"value": hclog.Fmt("%v", request.Plan.Raw),
	})

	translator := toCloudControl{tfToCfNameMap: r.tfToCcNameMap}
	desiredState, err := translator.AsString(ctx, request.Plan.Schema, request.Plan.Raw)

	if err != nil {
		response.Diagnostics.Append(DesiredStateErrorDiag("Plan", err))

		return
	}

	tflog.Debug(ctx, "Cloud Control DesiredState", map[string]interface{}{
		"value": desiredState,
	})

	targetState := make(map[string]any)
	err = json.Unmarshal([]byte(desiredState), &targetState)
	if err != nil {
		response.Diagnostics.Append(DesiredStateErrorDiag("Plan", err))
		return
	}
	output, err := cloudControlClient.CreateResourceWithContext(ctx, &cloudcontrol.CreateResourceInput{
		TypeName:    util.StringPtr(r.ccTypeName),
		RegionID:    r.provider.Region(ctx),
		ClientToken: util.StringPtr(util.GenerateToken(24)),
		TargetState: &targetState,
	})

	if err != nil {
		response.Diagnostics.Append(ServiceOperationErrorDiag("Cloud Control API", "CreateResource", err))
		return
	}

	if output == nil || output.OperationStatus == nil {
		response.Diagnostics.Append(ServiceOperationEmptyResultDiag("Cloud Control API", "CreateResource"))
		return
	}

	taskId := ""
	Success := false
	if *output.OperationStatus == base.FAILED {
		apiErr := fmt.Errorf("invoke create handler failed status,resp:%s,err:%v ", util.JsonString(output), nil)
		response.Diagnostics.Append(ServiceOperationErrorDiag("Cloud Control API Failed", "CreateResource", apiErr))
		return
	} else if *output.OperationStatus == base.SUCCESS {
		Success = true
	} else if *output.OperationStatus == base.IN_PROGRESS || *output.OperationStatus == base.PENDING {
		taskId = *output.TaskID
	} else {
		apiErr := fmt.Errorf("invoke create handler other status,resp:%s,err:%v ", util.JsonString(output), nil)
		response.Diagnostics.Append(ServiceOperationErrorDiag("Cloud Control API Failed", "CreateResource", apiErr))
		return
	}
	tflog.Info(ctx, "Cloud Control API CreateResource waiting task ......  ", map[string]interface{}{
		"TaskID":    hclog.Fmt("%v", taskId),
		"RequestID": hclog.Fmt("%v", output.GetRequestId()),
	})
	var event *cloudcontrol.ProgressEvent
	if !Success {
		event, _, err = tfcloudcontrol.AwaitTask(ctx, cloudControlClient, taskId)
		if err != nil {
			response.Diagnostics.Append(ServiceOperationErrorDiag("Cloud Control API Failed", "GetTask", err))
			return
		}
	}
	id := *event.Identifier

	if id != "" {
		if err := r.setId(ctx, id, &response.State); err != nil {
			response.Diagnostics.Append(ResourceIdentifierNotSetDiag(err))
			return
		}
	}

	// Produce a wholly-known new State by determining the final values for any attributes left unknown in the planned state.
	response.State.Raw = request.Plan.Raw

	// Set the "id" attribute.
	if err = r.setId(ctx, id, &response.State); err != nil {
		response.Diagnostics.Append(ResourceIdentifierNotSetDiag(err))
		return
	}

	response.Diagnostics.Append(r.populateUnknownValues(ctx, id, &response.State)...)
	if response.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Response.State.Raw", map[string]interface{}{
		"value": hclog.Fmt("%v", response.State.Raw),
	})

	traceExit(ctx, "Resource.Create")
}

func (r *genericResource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {

	ctx = r.bootstrapContext(ctx)

	traceEntry(ctx, "Resource.Read")

	tflog.Debug(ctx, "Request.State.Raw", map[string]interface{}{
		"value": hclog.Fmt("%v", request.State.Raw),
	})

	client := r.provider.CloudControlAPIClient(ctx)

	currentState := &request.State
	id, err := r.getId(ctx, currentState)

	if err != nil {
		response.Diagnostics.Append(ResourceIdentifierNotFoundDiag(err))

		return
	}

	description, err := r.describe(ctx, client, id)

	if tfresource.NotFound(err) {
		response.Diagnostics.Append(ResourceNotFoundWarningDiag(err))
		//response.State.RemoveResource(ctx)

		return
	}

	if err != nil {
		response.Diagnostics.Append(ServiceOperationErrorDiag("Cloud Control API", "GetResource", err))

		return
	}
	tflog.Debug(ctx, "Cloud Control API GetResource", map[string]interface{}{
		"value": util.ToString(description.ResourceDescription.Properties),
	})
	translator := toTerraform{cfToTfNameMap: r.ccToTfNameMap}
	schema := currentState.Schema
	val, err := translator.FromString(ctx, schema, util.ToString(description.ResourceDescription.Properties))

	if err != nil {
		response.Diagnostics.AddError(
			"Creation Of Terraform State Unsuccessful",
			fmt.Sprintf("Unable to create a Terraform State value from Cloud Control API Properties. This is typically an error with the Terraform provider implementation. Original Error: %s", err.Error()),
		)

		return
	}

	response.State = tfsdk.State{
		Schema: schema,
		Raw:    val,
	}

	// Copy over any write-only values.
	// They can only be in the current state.
	for _, path := range r.writeOnlyAttributePaths {
		response.Diagnostics.Append(copyStateValueAtPath(ctx, &response.State, &request.State, *path)...)
		if response.Diagnostics.HasError() {
			return
		}
	}

	// Set the "id" attribute.
	if err := r.setId(ctx, id, &response.State); err != nil {
		response.Diagnostics.Append(ResourceIdentifierNotSetDiag(err))

		return
	}

	tflog.Debug(ctx, "Response.State.Raw", map[string]interface{}{
		"value": hclog.Fmt("%v", response.State.Raw),
	})

	traceExit(ctx, "Resource.Read")
}

func (r *genericResource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	ctx = r.bootstrapContext(ctx)

	traceEntry(ctx, "Resource.Update")

	cloudControlClient := r.provider.CloudControlAPIClient(ctx)

	currentState := &request.State
	id, err := r.getId(ctx, currentState)

	if err != nil {
		response.Diagnostics.Append(ResourceIdentifierNotFoundDiag(err))

		return
	}

	// Clear any write-only values.
	// This forces patch document generation to always add values.
	currentStateRaw := currentState.Raw
	planRaw := request.Plan.Raw
	copyPlan := planRaw.Copy()

	if len(r.writeOnlyAttributePaths) > 0 {
		planRaw, err = tftypes.Transform(planRaw, func(tfPath *tftypes.AttributePath, val tftypes.Value) (tftypes.Value, error) {
			if len(tfPath.Steps()) < 1 {
				return val, nil
			}

			path, diags := attributePath(ctx, tfPath, request.Plan.Schema)
			if diags.HasError() {
				return val, ccdiag.DiagnosticsError(diags)
			}

			for _, woPath := range r.writeOnlyAttributePaths {
				if woPath.Equal(path) {
					//if write only value not change,set the val nil
					var currentValue attr.Value
					getCurrentDiags := currentState.GetAttribute(context.Background(), path, &currentValue)
					if !getCurrentDiags.HasError() {
						tfCurVal, tfErr := currentValue.ToTerraformValue(context.Background())
						if tfErr == nil {
							if val.Equal(tfCurVal) {
								return tftypes.NewValue(val.Type(), nil), nil
							}
						}
					}
					createOnly := false
					for _, createOnlyPath := range r.createOnlyAttributePaths {
						if createOnlyPath.Equal(path) {
							createOnly = true
							break
						}
					}
					if createOnly {
						return tftypes.NewValue(val.Type(), nil), nil
					} else {
						return val, nil
					}
				}

			}
			return val, nil
		})
		if err != nil {
			response.Diagnostics.Append(DesiredStateErrorDiag("Prior State", err))
			return
		}
	}
	translator := toCloudControl{tfToCfNameMap: r.tfToCcNameMap}
	currentDesiredState, err := translator.AsString(ctx, currentState.Schema, currentStateRaw)

	if err != nil {
		response.Diagnostics.Append(DesiredStateErrorDiag("Prior State", err))

		return
	}

	plannedDesiredState, err := translator.AsString(ctx, request.Plan.Schema, planRaw)

	if err != nil {
		response.Diagnostics.Append(DesiredStateErrorDiag("Plan", err))
		return
	}

	//for update set ,get new resource
	description, err := r.describeWithSysTag(ctx, r.provider.CloudControlAPIClient(ctx), id)

	if tfresource.NotFound(err) {
		response.Diagnostics.Append(ResourceNotFoundAfterWriteDiag(err))
		return
	}
	if err != nil {
		response.Diagnostics.Append(ServiceOperationErrorDiag("Cloud Control API", "GetResource", err))
		return
	}

	if description == nil {
		response.Diagnostics.Append(ServiceOperationEmptyResultDiag("Cloud Control API", "GetResource"))
		return
	}
	tflog.Debug(ctx, "Cloud Control API GetResource", map[string]interface{}{
		"value": util.ToString(description.ResourceDescription.Properties),
	})
	currentDesiredState = util.ToString(description.ResourceDescription.Properties)

	currentDesiredState, err = translateForUpdate(currentDesiredState, id, r.tfSchema.Attributes, r.ccToTfNameMap)
	plannedDesiredState, err = translateForUpdate(plannedDesiredState, id, r.tfSchema.Attributes, r.ccToTfNameMap)
	//plannedDesiredState中补全sys: volc: tag处理，保证patchDocument与CCAPI计算一致
	plannedDesiredState, err = MergeSystemTags(currentDesiredState, plannedDesiredState)
	if err != nil {
		response.Diagnostics.AddError(
			"Creation Of JSON Patch Unsuccessful",
			fmt.Sprintf("Unable to create a JSON Patch for resource update. This is typically an error with the Terraform provider implementation. Original Error: %s", err.Error()),
		)

		return
	}
	patchDocument, err := patchDocument(currentDesiredState, plannedDesiredState)

	if err != nil {
		response.Diagnostics.AddError(
			"Creation Of JSON Patch Unsuccessful",
			fmt.Sprintf("Unable to create a JSON Patch for resource update. This is typically an error with the Terraform provider implementation. Original Error: %s", err.Error()),
		)

		return
	}

	tflog.Debug(ctx, "Cloud Control API PatchDocument", map[string]interface{}{
		"value": patchDocument,
	})
	PatchDocumentArray := make([]any, 0)
	err = json.Unmarshal([]byte(patchDocument), &PatchDocumentArray)
	if err != nil {
		response.Diagnostics.Append(DesiredStateErrorDiag("Plan", err))
		return
	}
	output, err := cloudControlClient.UpdateResourceWithContext(ctx, &cloudcontrol.UpdateResourceInput{
		TypeName:      util.StringPtr(r.ccTypeName),
		RegionID:      util.StringPtr(r.provider.Region(ctx)),
		Identifier:    util.StringPtr(id),
		ClientToken:   util.StringPtr(util.GenerateToken(24)),
		PatchDocument: PatchDocumentArray,
	})
	if err != nil {
		apiErr := fmt.Errorf("invoke create handler failed status,resp:%s,err:%v ", util.JsonString(output), nil)
		response.Diagnostics.Append(ServiceOperationErrorDiag("Cloud Control API", "UpdateResource", apiErr))
		return
	}
	if output == nil || output.OperationStatus == nil {
		apiErr := fmt.Errorf("invoke create handler failed,resp:%s,err:%v ", util.JsonString(output), err)
		response.Diagnostics.Append(ServiceOperationErrorDiag("Cloud Control API", "UpdateResource", apiErr))
		return
	}
	taskId := ""
	Success := false
	status := *output.OperationStatus
	if status == base.FAILED {
		apiErr := fmt.Errorf("invoke update handler failed,resp:%s,err:%v ", util.JsonString(output), err)
		response.Diagnostics.Append(ServiceOperationErrorDiag("Cloud Control API", "UpdateResource", apiErr))
		return
	} else if status == base.SUCCESS {
		Success = true
	} else if status == base.IN_PROGRESS || status == base.PENDING {
		taskId = *output.TaskID
	} else {
		apiErr := fmt.Errorf("invoke update handler others status,resp:%s,err:%v ", util.JsonString(output), err)
		response.Diagnostics.Append(ServiceOperationErrorDiag("Cloud Control API", "UpdateResource", apiErr))
		return
	}
	tflog.Info(ctx, "Cloud Control API UpdateResource waiting task ...... ", map[string]interface{}{
		"TaskID":    hclog.Fmt("%v", taskId),
		"RequestID": hclog.Fmt("%v", output.GetRequestId()),
	})
	if !Success {
		_, _, err = tfcloudcontrol.AwaitTask(ctx, cloudControlClient, taskId)
		if err != nil {
			response.Diagnostics.Append(ServiceOperationWaiterErrorDiag("Cloud Control API", "UpdateResource", err))
			return
		}
	}
	//
	//if len(r.writeOnlyAttributePaths) > 0 {
	//	planRaw, err = tftypes.Transform(planRaw, func(tfPath *tftypes.AttributePath, val tftypes.Value) (tftypes.Value, error) {
	//		if len(tfPath.Steps()) < 1 {
	//			return val, nil
	//		}
	//
	//		path, diags := attributePath(ctx, tfPath, request.Plan.Schema)
	//		if diags.HasError() {
	//			return val, ccdiag.DiagnosticsError(diags)
	//		}
	//
	//		for _, woPath := range r.writeOnlyAttributePaths {
	//			if woPath.Equal(path) {
	//				return tftypes.NewValue(val.Type(), nil), nil
	//			}
	//		}
	//		return val, nil
	//	})
	//	if err != nil {
	//		response.Diagnostics.Append(DesiredStateErrorDiag("Prior State", err))
	//		return
	//	}
	//}

	// Produce a wholly-known new State by determining the final values for any attributes left unknown in the planned state.
	response.State.Raw = copyPlan

	response.Diagnostics.Append(r.populateUnknownValues(ctx, id, &response.State)...)
	if response.Diagnostics.HasError() {
		return
	}

	traceExit(ctx, "Resource.Update")
}

func (r *genericResource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	ctx = r.bootstrapContext(ctx)

	traceEntry(ctx, "Resource.Delete")

	conn := r.provider.CloudControlAPIClient(ctx)

	id, err := r.getId(ctx, &request.State)

	if err != nil {
		response.Diagnostics.Append(ResourceIdentifierNotFoundDiag(err))

		return
	}

	err = tfcloudcontrol.DeleteResource(ctx, conn, r.provider.Region(ctx), "", r.ccTypeName, id)

	if err != nil {
		response.Diagnostics.Append(ServiceOperationErrorDiag("Cloud Control API", "DeleteResource", err))

		return
	}

	response.State.RemoveResource(ctx)

	traceExit(ctx, "Resource.Delete")
}

func (r *genericResource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	ctx = r.bootstrapContext(ctx)

	traceEntry(ctx, "Resource.ImportState")

	tflog.Debug(ctx, "Request.ID", map[string]interface{}{
		"value": hclog.Fmt("%v", request.ID),
	})

	resource.ImportStatePassthroughID(ctx, idAttributePath, request, response)

	traceExit(ctx, "Resource.ImportState")
}

// ConfigValidators returns a list of functions which will all be performed during validation.
func (r *genericResource) ConfigValidators(context.Context) []resource.ConfigValidator {
	validators := make([]resource.ConfigValidator, 0)

	if len(r.configValidators) > 0 {
		validators = append(validators, r.configValidators...)
	}

	return validators
}

// describe returns the live state of the specified resource.
func (r *genericResource) describe(ctx context.Context, client *cloudcontrol.CloudControl, id string) (*cloudcontrol.GetResourceOutput, error) {
	return tfcloudcontrol.FindResourceByTypeNameAndID(ctx, client, r.provider.Region(ctx), r.ccTypeName, id)
}

// describe returns the live state of the specified resource.
func (r *genericResource) describeWithSysTag(ctx context.Context, client *cloudcontrol.CloudControl, id string) (*cloudcontrol.GetResourceOutput, error) {
	return tfcloudcontrol.FindResourceByTypeNameAndIDWithSysTag(ctx, client, r.provider.Region(ctx), r.ccTypeName, id)
}

// getId returns the resource's primary identifier value from State.
func (r *genericResource) getId(ctx context.Context, state *tfsdk.State) (string, error) {
	var val string
	diags := state.GetAttribute(ctx, idAttributePath, &val)

	if diags.HasError() {
		return "", ccdiag.DiagnosticsError(diags)
	}

	return val, nil
}

// setId sets the resource's primary identifier value in State.
func (r *genericResource) setId(ctx context.Context, val string, state *tfsdk.State) error {
	diags := state.SetAttribute(ctx, idAttributePath, val)

	if diags.HasError() {
		return ccdiag.DiagnosticsError(diags)
	}

	return nil
}

// populateUnknownValues populates and unknown values in State with values from the current resource description.
func (r *genericResource) populateUnknownValues(ctx context.Context, id string, state *tfsdk.State) diag.Diagnostics {
	var diags diag.Diagnostics

	unknowns, err := UnknownValuePaths(ctx, state.Raw)

	if err != nil {
		diags.AddError(
			"Creation Of Terraform State Unsuccessful",
			fmt.Sprintf("Unable to set Terraform State Unknown values from Cloud Control API Properties. This is typically an error with the Terraform provider implementation. Original Error: %s", err.Error()),
		)

		return diags
	}
	if len(unknowns) == 0 {
		return nil
	}

	description, err := r.describe(ctx, r.provider.CloudControlAPIClient(ctx), id)

	if tfresource.NotFound(err) {
		diags.Append(ResourceNotFoundAfterWriteDiag(err))

		return diags
	}

	if err != nil {
		diags.Append(ServiceOperationErrorDiag("Cloud Control API", "GetResource", err))

		return diags
	}

	if description == nil {
		diags.Append(ServiceOperationEmptyResultDiag("Cloud Control API", "GetResource"))

		return diags
	}

	err = SetUnknownValuesFromResourceModel(ctx, state, unknowns, util.ToString(description.ResourceDescription.Properties), r.ccToTfNameMap)

	if err != nil {
		diags.AddError(
			"Creation Of Terraform State Unsuccessful",
			fmt.Sprintf("Unable to set Terraform State Unknown values from Cloud Control API Properties. This is typically an error with the Terraform provider implementation. Original Error: %s", err.Error()),
		)

		return diags
	}

	//diags = SetReadOnlyFromResourceModel(ctx, state, r.readOnlyAttributePaths, util.ToString(description.ResourceDescription.Properties), r.ccToTfNameMap)
	return nil
}

// bootstrapContext injects the Cloud Control type name into logger contexts.
func (r *genericResource) bootstrapContext(ctx context.Context) context.Context {
	ctx = tflog.SetField(ctx, LoggingKeyCCType, r.ccTypeName)
	//ctx = r.provider.RegisterLogger(ctx)

	return ctx
}

func MergeSystemTags(
	currentDesiredState string,
	plannedDesiredState string,
) (string, error) {

	var current map[string]interface{}
	var planned map[string]interface{}

	if err := json.Unmarshal([]byte(currentDesiredState), &current); err != nil {
		return "", fmt.Errorf("unmarshal currentDesiredState failed: %w", err)
	}
	if err := json.Unmarshal([]byte(plannedDesiredState), &planned); err != nil {
		return "", fmt.Errorf("unmarshal plannedDesiredState failed: %w", err)
	}

	// 提取 Tags
	currentTags := extractTags(current)
	plannedTags := extractTags(planned)
	if currentTags == nil {
		return plannedDesiredState, nil
	}
	if plannedTags == nil {
		plannedTags = make([]map[string]interface{}, 0)
	}

	// 构建 planned 中已有 tag 的索引（Key → exists）
	existing := make(map[string]struct{})
	for _, t := range plannedTags {
		if key, ok := t["Key"].(string); ok {
			existing[key] = struct{}{}
		}
	}

	// 将 current 中的 sys:/volc: tag 合并进 planned
	for _, t := range currentTags {
		key, ok := t["Key"].(string)
		if !ok {
			continue
		}

		if strings.HasPrefix(key, "sys:") || strings.HasPrefix(key, "volc:") {
			if _, found := existing[key]; !found {
				plannedTags = append(plannedTags, t)
				existing[key] = struct{}{}
			}
		}
	}

	// 回写 Tags
	if len(plannedTags) > 0 {
		planned["Tags"] = plannedTags
	}

	merged, err := json.Marshal(planned)
	if err != nil {
		return "", fmt.Errorf("marshal merged plannedDesiredState failed: %w", err)
	}

	return string(merged), nil
}

func extractTags(obj map[string]interface{}) []map[string]interface{} {
	raw, ok := obj["Tags"]
	if !ok {
		return nil
	}

	arr, ok := raw.([]interface{})
	if !ok {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(arr))
	for _, item := range arr {
		if m, ok := item.(map[string]interface{}); ok {
			result = append(result, m)
		}
	}
	return result
}

// patchDocument returns a JSON Patch document describing the difference between `old` and `new`.
func patchDocument(old, new string) (string, error) {
	patch, err := jsonpatch.CreatePatch([]byte(old), []byte(new))

	if err != nil {
		return "", err
	}

	b, err := json.Marshal(patch)

	if err != nil {
		return "", err
	}

	return string(b), nil
}

// propertyPathToAttributePath returns the AttributePath for the specified JSON Pointer property path.
func (r *genericResource) propertyPathToAttributePath(propertyPath string) (*path.Path, error) {
	segments := strings.Split(propertyPath, "/")

	if got, expected := len(segments), 3; got < expected {
		return nil, fmt.Errorf("expected at least %d property path segments, got: %d", expected, got)
	}

	if got, expected := segments[0], ""; got != expected {
		return nil, fmt.Errorf("expected %q for the initial property path segment, got: %q", expected, got)
	}

	if got, expected := segments[1], "properties"; got != expected {
		return nil, fmt.Errorf("expected %q for the second property path segment, got: %q", expected, got)
	}

	attributePath := path.Empty()

	for _, segment := range segments[2:] {
		switch segment {
		case "", "*":
			return nil, fmt.Errorf("invalid property path segment: %q", segment)

		default:
			attributeName, ok := r.ccToTfNameMap[segment]
			if !ok {
				return nil, fmt.Errorf("attribute name mapping not found: %s", segment)
			}
			attributePath = attributePath.AtName(attributeName)
		}
	}

	return &attributePath, nil
}
