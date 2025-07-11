// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudcontrol

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/volcengine/terraform-provider-volcenginecc/internal/base"
	"github.com/volcengine/terraform-provider-volcenginecc/internal/cloudcontrol"
	"github.com/volcengine/terraform-provider-volcenginecc/internal/util"
)

func DeleteResource(ctx context.Context, cloudControlClient *cloudcontrol.CloudControl, region, roleARN, typeName, id string) error {
	tflog.Debug(ctx, "DeleteResource", map[string]interface{}{
		"cfTypeName": typeName,
		"id":         id,
	})

	resp, err := cloudControlClient.DeleteResourceWithContext(ctx, &cloudcontrol.DeleteResourceInput{
		TypeName:    util.StringPtr(typeName),
		RegionID:    util.StringPtr(region),
		Identifier:  util.StringPtr(id),
		ClientToken: util.StringPtr(util.GenerateToken(24)),
	})
	if err != nil {
		return err
	}
	if resp == nil || resp.OperationStatus == nil {
		return fmt.Errorf("call DeleteResource failed,resp:%s,err:%v ", util.JsonString(resp), err)

	}
	taskId := ""
	status := *resp.OperationStatus
	if status == base.FAILED {
		return fmt.Errorf("invoke DeleteResource handler failed,resp:%s ", util.JsonString(resp))
	} else if status == base.SUCCESS {
		return nil
	} else if status == base.IN_PROGRESS || status == base.PENDING {
		taskId = *resp.TaskID
	} else {
		return fmt.Errorf("invoke DeleteResource handler unknown failed,resp:%s ", util.JsonString(resp))
	}
	tflog.Info(ctx, "Cloud Control API DeleteResource waiting task ...... ", map[string]interface{}{
		"TaskID":    hclog.Fmt("%v", taskId),
		"RequestID": hclog.Fmt("%v", resp.GetRequestId()),
	})
	_, _, err = AwaitTask(ctx, cloudControlClient, taskId)
	if err != nil {
		return err
	}
	return nil
}

func InvokeGetTask(ctx context.Context, client *cloudcontrol.CloudControl, taskId string) (*cloudcontrol.ProgressEvent, error) {
	input := &cloudcontrol.GetTaskInput{
		TaskID: &taskId,
	}
	output, err := client.GetTaskWithContext(ctx, input)

	if err != nil {
		return nil, err
	}

	return &output.ProgressEvent, nil
}

func AwaitTask(ctx context.Context, client *cloudcontrol.CloudControl, taskId string) (*cloudcontrol.ProgressEvent, bool, error) {
	// 轮询
	times := 0
	AwaitTimeout := 24 * time.Hour
	SleepTime := 10 * time.Second
	for times < int(AwaitTimeout.Seconds()/SleepTime.Seconds()) {
		times++
		if times > 10 {
			time.Sleep(SleepTime)
		} else {
			time.Sleep(time.Duration(times) * time.Second)
		}
		output, err := InvokeGetTask(ctx, client, taskId)
		if err != nil {
			return nil, false, fmt.Errorf("invoke get task failed,resp:%s,err:%v ", util.JsonString(output), err)
		}

		status := *output.OperationStatus
		if status == base.FAILED {
			return nil, false, fmt.Errorf("invoke get task failed,resp:%s,err:%v ", util.JsonString(output), err)
		} else if status == base.SUCCESS {
			return output, true, nil
		} else if status == base.IN_PROGRESS || status == base.PENDING {
			continue
		} else {
			return nil, false, fmt.Errorf("invoke get task unknown failed,resp:%s,err:%v ", util.JsonString(output), err)
		}

	}
	return nil, false, fmt.Errorf("await Task Timeout")
}
