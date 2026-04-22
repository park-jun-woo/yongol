//ff:func feature=report type=test control=iteration dimension=1 topic=sarif
//ff:what test: Emit(..., catalog) — verifies full catalog rules are included, ruleIndex linkage, and rule meta fields
package sarif

import (
	"encoding/json"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	rulecatalog "github.com/park-jun-woo/yongol/pkg/rule/catalog"
	"github.com/park-jun-woo/yongol/pkg/validate"
)

// TestEmitFullCatalogEmpty: with a catalog but no diagnostics, rules[] is
// still populated (every catalogued rule) and results[] is empty.
func TestEmitFullCatalogEmpty(t *testing.T) {
	cat, err := rulecatalog.Load()
	if err != nil {
		t.Fatalf("catalog.Load: %v", err)
	}
	r := &validate.Report{Steps: []validate.StepResult{
		{Name: "step-a", Status: validate.StatusPass},
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
	if len(run.Tool.Driver.Rules) != cat.Len() {
		t.Errorf("driver.rules: got %d, want %d", len(run.Tool.Driver.Rules), cat.Len())
	}
	if len(run.Results) != 0 {
		t.Errorf("results: got %d, want 0", len(run.Results))
	}
	// Rule entry should carry shortDescription + helpUri + defaultConfiguration.
	first := run.Tool.Driver.Rules[0]
	if first.ShortDescription == nil || first.ShortDescription.Text == "" {
		t.Errorf("rules[0].shortDescription missing: %+v", first)
	}
	if first.HelpURI == "" {
		t.Errorf("rules[0].helpUri missing: %+v", first)
	}
	if first.DefaultConfiguration == nil {
		t.Errorf("rules[0].defaultConfiguration missing: %+v", first)
	}
}

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
