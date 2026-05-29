//ff:func feature=cli type=test control=sequence topic=format
//ff:what printReport default — 빈 포맷 문자열이 md 로 fallback 되는지

package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/validate"
)

// TestPrintReportFormatDefault verifies empty format falls back to md.
func TestPrintReportFormatDefault(t *testing.T) {
	r := &validate.Report{Steps: []validate.StepResult{
		{Name: "step-a", Status: validate.StatusPass},
	}}
	var buf bytes.Buffer
	if _, _, err := printReport(&buf, r, "", ""); err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if !strings.HasPrefix(buf.String(), "## Validation") {
		t.Errorf("empty format should default to md; got:\n%s", buf.String())
	}
}
