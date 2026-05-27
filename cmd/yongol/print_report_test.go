//ff:func feature=cli type=test control=sequence topic=format
//ff:what printReport test — 모든 포맷 분기 (md, json, sarif, unknown) 커버리지

package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/validate"
)

func TestPrintReport(t *testing.T) {
	passReport := &validate.Report{Steps: []validate.StepResult{
		{Name: "step-a", Status: validate.StatusPass},
	}}
	failReport := &validate.Report{Steps: []validate.StepResult{
		{
			Name:   "ssac",
			Status: validate.StatusFail,
			Diagnostics: []diagnostic.Diagnostic{
				{Level: diagnostic.LevelError, Message: "[S-01] missing arg", File: "a.ssac", Line: 10},
			},
		},
	}}

	t.Run("DefaultToMD", func(t *testing.T) {
		var buf bytes.Buffer
		_, _, err := printReport(&buf, passReport, "", "")
		if err != nil {
			t.Fatalf("unexpected err: %v", err)
		}
		if !strings.HasPrefix(buf.String(), "## Validation") {
			t.Errorf("expected md output, got: %q", buf.String())
		}
	})

	t.Run("MD", func(t *testing.T) {
		var buf bytes.Buffer
		_, _, err := printReport(&buf, passReport, "md", "")
		if err != nil {
			t.Fatalf("unexpected err: %v", err)
		}
		if !strings.Contains(buf.String(), "0 errors") {
			t.Errorf("expected 0 errors, got: %q", buf.String())
		}
	})

	t.Run("JSON", func(t *testing.T) {
		var buf bytes.Buffer
		errs, _, err := printReport(&buf, failReport, "json", "")
		if err == nil {
			t.Fatal("expected err")
		}
		if errs != 1 {
			t.Errorf("expected 1 error, got %d", errs)
		}
		var doc any
		if uerr := json.Unmarshal(buf.Bytes(), &doc); uerr != nil {
			t.Fatalf("invalid JSON: %v", uerr)
		}
	})

	t.Run("SARIF", func(t *testing.T) {
		var buf bytes.Buffer
		errs, _, err := printReport(&buf, failReport, "sarif", "")
		if err == nil {
			t.Fatal("expected err")
		}
		if errs != 1 {
			t.Errorf("expected 1 error, got %d", errs)
		}
		var doc map[string]any
		if uerr := json.Unmarshal(buf.Bytes(), &doc); uerr != nil {
			t.Fatalf("invalid JSON: %v", uerr)
		}
		if doc["version"] != "2.1.0" {
			t.Errorf("expected sarif version 2.1.0, got %v", doc["version"])
		}
	})

	t.Run("Unknown", func(t *testing.T) {
		var buf bytes.Buffer
		_, _, err := printReport(&buf, passReport, "yaml", "")
		if err == nil {
			t.Fatal("expected err for unknown format")
		}
		if !strings.Contains(err.Error(), "unknown format") {
			t.Errorf("expected 'unknown format' in error, got: %q", err.Error())
		}
	})
}
