//ff:func feature=report type=test control=iteration dimension=1 topic=json
//ff:what TestEmitEmpty — clean report 에서 zero counts + snake_case 키 노출 검증
package json

import (
	stdjson "encoding/json"
	"strings"
	"testing"

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
