//ff:func feature=report type=test control=sequence topic=sarif
//ff:what TestEmitFullCatalogRuleIndex — 진단의 ruleIndex 가 rules[] 의 id 와 역참조 일치
package sarif

import (
	"encoding/json"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	rulecatalog "github.com/park-jun-woo/yongol/pkg/rule/catalog"
	"github.com/park-jun-woo/yongol/pkg/validate"
)

// TestEmitFullCatalogRuleIndex: a fired diagnostic receives a ruleIndex that
// dereferences to the same id inside rules[].
func TestEmitFullCatalogRuleIndex(t *testing.T) {
	cat, err := rulecatalog.Load()
	if err != nil {
		t.Fatalf("catalog.Load: %v", err)
	}
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
			},
		},
	}}
	data, err := Emit(r, "v0.1.21", "", cat)
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
	res := run.Results[0]
	if res.RuleID != "S-27" {
		t.Errorf("results[0].ruleId: got %q, want S-27", res.RuleID)
	}
	if res.RuleIndex == nil {
		t.Fatalf("results[0].ruleIndex should be populated")
	}
	idx := *res.RuleIndex
	if idx < 0 || idx >= len(run.Tool.Driver.Rules) {
		t.Fatalf("ruleIndex %d out of bounds (rules len=%d)", idx, len(run.Tool.Driver.Rules))
	}
	if run.Tool.Driver.Rules[idx].ID != "S-27" {
		t.Errorf("rules[%d].id: got %q, want S-27", idx, run.Tool.Driver.Rules[idx].ID)
	}
}
