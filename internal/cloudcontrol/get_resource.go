package cloudcontrol

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
)

const opGetResourceCommon = "GetResource"

func (c *CloudControl) GetResourceRequest(input *GetResourceInput) (req *request.Request, output *GetResourceOutput) {
	op := &request.Operation{
		Name:       opGetResourceCommon,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetResourceInput{}
	}

	output = &GetResourceOutput{}
	req = c.newRequest(op, input, output)
	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")
	return
}

func (c *CloudControl) GetResourceWithContext(ctx volcengine.Context, input *GetResourceInput, opts ...request.Option) (*GetResourceOutput, error) {
	req, out := c.GetResourceRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type (
	GetResourceInput struct {
		TypeName   *string
		RegionID   *string
		Identifier *string
	}

	GetResourceOutput struct {
		ResponseMetadata *response.ResponseMetadata

		TypeName            *string
		ResourceDescription *ResourceDescriptionForGetResourceOutput
	}

	ResourceDescriptionForGetResourceOutput struct {
		Identifier *string
		Properties *string
	}
)
