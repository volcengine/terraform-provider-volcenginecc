// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudcontrol

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/volcengine/terraform-provider-volcenginecc/internal/cloudcontrol"
	"github.com/volcengine/terraform-provider-volcenginecc/internal/tfresource"
)

func FindResourceByTypeNameAndID(ctx context.Context, client *cloudcontrol.CloudControl, regionId, typeName, id string) (*cloudcontrol.GetResourceOutput, error) {
	tflog.Debug(ctx, "FindResourceByTypeNameAndID", map[string]interface{}{
		"cfTypeName": typeName,
		"id":         id,
	})
	output, err := client.GetResourceWithContext(ctx, &cloudcontrol.GetResourceInput{
		TypeName:   &typeName,
		RegionID:   &regionId,
		Identifier: &id,
	})
	if err != nil {
		if strings.Contains(err.Error(), "HandlerErrorCode: NotFound") {
			return nil, &tfresource.NotFoundError{LastError: err}

		}
		return nil, err
	}
	if output == nil || output.ResourceDescription == nil {
		return nil, &tfresource.NotFoundError{Message: "Empty result"}
	}
	return output, nil
}
