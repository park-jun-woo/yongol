//ff:func feature=cli type=test control=sequence topic=json
//ff:what printReportJSON test — JSON 리포트 출력 검증

package main

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/validate"
)

func TestPrintReportJSON(t *testing.T) {
	t.Run("WithErrors", func(t *testing.T) {
		r := &validate.Report{Steps: []validate.StepResult{
			{
				Name:   "ssac",
				Status: validate.StatusFail,
				Diagnostics: []diagnostic.Diagnostic{
					{Level: diagnostic.LevelError, Message: "S-01: bad"},
					{Level: diagnostic.LevelWarning, Message: "S-02: warn"},
				},
			},
		}}
		var buf bytes.Buffer
		errs, warns, err := printReportJSON(&buf, r, "")
		if err == nil {
			t.Fatal("expected err for failures")
		}
		if errs != 1 {
			t.Errorf("expected 1 error, got %d", errs)
		}
		if warns != 1 {
			t.Errorf("expected 1 warning, got %d", warns)
		}
		var doc any
		if uerr := json.Unmarshal(buf.Bytes(), &doc); uerr != nil {
			t.Fatalf("invalid JSON: %v", uerr)
		}
	})

	t.Run("NoErrors", func(t *testing.T) {
		r := &validate.Report{Steps: []validate.StepResult{
			{Name: "step", Status: validate.StatusPass},
		}}
		var buf bytes.Buffer
		errs, warns, err := printReportJSON(&buf, r, "")
		if err != nil {
			t.Fatalf("unexpected err: %v", err)
		}
		if errs != 0 || warns != 0 {
			t.Errorf("expected (0,0), got (%d,%d)", errs, warns)
		}
	})

	t.Run("WriteError", func(t *testing.T) {
		r := &validate.Report{Steps: []validate.StepResult{
			{Name: "step", Status: validate.StatusPass},
		}}
		_, _, err := printReportJSON(&failWriter{}, r, "")
		if err == nil {
			t.Fatal("expected write error")
		}
	})

	t.Run("CatalogChecksPopulated", func(t *testing.T) {
		// Exercises the rulecatalog.Load + cat.Len() path: the emitted JSON
		// summary.checks must reflect the embedded catalog size (> 0), proving
		// the non-error catalog branch wired the count into the document.
		r := &validate.Report{Steps: []validate.StepResult{
			{Name: "step", Status: validate.StatusPass},
		}}
		var buf bytes.Buffer
		if _, _, err := printReportJSON(&buf, r, "/tmp/specs"); err != nil {
			t.Fatalf("unexpected err: %v", err)
		}
		var doc struct {
			Summary struct {
				Checks int `json:"checks"`
			} `json:"summary"`
		}
		if err := json.Unmarshal(buf.Bytes(), &doc); err != nil {
			t.Fatalf("invalid JSON: %v", err)
		}
		if doc.Summary.Checks <= 0 {
			t.Fatalf("expected summary.checks > 0 from embedded catalog, got %d", doc.Summary.Checks)
		}
	})
}
