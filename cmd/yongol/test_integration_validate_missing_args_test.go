//ff:func feature=cli type=test control=sequence
//ff:what validate missing-args — 인자 없을 때 *usageError (exit 2)

package main

import "testing"

// TestIntegrationValidate_MissingArgs verifies exit 2 (usage error) when no
// specs-dir is supplied. PhaseC01 mapped cobra.RangeArgs failures through
// usageArgs → *usageError; this test enforces that contract.
func TestIntegrationValidate_MissingArgs(t *testing.T) {
	_, _, err := runCmd(t, "validate")
	if err == nil {
		t.Fatal("expected usage error for missing arg, got nil")
	}
	if !isUsageError(err) {
		t.Fatalf("expected *usageError (exit 2), got %T: %v", err, err)
	}
}
