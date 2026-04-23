//ff:func feature=cli type=test control=sequence topic=format
//ff:what validate -f foo — unknown 포맷 시 *usageError (exit 2)

package main

import (
	"strings"
	"testing"
)

// TestIntegrationValidate_FormatUnknown verifies `-f foo` surfaces as a
// *usageError (exit 2) per PhaseC01 exit-code semantics.
func TestIntegrationValidate_FormatUnknown(t *testing.T) {
	specs := zenflowSpecsDir(t)
	_, _, err := runCmd(t, "validate", specs, "-f", "yaml")
	if err == nil {
		t.Fatal("expected error for unknown -f value, got nil")
	}
	if !isUsageError(err) {
		t.Fatalf("unknown -f value should be *usageError (exit 2), got %T: %v", err, err)
	}
	if !strings.Contains(err.Error(), "invalid --format") {
		t.Errorf("err should mention 'invalid --format'; got %q", err.Error())
	}
}
