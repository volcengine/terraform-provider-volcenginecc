// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package naming_test

import (
	"testing"

	"github.com/volcengine/terraform-provider-volcenginecc/internal/naming"
)

func TestParseCloudControlTypeName(t *testing.T) {
	testCases := []struct {
		TestName             string
		Value                string
		ExpectError          bool
		ExpectedOrganization string
		ExpectedService      string
		ExpectedResource     string
	}{
		{
			TestName:    "empty string",
			Value:       "",
			ExpectError: true,
		},
		{
			TestName:    "whitespace string",
			Value:       "  ",
			ExpectError: true,
		},
		{
			TestName:    "incorrect number of segments",
			Value:       "ve::EC2",
			ExpectError: true,
		},
		{
			TestName:    "invalid type name",
			Value:       "ve::KMS::WayTooLongAResourceName000000000000000000000000000000000000000012",
			ExpectError: true,
		},
		{
			TestName:    "Terraform type name",
			Value:       "ve_kms_key",
			ExpectError: true,
		},
		{
			TestName:             "valid type name",
			Value:                "ve::KMS::Key",
			ExpectedOrganization: "ve",
			ExpectedService:      "KMS",
			ExpectedResource:     "Key",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.TestName, func(t *testing.T) {
			gotOrganization, gotService, gotResource, err := naming.ParseCloudControlTypeName(testCase.Value)

			if err == nil && testCase.ExpectError {
				t.Fatalf("expected error, got no error")
			}

			if err != nil && !testCase.ExpectError {
				t.Fatalf("got unexpected error: %s", err)
			}

			if gotOrganization != testCase.ExpectedOrganization {
				t.Errorf("expected Organization: %s, got: %s", testCase.ExpectedOrganization, gotOrganization)
			}
			if gotService != testCase.ExpectedService {
				t.Errorf("expected Service: %s, got: %s", testCase.ExpectedService, gotService)
			}
			if gotResource != testCase.ExpectedResource {
				t.Errorf("expected Resource: %s, got: %s", testCase.ExpectedResource, gotResource)
			}
		})
	}
}

func TestCreateTerraformTypeName(t *testing.T) {
	testCases := []struct {
		TestName      string
		Organization  string
		Service       string
		Resource      string
		ExpectedValue string
	}{
		{
			TestName:      "valid type name",
			Organization:  "ve",
			Service:       "kms",
			Resource:      "key",
			ExpectedValue: "ve_kms_key",
		},
		{
			TestName:      "valid type name multiple underscores in resource",
			Organization:  "ve",
			Service:       "logs",
			Resource:      "log_group",
			ExpectedValue: "ve_logs_log_group",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.TestName, func(t *testing.T) {
			gotValue := naming.CreateTerraformTypeName(testCase.Organization, testCase.Service, testCase.Resource)

			if gotValue != testCase.ExpectedValue {
				t.Errorf("expected type name: %s, got: %s", testCase.ExpectedValue, gotValue)
			}
		})
	}
}

func TestParseTerraformTypeName(t *testing.T) {
	testCases := []struct {
		TestName             string
		Value                string
		ExpectError          bool
		ExpectedOrganization string
		ExpectedService      string
		ExpectedResource     string
	}{
		{
			TestName:    "empty string",
			Value:       "",
			ExpectError: true,
		},
		{
			TestName:    "whitespace string",
			Value:       "  ",
			ExpectError: true,
		},
		{
			TestName:    "incorrect number of segments",
			Value:       "ve_ec2",
			ExpectError: true,
		},
		{
			TestName:    "invalid type name",
			Value:       "ve_kms_k-y",
			ExpectError: true,
		},
		{
			TestName:    "CloudControl type name",
			Value:       "ve::KMS::Key",
			ExpectError: true,
		},
		{
			TestName:             "valid type name",
			Value:                "ve_kms_key",
			ExpectedOrganization: "ve",
			ExpectedService:      "kms",
			ExpectedResource:     "key",
		},
		{
			TestName:             "valid type name multiple underscores in resource",
			Value:                "ve_logs_log_group",
			ExpectedOrganization: "ve",
			ExpectedService:      "logs",
			ExpectedResource:     "log_group",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.TestName, func(t *testing.T) {
			gotOrganization, gotService, gotResource, err := naming.ParseTerraformTypeName(testCase.Value)

			if err == nil && testCase.ExpectError {
				t.Fatalf("expected error, got no error")
			}

			if err != nil && !testCase.ExpectError {
				t.Fatalf("got unexpected error: %s", err)
			}

			if gotOrganization != testCase.ExpectedOrganization {
				t.Errorf("expected Organization: %s, got: %s", testCase.ExpectedOrganization, gotOrganization)
			}
			if gotService != testCase.ExpectedService {
				t.Errorf("expected Service: %s, got: %s", testCase.ExpectedService, gotService)
			}
			if gotResource != testCase.ExpectedResource {
				t.Errorf("expected Resource: %s, got: %s", testCase.ExpectedResource, gotResource)
			}
		})
	}
}
