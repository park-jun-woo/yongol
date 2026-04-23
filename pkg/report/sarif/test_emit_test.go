//ff:func feature=report type=test control=sequence topic=sarif
//ff:what TestEmitEmpty — 빈 리포트 → driver metadata + 0 results 검증
package sarif

import (
	"encoding/json"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/validate"
)

// TestEmitEmpty covers the minimal empty report — no diagnostics, no rules.
func TestEmitEmpty(t *testing.T) {
	r := &validate.Report{Steps: []validate.StepResult{
		{Name: "step-a", Status: validate.StatusPass},
	}}
	data, err := Emit(r, "v0.1.20", "", nil)
	if err != nil {
		t.Fatalf("Emit: %v", err)
	}
	var doc Document
	if err := json.Unmarshal(data, &doc); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if doc.Version != "2.1.0" {
		t.Errorf("version: got %q, want 2.1.0", doc.Version)
	}
	if doc.Schema == "" {
		t.Errorf("expected $schema to be populated")
	}
	if len(doc.Runs) != 1 {
		t.Fatalf("runs: got %d, want 1", len(doc.Runs))
	}
	run := doc.Runs[0]
	if run.Tool.Driver.Name != "yongol" {
		t.Errorf("driver name: got %q, want yongol", run.Tool.Driver.Name)
	}
	if run.Tool.Driver.Version != "v0.1.20" {
		t.Errorf("driver version: got %q, want v0.1.20", run.Tool.Driver.Version)
	}
	if run.Tool.Driver.InformationURI == "" {
		t.Errorf("informationUri should be populated")
	}
	if len(run.Tool.Driver.Rules) != 0 {
		t.Errorf("rules should be empty for no diagnostics, got %d", len(run.Tool.Driver.Rules))
	}
	if len(run.Results) != 0 {
		t.Errorf("results should be empty, got %d", len(run.Results))
	}
}
