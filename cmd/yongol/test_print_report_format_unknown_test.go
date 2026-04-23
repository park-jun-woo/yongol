//ff:func feature=cli type=test control=sequence topic=format
//ff:what printReport unknown — 알 수 없는 포맷 시 에러 반환

package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/validate"
)

// TestPrintReportFormatUnknown verifies the dispatcher rejects unknown formats
// with a descriptive error rather than silently falling back.
func TestPrintReportFormatUnknown(t *testing.T) {
	r := &validate.Report{Steps: []validate.StepResult{
		{Name: "step-a", Status: validate.StatusPass},
	}}
	var buf bytes.Buffer
	_, _, err := printReport(&buf, r, "yaml", "")
	if err == nil {
		t.Fatal("expected error for unknown format, got nil")
	}
	if !strings.Contains(err.Error(), "unknown format") {
		t.Errorf("err should mention 'unknown format'; got %q", err.Error())
	}
}
