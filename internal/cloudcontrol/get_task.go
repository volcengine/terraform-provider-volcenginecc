package cloudcontrol

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
)

const opGetTaskCommon = "GetTask"

func (c *CloudControl) GetTaskRequest(input *GetTaskInput) (req *request.Request, output *GetTaskOutput) {
	op := &request.Operation{
		Name:       opGetTaskCommon,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetTaskInput{}
	}

	output = &GetTaskOutput{}
	req = c.newRequest(op, input, output)
	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")
	return
}

func (c *CloudControl) GetTaskWithContext(ctx volcengine.Context, input *GetTaskInput, opts ...request.Option) (*GetTaskOutput, error) {
	req, out := c.GetTaskRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type (
	GetTaskInput struct {
		TaskID *string
	}

	GetTaskOutput struct {
		ResponseMetadata *response.ResponseMetadata
		ProgressEvent
	}
)
