//ff:func feature=cli type=test-helper control=sequence
//ff:what subtestTestPrintReportSARIFWithErrors — WithErrors 서브테스트
package main

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/validate"
)

func subtestTestPrintReportSARIFWithErrors(t *testing.T) {

	r := &validate.Report{Steps: []validate.StepResult{
		{
			Name:   "ssac",
			Status: validate.StatusFail,
			Diagnostics: []diagnostic.Diagnostic{
				{Level: diagnostic.LevelError, Message: "[S-01] bad"},
			},
		},
	}}
	var buf bytes.Buffer
	errs, _, err := printReportSARIF(&buf, r, "")
	if err == nil {
		t.Fatal("expected err for failures")
	}
	if errs != 1 {
		t.Errorf("expected 1 error, got %d", errs)
	}
	var doc map[string]any
	if uerr := json.Unmarshal(buf.Bytes(), &doc); uerr != nil {
		t.Fatalf("invalid JSON: %v", uerr)
	}
	if doc["version"] != "2.1.0" {
		t.Errorf("expected SARIF 2.1.0, got %v", doc["version"])
	}

}
