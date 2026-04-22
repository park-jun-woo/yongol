//ff:func feature=report type=test control=iteration dimension=1 topic=json
//ff:what test: Emit — flat JSON snake_case 구조 + summary 카운팅 + rule_id 추출 검증
package json

import (
	stdjson "encoding/json"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/validate"
)

// TestEmitEmpty: clean report → zero errors, zero warnings, empty diagnostics.
func TestEmitEmpty(t *testing.T) {
	r := &validate.Report{Steps: []validate.StepResult{
		{Name: "step-a", Status: validate.StatusPass},
	}}
	data, err := Emit(r, "v0.1.21", "./specs", 150)
	if err != nil {
		t.Fatalf("Emit: %v", err)
	}
	var doc Document
	if err := stdjson.Unmarshal(data, &doc); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if doc.YongolVersion != "v0.1.21" {
		t.Errorf("yongol_version: got %q, want v0.1.21", doc.YongolVersion)
	}
	if doc.SpecsDir != "./specs" {
		t.Errorf("specs_dir: got %q, want ./specs", doc.SpecsDir)
	}
	if doc.Summary.Errors != 0 || doc.Summary.Warnings != 0 {
		t.Errorf("summary error/warning: %+v", doc.Summary)
	}
	if doc.Summary.Checks != 150 {
		t.Errorf("summary.checks: got %d, want 150", doc.Summary.Checks)
	}
	if len(doc.Diagnostics) != 0 {
		t.Errorf("diagnostics should be empty, got %d", len(doc.Diagnostics))
	}
	// Verify snake_case keys are literally present in the output.
	raw := string(data)
	for _, key := range []string{
		`"yongol_version"`, `"specs_dir"`, `"summary"`, `"diagnostics"`,
	} {
		if !strings.Contains(raw, key) {
			t.Errorf("missing snake_case key %s in output:\n%s", key, raw)
		}
	}
}

// TestEmitWithDiagnostics: rule_id extraction + uppercase level + count.
func TestEmitWithDiagnostics(t *testing.T) {
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
					Message: "[S-27] 변수 foo 미선언",
				},
				{
					File:    "service/workflow/update.ssac",
					Line:    8,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelWarning,
					Message: "[S-36] stale @response",
				},
			},
		},
	}}
	data, err := Emit(r, "v0.1.21", "", 182)
	if err != nil {
		t.Fatalf("Emit: %v", err)
	}
	var doc Document
	if err := stdjson.Unmarshal(data, &doc); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if doc.Summary.Errors != 1 || doc.Summary.Warnings != 1 {
		t.Errorf("summary counts: %+v", doc.Summary)
	}
	if doc.Summary.Checks != 182 {
		t.Errorf("summary.checks: got %d, want 182", doc.Summary.Checks)
	}
	if len(doc.Diagnostics) != 2 {
		t.Fatalf("diagnostics: got %d, want 2", len(doc.Diagnostics))
	}
	d0 := doc.Diagnostics[0]
	if d0.RuleID != "S-27" || d0.Level != "ERROR" {
		t.Errorf("diagnostics[0]: rule_id=%q level=%q, want S-27/ERROR", d0.RuleID, d0.Level)
	}
	if d0.Line != 15 || d0.File != "service/auth/login.ssac" {
		t.Errorf("diagnostics[0] location: %+v", d0)
	}
	if strings.Contains(d0.Message, "[S-27]") {
		t.Errorf("diagnostics[0].message should not include [S-27]: %q", d0.Message)
	}
	d1 := doc.Diagnostics[1]
	if d1.Level != "WARNING" {
		t.Errorf("diagnostics[1].level: got %q, want WARNING", d1.Level)
	}
}

// TestEmitRuleIDMissing: plain message (no prefix) → rule_id is empty string.
func TestEmitRuleIDMissing(t *testing.T) {
	r := &validate.Report{Steps: []validate.StepResult{
		{
			Name:   "misc",
			Status: validate.StatusFail,
			Diagnostics: []diagnostic.Diagnostic{
				{
					File:    "x.ssac",
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: "plain error without rule id",
				},
			},
		},
	}}
	data, err := Emit(r, "v0.1.21", "", 0)
	if err != nil {
		t.Fatalf("Emit: %v", err)
	}
	var doc Document
	if err := stdjson.Unmarshal(data, &doc); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if len(doc.Diagnostics) != 1 {
		t.Fatalf("diagnostics: got %d, want 1", len(doc.Diagnostics))
	}
	if doc.Diagnostics[0].RuleID != "" {
		t.Errorf("rule_id should be empty, got %q", doc.Diagnostics[0].RuleID)
	}
}
