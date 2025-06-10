package cloudcontrol

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
)

const opListResourceCommon = "ListResources"

func (c *CloudControl) ListResourceRequest(input *ListResourceInput) (req *request.Request, output *ListResourceOutput) {
	op := &request.Operation{
		Name:       opListResourceCommon,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListResourceInput{}
	}

	output = &ListResourceOutput{}
	req = c.newRequest(op, input, output)
	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")
	return
}

func (c *CloudControl) ListResourceWithContext(ctx volcengine.Context, input *ListResourceInput, opts ...request.Option) (*ListResourceOutput, error) {
	req, out := c.ListResourceRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type (
	ListResourceInput struct {
		TypeName   *string
		RegionID   *string
		MaxResults int
		NextToken  *string
	}

	ListResourceOutput struct {
		ResponseMetadata     *response.ResponseMetadata
		TypeName             *string
		NextToken            *string
		ResourceDescriptions []*ResourceDescriptionForListResourceOutput
	}

	ResourceDescriptionForListResourceOutput struct {
		Identifier *string
		Properties *string
	}
)
