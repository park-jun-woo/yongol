//ff:func feature=cli type=test control=sequence topic=sarif
//ff:what printReportSARIF test — SARIF 2.1.0 리포트 출력 검증

package main

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/validate"
)

func TestPrintReportSARIF(t *testing.T) {
	t.Run("WithErrors", func(t *testing.T) {
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
	})

	t.Run("NoErrors", func(t *testing.T) {
		r := &validate.Report{Steps: []validate.StepResult{
			{Name: "step", Status: validate.StatusPass},
		}}
		var buf bytes.Buffer
		errs, warns, err := printReportSARIF(&buf, r, "")
		if err != nil {
			t.Fatalf("unexpected err: %v", err)
		}
		if errs != 0 || warns != 0 {
			t.Errorf("expected (0,0), got (%d,%d)", errs, warns)
		}
	})

	t.Run("WarningsOnly", func(t *testing.T) {
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
	})

	t.Run("WriteError", func(t *testing.T) {
		r := &validate.Report{Steps: []validate.StepResult{
			{Name: "step", Status: validate.StatusPass},
		}}
		_, _, err := printReportSARIF(&failWriter{}, r, "")
		if err == nil {
			t.Fatal("expected write error")
		}
	})
}
