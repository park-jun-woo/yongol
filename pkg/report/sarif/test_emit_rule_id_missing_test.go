//ff:func feature=report type=test control=sequence topic=sarif
//ff:what TestEmitRuleIDMissing — 전치 RULE-ID 없는 메시지는 ruleId 비어 있음
package sarif

import (
	"encoding/json"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/validate"
)

// TestEmitRuleIDMissing — diagnostic messages without a [RULE-ID] prefix keep
// ruleId empty and do not contribute to tool.driver.rules.
func TestEmitRuleIDMissing(t *testing.T) {
	r := &validate.Report{Steps: []validate.StepResult{
		{
			Name:   "misc",
			Status: validate.StatusFail,
			Diagnostics: []diagnostic.Diagnostic{
				{
					File:    "x.ssac",
					Line:    1,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: "plain error message without rule id",
				},
			},
		},
	}}
	data, err := Emit(r, "v0.1.20", "", nil)
	if err != nil {
		t.Fatalf("Emit: %v", err)
	}
	var doc Document
	if err := json.Unmarshal(data, &doc); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	run := doc.Runs[0]
	if len(run.Results) != 1 {
		t.Fatalf("results: got %d, want 1", len(run.Results))
	}
	if run.Results[0].RuleID != "" {
		t.Errorf("ruleId should be empty, got %q", run.Results[0].RuleID)
	}
	if len(run.Tool.Driver.Rules) != 0 {
		t.Errorf("driver.rules should be empty (no fired rule id), got %d",
			len(run.Tool.Driver.Rules))
	}
}
