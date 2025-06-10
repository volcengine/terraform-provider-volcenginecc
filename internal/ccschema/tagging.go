// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ccschema

type Tagging struct {
	Taggable     *bool                `json:"taggable,omitempty"`
	TagOnCreate  *bool                `json:"tagOnCreate,omitempty"`
	TagUpdatable *bool                `json:"tagUpdatable,omitempty"`
	TagProperty  *PropertyJsonPointer `json:"tagProperty,omitempty"`
	Permissions  []string             `json:"permissions,omitempty"`
}
