// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudcontrol

import (
	"errors"
	"net/http"
	"strings"

	"github.com/volcengine/terraform-provider-volcenginecc/internal/cloudcontrol"
	"github.com/volcengine/terraform-provider-volcenginecc/internal/tfresource"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineerr"
)

const (
	cloudControlNotFoundHandlerErrorCode = "(HandlerErrorCode: NotFound)"

	ecsInstanceTypeName                  = "Volcengine::ECS::Instance"
	ecsDeletionProtectionCode            = "AccessDenied"
	ecsDeletionProtectionStatusSubstring = "InvalidOperation.DeletionProtection:"
)

// isECSDeletionProtectionConflict identifies the exact delete rejection that
// can be recovered by disabling the instance's declared protection first.
func isECSDeletionProtectionConflict(typeName string, event *cloudcontrol.ProgressEvent) bool {
	return typeName == ecsInstanceTypeName &&
		event != nil &&
		event.ErrorCode != nil && *event.ErrorCode == ecsDeletionProtectionCode &&
		event.StatusMessage != nil && strings.Contains(*event.StatusMessage, ecsDeletionProtectionStatusSubstring)
}

// wrapCloudControlNotFound converts errors that explicitly identify a missing
// Cloud Control resource into the provider's shared NotFoundError type.
func wrapCloudControlNotFound(err error) error {
	if err == nil || tfresource.NotFound(err) {
		return err
	}

	if isCloudControlNotFoundError(err) {
		return &tfresource.NotFoundError{LastError: err}
	}

	return err
}

// isCloudControlNotFoundError checks structured SDK error data first and uses
// the exact Cloud Control handler marker only as a compatibility fallback.
func isCloudControlNotFoundError(err error) bool {
	var requestFailure volcengineerr.RequestFailure
	if errors.As(err, &requestFailure) && requestFailure.StatusCode() == http.StatusNotFound {
		return true
	}

	var sdkError volcengineerr.Error
	if errors.As(err, &sdkError) && isCloudControlNotFoundCode(sdkError.Code()) {
		return true
	}

	return strings.Contains(err.Error(), cloudControlNotFoundHandlerErrorCode)
}

// isCloudControlNotFoundProgressEvent reports whether a failed asynchronous
// operation explicitly returned a resource-not-found handler error.
func isCloudControlNotFoundProgressEvent(event *cloudcontrol.ProgressEvent) bool {
	if event == nil {
		return false
	}

	if event.ErrorCode != nil && isCloudControlNotFoundCode(*event.ErrorCode) {
		return true
	}

	return event.StatusMessage != nil && strings.Contains(*event.StatusMessage, cloudControlNotFoundHandlerErrorCode)
}

// isCloudControlNotFoundCode accepts the generic Cloud Control NotFound code
// and service-specific codes whose final, complete segment is NotFound.
func isCloudControlNotFoundCode(code string) bool {
	return code == "NotFound" || strings.HasSuffix(code, ".NotFound")
}
