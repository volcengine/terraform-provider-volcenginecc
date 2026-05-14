// Copyright (c) 2025 Beijing Volcano Engine Technology Co., Ltd.
// SPDX-License-Identifier: MPL-2.0

// Package customresources lets hand-written code override the resource
// implementation produced by the auto-generated `*_resource_gen.go` files
// without modifying those generated files (which would be lost on the next
// `make resources` run).
//
// Usage:
//
//  1. Auto-generated code in `internal/volcengine/.../foo_resource_gen.go`
//     registers itself as usual via `registry.AddResourceFactory("foo", fooResource)`.
//  2. A hand-written file in the same package, e.g. `foo_resource.go`, calls
//     `customresources.Register("foo", wrapFoo)` from its `init()` function.
//  3. The provider's `Resources()` method consults this registry and replaces
//     the auto-generated resource with the wrapped one.
//
// The wrapper receives the inner (auto-generated) resource so it can delegate
// most of the framework lifecycle while customizing only the parts it needs
// (typically Create or Read).
package customresources

import (
	"context"
	"sync"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Factory wraps the auto-generated inner resource and returns the resource
// that should be exposed to the framework in its place.
type Factory func(ctx context.Context, inner resource.Resource) (resource.Resource, error)

var (
	mu       sync.Mutex
	registry = make(map[string]Factory)
)

// Register associates a custom Factory with a Terraform resource type name.
// Registering the same name twice panics to surface conflicts at startup.
func Register(typeName string, f Factory) {
	mu.Lock()
	defer mu.Unlock()
	if _, exists := registry[typeName]; exists {
		panic("customresources: duplicate registration for " + typeName)
	}
	registry[typeName] = f
}

// Lookup returns the custom Factory for typeName, if any.
func Lookup(typeName string) (Factory, bool) {
	mu.Lock()
	defer mu.Unlock()
	f, ok := registry[typeName]
	return f, ok
}
