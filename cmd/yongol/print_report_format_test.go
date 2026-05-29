//ff:func feature=cli type=test control=sequence topic=format
//ff:what printReport md — 기본 markdown 분기 검증

package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/validate"
)

// TestPrintReportFormatMD covers the default md branch. Verifies the markdown
// renderer emits the familiar "N errors, M warnings" footer plus an H2
// Validation header.
func TestPrintReportFormatMD(t *testing.T) {
	r := &validate.Report{Steps: []validate.StepResult{
		{Name: "step-a", Status: validate.StatusPass},
	}}
	var buf bytes.Buffer
	errs, warns, err := printReport(&buf, r, "md", "")
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if errs != 0 || warns != 0 {
		t.Fatalf("expected (0,0), got (%d,%d)", errs, warns)
	}
	out := buf.String()
	if !strings.HasPrefix(out, "## Validation") {
		t.Errorf("expected md output to start with `## Validation`, got:\n%s", out)
	}
	if !strings.Contains(out, "0 errors, 0 warnings") {
		t.Errorf("expected md output to contain `0 errors, 0 warnings`, got:\n%s", out)
	}
}
