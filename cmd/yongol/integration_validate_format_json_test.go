//ff:func feature=cli type=test control=iteration dimension=1 topic=format
//ff:what validate -f json — yongol 전용 flat JSON + summary.checks 검증

package main

import (
	"encoding/json"
	"testing"
)

// TestIntegrationValidate_FormatJSON verifies the json branch produces the
// yongol bespoke flat document with snake_case keys and a populated
// summary.checks value sourced from the embedded catalog.
func TestIntegrationValidate_FormatJSON(t *testing.T) {
	specs := zenflowSpecsDir(t)
	stdout, _, err := runCmd(t, "validate", specs, "-f", "json")
	if err != nil {
		t.Fatalf("unexpected err: %v\nstdout:\n%s", err, stdout)
	}
	var doc map[string]any
	if uerr := json.Unmarshal([]byte(stdout), &doc); uerr != nil {
		t.Fatalf("json output is not valid JSON: %v\n%s", uerr, stdout)
	}
	for _, key := range []string{"yongol_version", "specs_dir", "summary", "diagnostics"} {
		if _, ok := doc[key]; !ok {
			t.Errorf("missing top-level snake_case key %q", key)
		}
	}
	summary, ok := doc["summary"].(map[string]any)
	if !ok {
		t.Fatalf("summary is not an object: %T", doc["summary"])
	}
	if summary["errors"].(float64) != 0 {
		t.Errorf("summary.errors: want 0, got %v", summary["errors"])
	}
	if summary["warnings"].(float64) != 0 {
		t.Errorf("summary.warnings: want 0, got %v", summary["warnings"])
	}
	if summary["checks"].(float64) < 100 {
		t.Errorf("summary.checks: want >=100, got %v", summary["checks"])
	}
	diags, ok := doc["diagnostics"].([]any)
	if !ok {
		t.Fatalf("diagnostics is not an array: %T", doc["diagnostics"])
	}
	if len(diags) != 0 {
		t.Errorf("diagnostics: want 0 on clean zenflow run, got %d", len(diags))
	}
}
