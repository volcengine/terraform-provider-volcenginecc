package cloudcontrol

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
)

const opUpdateResourceCommon = "UpdateResource"

func (c *CloudControl) UppdateRequest(input *UpdateResourceInput) (req *request.Request, output *UpdateResourceOutput) {
	op := &request.Operation{
		Name:       opUpdateResourceCommon,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateResourceInput{}
	}

	output = &UpdateResourceOutput{}
	req = c.newRequest(op, input, output)
	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")
	return
}

func (c *CloudControl) UpdateResourceWithContext(ctx volcengine.Context, input *UpdateResourceInput, opts ...request.Option) (*UpdateResourceOutput, error) {
	req, out := c.UppdateRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type (
	UpdateResourceInput struct {
		TypeName      *string
		RegionID      *string
		Identifier    *string
		ClientToken   *string
		PatchDocument []any
	}

	UpdateResourceOutput struct {
		Metadata        *response.ResponseMetadata
		EventTime       *string
		TypeName        *string
		Operation       *string
		OperationStatus *string
		TaskID          *string
	}
)

func (output *UpdateResourceOutput) GetRequestId() string {
	if output == nil || output.Metadata == nil {
		return "-"
	}
	reqId := output.Metadata.RequestId
	if reqId == "" {
		return "-"
	}
	return output.Metadata.RequestId
}
