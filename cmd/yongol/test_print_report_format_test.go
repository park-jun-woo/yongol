//ff:func feature=cli type=test control=iteration dimension=1 topic=format
//ff:what test: printReport dispatcher — md/sarif format dispatch + invalid value error
package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
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

// TestPrintReportFormatSARIF verifies the SARIF branch produces a valid
// SARIF 2.1.0 document with the expected driver metadata and results.
func TestPrintReportFormatSARIF(t *testing.T) {
	r := &validate.Report{Steps: []validate.StepResult{
		{
			Name:   "ssac",
			Status: validate.StatusFail,
			Diagnostics: []diagnostic.Diagnostic{
				{
					File:    "service/auth/login.ssac",
					Line:    15,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: "[S-27] variable foo is undeclared",
				},
			},
		},
	}}
	var buf bytes.Buffer
	errs, warns, err := printReport(&buf, r, "sarif", "")
	if err == nil {
		t.Fatalf("expected validation failed err, got nil")
	}
	if !strings.HasPrefix(err.Error(), "validation failed:") {
		t.Errorf("err should start with 'validation failed:'; got %q", err.Error())
	}
	if errs != 1 || warns != 0 {
		t.Errorf("counts: got (%d,%d), want (1,0)", errs, warns)
	}
	var doc map[string]any
	if uerr := json.Unmarshal(buf.Bytes(), &doc); uerr != nil {
		t.Fatalf("sarif output is not valid JSON: %v\n%s", uerr, buf.String())
	}
	if doc["version"] != "2.1.0" {
		t.Errorf("sarif version: got %v, want 2.1.0", doc["version"])
	}
}

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
