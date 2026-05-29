//ff:func feature=cli type=test control=sequence topic=format
//ff:what validate -f md — markdown 렌더링 regression gate

package main

import (
	"strings"
	"testing"
)

// TestIntegrationValidate_FormatMD runs `yongol validate <specs> -f md` and
// expects the GFM-lite header plus the preserved "0 errors, 0 warnings"
// footer. Protects the md dispatcher from regressions.
func TestIntegrationValidate_FormatMD(t *testing.T) {
	specs := zenflowSpecsDir(t)
	stdout, _, err := runCmd(t, "validate", specs, "-f", "md")
	if err != nil {
		t.Fatalf("unexpected err: %v\nstdout:\n%s", err, stdout)
	}
	if !strings.HasPrefix(stdout, "## Validation") {
		t.Errorf("expected md output to begin with `## Validation`, got:\n%s", stdout)
	}
	if !strings.Contains(stdout, "0 errors, 0 warnings") {
		t.Errorf("expected `0 errors, 0 warnings` footer, got:\n%s", stdout)
	}
}
