//ff:func feature=report type=test control=sequence topic=json
//ff:what TestEmitWithDiagnostics — rule_id 추출 + uppercase level + summary 카운트
package json

import (
	stdjson "encoding/json"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/validate"
)

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
					Message: "[S-27] variable foo is not declared",
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
