//ff:func feature=cli type=test control=sequence
//ff:what generate args-count — artsDir 누락 시 *usageError (exit 2)

package main

import (
	"testing"
)

// TestIntegrationGenerate_ArgsCount verifies that `generate <specs>` (artsdir
// missing) exits with code 2 via *usageError, matching validate's
// RangeArgs→usageArgs contract from PhaseC01.
func TestIntegrationGenerate_ArgsCount(t *testing.T) {
	specs := zenflowSpecsDir(t)
	_, _, err := runCmd(t, "generate", specs)
	if err == nil {
		t.Fatal("expected usage error for missing artsDir, got nil")
	}
	if !isUsageError(err) {
		t.Fatalf("expected *usageError (exit 2), got %T: %v", err, err)
	}
}
