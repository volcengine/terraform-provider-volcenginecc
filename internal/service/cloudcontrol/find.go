// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudcontrol

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/volcengine/terraform-provider-volcenginecc/internal/cloudcontrol"
	"github.com/volcengine/terraform-provider-volcenginecc/internal/tfresource"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
)

func FindResourceByTypeNameAndIDWithSysTag(ctx context.Context, client *cloudcontrol.CloudControl, regionId, typeName, id string) (*cloudcontrol.GetResourceOutput, error) {
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
	err = NormalizeResourceDescription(output.ResourceDescription)
	if err != nil {
		return nil, &tfresource.NotFoundError{Message: "normalize resource description err", LastError: err}
	}
	return output, nil
}

func NormalizeResourceDescription(desc *cloudcontrol.ResourceDescriptionForGetResourceOutput) error {
	if desc == nil || desc.Properties == nil || *desc.Properties == "" {
		return nil
	}

	var props map[string]interface{}
	if err := json.Unmarshal([]byte(*desc.Properties), &props); err != nil {
		return fmt.Errorf("failed to unmarshal ResourceDescription.Properties: %w", err)
	}
	// 如果有 Tags，过滤掉 sys: 或者 volc: 前缀
	if rawTags, ok := props["Tags"]; ok {
		if tagsArr, ok := rawTags.([]interface{}); ok {
			cleanedArr := make([]interface{}, 0, len(tagsArr))
			for _, item := range tagsArr {
				if tagMap, ok := item.(map[string]interface{}); ok {
					key, keyOk := tagMap["Key"].(string)
					value, valueOk := tagMap["Value"].(string)

					if !keyOk || !valueOk {
						cleanedArr = append(cleanedArr, tagMap)
						continue
					}
					// 过滤掉以 sys: 或 volc: 开头的 Key
					if strings.HasPrefix(key, "sys:") || strings.HasPrefix(key, "volc:") {
						continue
					}
					cleaned := make(map[string]interface{})
					cleaned["Key"] = key
					cleaned["Value"] = value
					if len(cleaned) > 0 {
						cleanedArr = append(cleanedArr, cleaned)
					}
				}
			}
			props["Tags"] = cleanedArr
		}
	}
	normalized, err := json.Marshal(props)
	if err != nil {
		return fmt.Errorf("failed to marshal normalized properties: %w", err)
	}

	desc.Properties = volcengine.String(string(normalized))
	return nil
}
