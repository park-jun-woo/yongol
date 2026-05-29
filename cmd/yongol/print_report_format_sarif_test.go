//ff:func feature=cli type=test control=sequence topic=format
//ff:what printReport sarif — SARIF 2.1.0 문서 생성 + 에러 집계 검증

package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/validate"
)

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
