//ff:func feature=report type=test control=sequence topic=sarif
//ff:what TestEmitFullCatalogEmpty — 카탈로그 있으면 rules[] 전체 규칙 방출, results 는 비어 있음
package sarif

import (
	"encoding/json"
	"testing"

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
