//ff:func feature=report type=test control=sequence topic=json
//ff:what TestEmitRuleIDMissing — rule_id 없는 메시지는 빈 문자열로 유지
package json

import (
	stdjson "encoding/json"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/validate"
)

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
