//ff:func feature=orchestrator type=test control=sequence
//ff:what Level constants distinct — LevelError ≠ LevelWarning
package diagnostic_test

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// TestLevel_DistinctValues verifies that LevelError and LevelWarning have distinct values.
func TestLevel_DistinctValues(t *testing.T) {
	if diagnostic.LevelError == diagnostic.LevelWarning {
		t.Errorf("LevelError and LevelWarning must differ, both=%q", diagnostic.LevelError)
	}
}
