//ff:func feature=cli type=test control=sequence topic=md
//ff:what printReportMD test — GFM-lite 리포트 출력 검증

package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/validate"
)

func TestPrintReportMD(t *testing.T) {
	t.Run("AllPass", func(t *testing.T) {
		r := &validate.Report{Steps: []validate.StepResult{
			{Name: "step-a", Status: validate.StatusPass},
		}}
		var buf bytes.Buffer
		errs, warns, err := printReportMD(&buf, r)
		if err != nil {
			t.Fatalf("unexpected err: %v", err)
		}
		if errs != 0 || warns != 0 {
			t.Errorf("expected (0,0), got (%d,%d)", errs, warns)
		}
		out := buf.String()
		if !strings.Contains(out, "## Validation") {
			t.Errorf("expected header, got: %q", out)
		}
		if !strings.Contains(out, "0 errors, 0 warnings") {
			t.Errorf("expected footer, got: %q", out)
		}
	})

	t.Run("WithErrorsAndWarnings", func(t *testing.T) {
		r := &validate.Report{Steps: []validate.StepResult{
			{
				Name:   "ssac",
				Status: validate.StatusFail,
				Diagnostics: []diagnostic.Diagnostic{
					{Level: diagnostic.LevelError, Message: "S-01: bad"},
					{Level: diagnostic.LevelWarning, Message: "S-02: check"},
				},
			},
		}}
		var buf bytes.Buffer
		errs, warns, err := printReportMD(&buf, r)
		if err == nil {
			t.Fatal("expected err for failures")
		}
		if errs != 1 || warns != 1 {
			t.Errorf("expected (1,1), got (%d,%d)", errs, warns)
		}
		out := buf.String()
		if !strings.Contains(out, "### Errors") {
			t.Errorf("expected Errors section, got: %q", out)
		}
		if !strings.Contains(out, "### Warnings") {
			t.Errorf("expected Warnings section, got: %q", out)
		}
	})
}
