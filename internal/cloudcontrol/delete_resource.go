package cloudcontrol

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
)

const opDeleteResourceCommon = "DeleteResource"

func (c *CloudControl) DeleteResourceRequest(input *DeleteResourceInput) (req *request.Request, output *DeleteResourceOutput) {
	op := &request.Operation{
		Name:       opDeleteResourceCommon,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteResourceInput{}
	}

	output = &DeleteResourceOutput{}
	req = c.newRequest(op, input, output)
	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")
	return
}

func (c *CloudControl) DeleteResourceWithContext(ctx volcengine.Context, input *DeleteResourceInput, opts ...request.Option) (*DeleteResourceOutput, error) {
	req, out := c.DeleteResourceRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type (
	DeleteResourceInput struct {
		TypeName    *string
		Identifier  *string
		RegionID    *string
		ClientToken *string
	}

	DeleteResourceOutput struct {
		Metadata *response.ResponseMetadata

		ProgressEvent
	}
)

func (output *DeleteResourceOutput) GetRequestId() string {
	if output == nil || output.Metadata == nil {
		return "-"
	}
	reqId := output.Metadata.RequestId
	if reqId == "" {
		return "-"
	}
	return output.Metadata.RequestId
}
