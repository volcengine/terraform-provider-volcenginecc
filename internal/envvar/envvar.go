// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package envvar

import (
	"os"
)

// Standard Volcengine environment variables used in the Terraform Volcengine Provider testing.
const (
	DefaultRegion = "VOLCNEINGE_REGION"
)

// GetWithDefault gets an environment variable value if non-empty or returns the default.
func GetWithDefault(variable string, defaultValue string) string {
	value := os.Getenv(variable)

	if value == "" {
		return defaultValue
	}

	return value
}
