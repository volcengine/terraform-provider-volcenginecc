// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudcontrol

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/volcengine/terraform-provider-volcenginecc/internal/cloudcontrol"
	"github.com/volcengine/terraform-provider-volcenginecc/internal/util"
)

func ListResourcesByTypeName(ctx context.Context, client *cloudcontrol.CloudControl, region, typeName string) (*cloudcontrol.ListResourceOutput, error) {
	tflog.Debug(ctx, "ListResourcesByTypeName", map[string]interface{}{
		"cfTypeName": typeName,
	})

	input := &cloudcontrol.ListResourceInput{
		TypeName:   util.StringPtr(typeName),
		RegionID:   util.StringPtr(region),
		MaxResults: 20,
	}
	var nextToken *string
	result := &cloudcontrol.ListResourceOutput{
		ResourceDescriptions: make([]*cloudcontrol.ResourceDescriptionForListResourceOutput, 0),
	}
	for {
		resp, err := client.ListResourceWithContext(ctx, input)
		if err != nil {
			return nil, fmt.Errorf("call GetResource failed,resp:%s,err:%v ", util.JsonString(resp), err)
		}
		result.ResourceDescriptions = append(result.ResourceDescriptions, resp.ResourceDescriptions...)
		if resp.NextToken == nil || *resp.NextToken == "" {
			break
		}
		nextToken = resp.NextToken
		input.NextToken = nextToken

	}
	return result, nil
}
