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
	"github.com/volcengine/terraform-provider-volcenginecc/internal/tfresource"
	"github.com/volcengine/terraform-provider-volcenginecc/internal/util"
)

func DeleteResource(ctx context.Context, cloudControlClient *cloudcontrol.CloudControl, region, roleARN, typeName, id string) error {
	tflog.Debug(ctx, "DeleteResource", map[string]interface{}{
		"cfTypeName": typeName,
		"id":         id,
	})

	resp, err := deleteResourceOnce(ctx, cloudControlClient, region, typeName, id)
	if err != nil {
		return wrapCloudControlNotFound(err)
	}
	// An ECS instance that still declares deletion protection rejects the
	// delete outright. The declared intent is to remove it, so clear the
	// protection the resource itself carries and delete once more.
	if resp != nil && isECSDeletionProtectionConflict(typeName, &resp.ProgressEvent) {
		if updateErr := disableECSDeletionProtection(ctx, cloudControlClient, region, typeName, id); updateErr != nil {
			return fmt.Errorf("disable ECS deletion protection before delete: %w", updateErr)
		}
		resp, err = deleteResourceOnce(ctx, cloudControlClient, region, typeName, id)
		if err != nil {
			return wrapCloudControlNotFound(err)
		}
	}
	if resp == nil || resp.OperationStatus == nil {
		return fmt.Errorf("call DeleteResource failed,resp:%s,err:%v ", util.JsonString(resp), err)

	}
	taskId := ""
	status := *resp.OperationStatus
	if status == base.FAILED {
		err := fmt.Errorf("invoke DeleteResource handler failed,resp:%s ", util.JsonString(resp))
		if isCloudControlNotFoundProgressEvent(&resp.ProgressEvent) {
			return &tfresource.NotFoundError{LastError: err}
		}
		return err
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

func deleteResourceOnce(ctx context.Context, cloudControlClient *cloudcontrol.CloudControl, region, typeName, id string) (*cloudcontrol.DeleteResourceOutput, error) {
	return cloudControlClient.DeleteResourceWithContext(ctx, &cloudcontrol.DeleteResourceInput{
		TypeName:    util.StringPtr(typeName),
		RegionID:    util.StringPtr(region),
		Identifier:  util.StringPtr(id),
		ClientToken: util.StringPtr(util.GenerateToken(32)),
	})
}

// disableECSDeletionProtection clears the instance's declared deletion
// protection so the pending delete can proceed.
func disableECSDeletionProtection(ctx context.Context, cloudControlClient *cloudcontrol.CloudControl, region, typeName, id string) error {
	resp, err := cloudControlClient.UpdateResourceWithContext(ctx, &cloudcontrol.UpdateResourceInput{
		TypeName:    util.StringPtr(typeName),
		RegionID:    util.StringPtr(region),
		Identifier:  util.StringPtr(id),
		ClientToken: util.StringPtr(util.GenerateToken(32)),
		PatchDocument: []any{map[string]any{
			"op":    "replace",
			"path":  "/DeletionProtection",
			"value": false,
		}},
	})
	if err != nil {
		return err
	}
	if resp == nil || resp.OperationStatus == nil {
		return fmt.Errorf("call UpdateResource failed,resp:%s", util.JsonString(resp))
	}
	status := *resp.OperationStatus
	if status == base.SUCCESS {
		return nil
	}
	if status != base.IN_PROGRESS && status != base.PENDING {
		return fmt.Errorf("invoke UpdateResource handler failed,resp:%s", util.JsonString(resp))
	}
	if resp.TaskID == nil || *resp.TaskID == "" {
		return fmt.Errorf("invoke UpdateResource handler returned no task,resp:%s", util.JsonString(resp))
	}
	_, _, err = AwaitTask(ctx, cloudControlClient, *resp.TaskID)
	return err
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
			err := fmt.Errorf("invoke get task failed,resp:%s ", util.JsonString(output))
			if isCloudControlNotFoundProgressEvent(output) {
				return nil, false, &tfresource.NotFoundError{LastError: err}
			}
			return nil, false, err
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
