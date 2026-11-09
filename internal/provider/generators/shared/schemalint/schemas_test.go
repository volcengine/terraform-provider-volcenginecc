package schemalint

import (
	"path/filepath"
	"testing"

	"github.com/volcengine/terraform-provider-volcenginecc/internal/ccschema"
)

const schemaCacheDirectory = "../../../../service/cloudcontrol/schemas"

// TestSchemas validates every cached resource schema against the naming
// baseline and requires each baseline entry to still exempt a violation, so
// the baseline shrinks as published schemas are fixed.
func TestSchemas(t *testing.T) {
	t.Parallel()

	baseline, err := DefaultBaseline()
	if err != nil {
		t.Fatal(err)
	}

	files, err := filepath.Glob(filepath.Join(schemaCacheDirectory, "*.json"))
	if err != nil {
		t.Fatal(err)
	}
	if len(files) == 0 {
		t.Fatalf("no schemas found in %s", schemaCacheDirectory)
	}

	typeNames := make(map[string]bool)
	for _, file := range files {
		schema, err := ccschema.NewResourceJsonSchemaPath(file)
		if err != nil {
			t.Fatalf("%s: %s", file, err)
		}
		resource, err := schema.Resource()
		if err != nil {
			t.Fatalf("%s: %s", file, err)
		}
		if err := resource.Expand(); err != nil {
			t.Fatalf("%s: %s", file, err)
		}
		typeName := *resource.TypeName
		typeNames[typeName] = true

		for _, v := range Validate(resource, baseline) {
			t.Errorf("%s %s", typeName, v)
		}

		violating := make(map[string]bool)
		for _, v := range Validate(resource, nil) {
			violating[v.Path] = true
		}
		for _, path := range baseline[typeName] {
			if !violating[path] {
				t.Errorf("stale naming baseline entry %s %s: remove it from naming_baseline.json", typeName, path)
			}
		}
	}

	for typeName := range baseline {
		if !typeNames[typeName] {
			t.Errorf("stale naming baseline type %s: remove it from naming_baseline.json", typeName)
		}
	}
}
