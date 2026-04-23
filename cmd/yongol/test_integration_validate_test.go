//ff:func feature=cli type=test control=sequence
//ff:what validate happy — zenflow specs 로 clean report (0 errors)

package main

import (
	"strings"
	"testing"
)

// TestIntegrationValidate_Happy runs `yongol validate <zenflow-specs>` and
// expects a clean report (0 errors, 0 warnings) with exit 0. Regression gate
// for the end-to-end validate pipeline — any upstream parser/rule that
// starts flagging zenflow breaks this.
func TestIntegrationValidate_Happy(t *testing.T) {
	specs := zenflowSpecsDir(t)
	stdout, _, err := runCmd(t, "validate", specs)
	if err != nil {
		t.Fatalf("validate happy: unexpected error: %v\nstdout:\n%s", err, stdout)
	}
	if !strings.Contains(stdout, "0 errors, 0 warnings") {
		t.Errorf("expected stdout to contain `0 errors, 0 warnings`, got:\n%s", stdout)
	}
}
