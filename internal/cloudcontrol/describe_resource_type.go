package cloudcontrol

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
)

const opDescribeResourceTypeCommon = "DescribeResourceType"

func (c *CloudControl) DescribeResourceTypeRequest(input *DescribeResourceTypeInput) (req *request.Request, output *DescribeResourceTypeOutput) {
	op := &request.Operation{
		Name:       opDescribeResourceTypeCommon,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeResourceTypeInput{}
	}

	output = &DescribeResourceTypeOutput{}
	req = c.newRequest(op, input, output)
	return
}

func (c *CloudControl) DescribeResourceTypeWithContext(ctx volcengine.Context, input *DescribeResourceTypeInput, opts ...request.Option) (*DescribeResourceTypeOutput, error) {
	req, out := c.DescribeResourceTypeRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type (
	DescribeResourceTypeInput struct {
		TypeName    *string
		Identifier  *string
		RegionID    *string
		ClientToken *string
	}

	DescribeResourceTypeOutput struct {
		Metadata *response.ResponseMetadata

		TypeName *string
		Schema   any
	}
)
