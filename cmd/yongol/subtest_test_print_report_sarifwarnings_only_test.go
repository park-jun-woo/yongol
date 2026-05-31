//ff:func feature=cli type=test-helper control=sequence
//ff:what subtestTestPrintReportSARIFWarningsOnly — WarningsOnly 서브테스트
package main

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/validate"
)

func subtestTestPrintReportSARIFWarningsOnly(t *testing.T) {

	r := &validate.Report{Steps: []validate.StepResult{
		{
			Name:   "check",
			Status: validate.StatusPass,
			Diagnostics: []diagnostic.Diagnostic{
				{Level: diagnostic.LevelWarning, Message: "[W-01] warn"},
			},
		},
	}}
	var buf bytes.Buffer
	errs, warns, err := printReportSARIF(&buf, r, "/tmp/specs")
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if errs != 0 {
		t.Errorf("expected 0 errors, got %d", errs)
	}
	if warns != 1 {
		t.Errorf("expected 1 warning, got %d", warns)
	}
	var doc map[string]any
	if uerr := json.Unmarshal(buf.Bytes(), &doc); uerr != nil {
		t.Fatalf("invalid JSON: %v", uerr)
	}

}
