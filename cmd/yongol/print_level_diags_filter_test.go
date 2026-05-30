//ff:func feature=cli type=test control=sequence
//ff:what printLevelDiags level filter test — 불일치 레벨 진단 필터링 검증

package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/validate"
)

func TestPrintLevelDiagsFiltersMismatchedLevel(t *testing.T) {
	var buf bytes.Buffer
	step := validate.StepResult{
		Name: "filter-test",
		Diagnostics: []diagnostic.Diagnostic{
			{Level: diagnostic.LevelWarning, Message: "W-01: something"},
		},
	}
	printLevelDiags(&buf, step, diagnostic.LevelError)
	if buf.Len() != 0 {
		t.Errorf("expected no output when no matching level, got: %q", buf.String())
	}
}

func TestPrintLevelDiagsBranches(t *testing.T) {
	var buf bytes.Buffer
	step := validate.StepResult{
		Name: "step-x",
		Diagnostics: []diagnostic.Diagnostic{
			// Skipped: wrong level.
			{Level: diagnostic.LevelWarning, Message: "skip me"},
			// Matching, explicit Advice field, with file/line.
			{Level: diagnostic.LevelError, Message: "E-01: boom", Advice: "do this", File: "a.yaml", Line: 7},
			// Matching, no Advice field, inline advice split from Message.
			{Level: diagnostic.LevelError, Message: "E-02: bad → Advice: try again"},
			// Matching, no advice at all.
			{Level: diagnostic.LevelError, Message: "E-03: plain"},
		},
	}
	printLevelDiags(&buf, step, diagnostic.LevelError)
	out := buf.String()
	if !strings.Contains(out, "[step-x]") {
		t.Errorf("expected header, got: %q", out)
	}
	if !strings.Contains(out, "E-01: boom") || !strings.Contains(out, "↳ Advice: do this") {
		t.Errorf("expected explicit advice diag, got: %q", out)
	}
	if !strings.Contains(out, "E-02: bad") || !strings.Contains(out, "↳ Advice: try again") {
		t.Errorf("expected split advice diag, got: %q", out)
	}
	if !strings.Contains(out, "E-03: plain") {
		t.Errorf("expected plain diag, got: %q", out)
	}
	if strings.Contains(out, "skip me") {
		t.Errorf("warning diag should be filtered, got: %q", out)
	}
}
