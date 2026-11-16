// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudcontrol

import (
	"errors"
	"net/http"
	"testing"

	ccsdk "github.com/volcengine/terraform-provider-volcenginecc/internal/cloudcontrol"
	"github.com/volcengine/terraform-provider-volcenginecc/internal/tfresource"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineerr"
)

// TestWrapCloudControlNotFound verifies that only explicit Cloud Control
// resource-not-found responses are converted to tfresource.NotFoundError.
func TestWrapCloudControlNotFound(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		err      error
		expected bool
	}{
		"exact SDK code": {
			err:      volcengineerr.New("NotFound", "resource does not exist", nil),
			expected: true,
		},
		"HTTP 404": {
			err: volcengineerr.NewRequestFailure(
				volcengineerr.New("UnknownError", "resource does not exist", nil),
				http.StatusNotFound,
				"request-id",
			),
			expected: true,
		},
		"exact handler marker": {
			err:      errors.New("handler failed (HandlerErrorCode: NotFound)"),
			expected: true,
		},
		"service-specific NotFound code": {
			err:      volcengineerr.New("InvalidNatGateway.NotFound", "resource does not exist", nil),
			expected: true,
		},
		"partial service code suffix is not enough": {
			err: volcengineerr.New("InvalidNatGateway.NotFoundExtra", "different service error", nil),
		},
		"partial handler marker is not enough": {
			err: errors.New("handler failed (HandlerErrorCode: NotFoundExtra)"),
		},
		"ordinary error": {
			err: errors.New("permission denied"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			err := wrapCloudControlNotFound(test.err)
			if got := tfresource.NotFound(err); got != test.expected {
				t.Fatalf("tfresource.NotFound() = %t, want %t; error: %v", got, test.expected, err)
			}
		})
	}
}

// TestIsCloudControlNotFoundProgressEvent verifies the strict ErrorCode and
// StatusMessage boundaries used for failed synchronous and asynchronous tasks.
func TestIsCloudControlNotFoundProgressEvent(t *testing.T) {
	t.Parallel()

	notFound := "NotFound"
	serviceCode := "InvalidNatGateway.NotFound"
	partialServiceCode := "InvalidNatGateway.NotFoundExtra"
	exactMarker := "handler failed (HandlerErrorCode: NotFound)"
	partialMarker := "handler failed (HandlerErrorCode: NotFoundExtra)"

	tests := map[string]struct {
		event    *ccsdk.ProgressEvent
		expected bool
	}{
		"nil event": {},
		"exact error code": {
			event:    &ccsdk.ProgressEvent{ErrorCode: &notFound},
			expected: true,
		},
		"service-specific NotFound code": {
			event:    &ccsdk.ProgressEvent{ErrorCode: &serviceCode},
			expected: true,
		},
		"partial service code suffix is not enough": {
			event: &ccsdk.ProgressEvent{ErrorCode: &partialServiceCode},
		},
		"exact status marker": {
			event:    &ccsdk.ProgressEvent{StatusMessage: &exactMarker},
			expected: true,
		},
		"partial status marker is not enough": {
			event: &ccsdk.ProgressEvent{StatusMessage: &partialMarker},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if got := isCloudControlNotFoundProgressEvent(test.event); got != test.expected {
				t.Fatalf("isCloudControlNotFoundProgressEvent() = %t, want %t", got, test.expected)
			}
		})
	}
}

// TestIsECSDeletionProtectionConflict verifies that only an ECS instance
// rejected for its declared deletion protection is treated as recoverable.
func TestIsECSDeletionProtectionConflict(t *testing.T) {
	t.Parallel()

	accessDenied := ecsDeletionProtectionCode
	otherCode := "OtherError"
	exactMessage := "AccessDenied: InvalidOperation.DeletionProtection: deletion is protected"
	unrelatedMessage := "AccessDenied: InvalidOperation.Something: denied"

	tests := map[string]struct {
		typeName string
		event    *ccsdk.ProgressEvent
		expected bool
	}{
		"ECS instance rejected for deletion protection": {
			typeName: ecsInstanceTypeName,
			event:    &ccsdk.ProgressEvent{ErrorCode: &accessDenied, StatusMessage: &exactMessage},
			expected: true,
		},
		"another resource type with the same rejection": {
			typeName: "Volcengine::ECS::KeyPair",
			event:    &ccsdk.ProgressEvent{ErrorCode: &accessDenied, StatusMessage: &exactMessage},
			expected: false,
		},
		"different error code": {
			typeName: ecsInstanceTypeName,
			event:    &ccsdk.ProgressEvent{ErrorCode: &otherCode, StatusMessage: &exactMessage},
			expected: false,
		},
		"access denied for an unrelated reason": {
			typeName: ecsInstanceTypeName,
			event:    &ccsdk.ProgressEvent{ErrorCode: &accessDenied, StatusMessage: &unrelatedMessage},
			expected: false,
		},
		"no event": {
			typeName: ecsInstanceTypeName,
			event:    nil,
			expected: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if got := isECSDeletionProtectionConflict(test.typeName, test.event); got != test.expected {
				t.Fatalf("isECSDeletionProtectionConflict = %v, want %v", got, test.expected)
			}
		})
	}
}
