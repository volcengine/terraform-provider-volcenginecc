// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package generic

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

// DataSourceOptionsFunc is a type alias for a data source type functional option.
type DataSourceOptionsFunc func(*genericDataSource) error

// dataSourceWithAttributeNameMap is a helper function to construct functional options
// that set a data source type's attribute name maps.
// If multiple dataSourceWithAttributeNameMap calls are made, the last call overrides
// the previous calls' values.
func dataSourceWithAttributeNameMap(v map[string]string) DataSourceOptionsFunc {
	return func(o *genericDataSource) error {
		if _, ok := v["id"]; !ok {
			// Synthesize a mapping for the reserved top-level "id" attribute.
			v["id"] = "ID"
		}

		cfToTfNameMap := make(map[string]string, len(v))

		for tfName, cfName := range v {
			_, ok := cfToTfNameMap[cfName]
			if ok {
				return fmt.Errorf("duplicate attribute name mapping for Cloud Control property %s", cfName)
			}
			cfToTfNameMap[cfName] = tfName
		}

		o.ccToTfNameMap = cfToTfNameMap

		return nil
	}
}

// dataSourceWithCloudControlTypeName is a helper function to construct functional options
// that set a resource type's Cloud Control type name.
// If multiple dataSourceWithCloudControlTypeName calls are made, the last call overrides
// the previous calls' values.
func dataSourceWithCloudControlTypeName(v string) DataSourceOptionsFunc {
	return func(o *genericDataSource) error {
		o.ccTypeName = v

		return nil
	}
}

// dataSourceWithTerraformSchema is a helper function to construct functional options
// that set a resource type's Terraform schema.
// If multiple dataSourceWithTerraformSchema calls are made, the last call overrides
// the previous calls' values.
func dataSourceWithTerraformSchema(v schema.Schema) DataSourceOptionsFunc {
	return func(o *genericDataSource) error {
		o.tfSchema = v

		return nil
	}
}

// dataSourceWithTerraformTypeName is a helper function to construct functional options
// that set a resource type's Terraform type name.
// If multiple dataSourceWithTerraformTypeName calls are made, the last call overrides
// the previous calls' values.
func dataSourceWithTerraformTypeName(v string) DataSourceOptionsFunc {
	return func(o *genericDataSource) error {
		o.tfTypeName = v

		return nil
	}
}

type DataSourceOptions []DataSourceOptionsFunc

// WithAttributeNameMap is a helper function to construct functional options
// that set a resource type's attribute name map, append that function to the
// current slice of functional options and return the new slice of options.
// It is intended to be chained with other similar helper functions in a builder pattern.
func (opts DataSourceOptions) WithAttributeNameMap(v map[string]string) DataSourceOptions {
	return append(opts, dataSourceWithAttributeNameMap(v))
}

// WithCloudControlTypeName is a helper function to construct functional options
// that set a resource type's Cloud Control type name, append that function to the
// current slice of functional options and return the new slice of options.
// It is intended to be chained with other similar helper functions in a builder pattern.
func (opts DataSourceOptions) WithCloudControlTypeName(v string) DataSourceOptions {
	return append(opts, dataSourceWithCloudControlTypeName(v))
}

// WithTerraformSchema is a helper function to construct functional options
// that set a resource type's Terraform schema, append that function to the
// current slice of functional options and return the new slice of options.
// It is intended to be chained with other similar helper functions in a builder pattern.
func (opts DataSourceOptions) WithTerraformSchema(v schema.Schema) DataSourceOptions {
	return append(opts, dataSourceWithTerraformSchema(v))
}

// WithTerraformTypeName is a helper function to construct functional options
// that set a resource type's Terraform type name, append that function to the
// current slice of functional options and return the new slice of options.
// It is intended to be chained with other similar helper functions in a builder pattern.
func (opts DataSourceOptions) WithTerraformTypeName(v string) DataSourceOptions {
	return append(opts, dataSourceWithTerraformTypeName(v))
}

type genericDataSource struct {
	ccToTfNameMap map[string]string // Map of Cloud Control property name to Terraform attribute name
	ccTypeName    string            // Cloud Control type name for the resource type
	tfSchema      schema.Schema     // Terraform schema for the data source type
	tfTypeName    string            // Terraform type name for data source type
}
