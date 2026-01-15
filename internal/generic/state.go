// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package generic

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// copyStateValueAtPath copies the value at a specified path from source State to destination State.
func copyStateValueAtPath(ctx context.Context, dst, src *tfsdk.State, p path.Path) diag.Diagnostics {
	var diags diag.Diagnostics
	var val attr.Value
	currentPath := path.Empty()
	currentPathList := []path.Path{currentPath}
	for _, step := range p.Steps() {
		stepNeedCopyList := make([]path.Path, 0)
		for _, cur := range currentPathList {
			currentPathTemp := cur.AtName(step.String())
			diags = src.GetAttribute(ctx, currentPathTemp, &val)
			if diags.HasError() {
				return diags
			}
			// 如果遇到 List 类型，则递归遍历元素
			switch v := val.(type) {
			case types.List:
				if v.IsNull() || v.IsUnknown() {
					continue
				}
				for i := range v.Elements() {
					// 在 Path 上增加 index 并递归处理
					subPath := currentPathTemp.AtListIndex(i)
					stepNeedCopyList = append(stepNeedCopyList, subPath)
				}
			case types.Set:
				//如果遇到Set类型，直接退出，不进行WriteOnly字段的赋值
				return diags
			default:
				//继续遍历
				stepNeedCopyList = append(stepNeedCopyList, currentPathTemp)
			}
		}
		currentPathList = stepNeedCopyList
	}
	for _, copyPath := range currentPathList {
		diags = copyStateValue(ctx, dst, src, copyPath)
		if diags.HasError() {
			return diags
		}
	}
	return diags
}

// copyStateValue copies the value at a specified path from source State to destination State.
func copyStateValue(ctx context.Context, dst, src *tfsdk.State, p path.Path) diag.Diagnostics {
	var diags diag.Diagnostics
	var val attr.Value

	diags.Append(src.GetAttribute(ctx, p, &val)...)
	if diags.HasError() {
		return diags
	}

	diags.Append(dst.SetAttribute(ctx, p, val)...)
	if diags.HasError() {
		return diags
	}

	return diags
}
