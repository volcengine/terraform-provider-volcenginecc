package schemalint

import (
	_ "embed"
	"encoding/json"
	"fmt"
)

// Baseline maps a Cloud Control type name to the violation paths published
// before the naming rules took effect. Those fields keep their names and are
// exempt from the checks. Entries are only ever removed: new fields must
// follow the rules instead of being added here.
type Baseline map[string][]string

//go:embed naming_baseline.json
var baselineJSON []byte

// DefaultBaseline returns the embedded naming baseline.
func DefaultBaseline() (Baseline, error) {
	var baseline Baseline
	if err := json.Unmarshal(baselineJSON, &baseline); err != nil {
		return nil, fmt.Errorf("parsing schema naming baseline: %w", err)
	}

	return baseline, nil
}
