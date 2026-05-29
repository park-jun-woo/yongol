//ff:func feature=cli type=test control=sequence topic=format
//ff:what validate -f sarif — SARIF driver.rules 에 150+ 항목 포함 (PhaseF02)

package main

import (
	"encoding/json"
	"testing"
)

// TestIntegrationValidate_FormatSARIFFullCatalog asserts the PhaseF02 upgrade:
// SARIF rules[] carries the full rulebook catalog (150+ entries) with
// shortDescription populated.
func TestIntegrationValidate_FormatSARIFFullCatalog(t *testing.T) {
	specs := zenflowSpecsDir(t)
	stdout, _, err := runCmd(t, "validate", specs, "-f", "sarif")
	if err != nil {
		t.Fatalf("unexpected err: %v\nstdout:\n%s", err, stdout)
	}
	var doc map[string]any
	if uerr := json.Unmarshal([]byte(stdout), &doc); uerr != nil {
		t.Fatalf("sarif output is not valid JSON: %v\n%s", uerr, stdout)
	}
	runs := doc["runs"].([]any)
	driver := runs[0].(map[string]any)["tool"].(map[string]any)["driver"].(map[string]any)
	rules, ok := driver["rules"].([]any)
	if !ok {
		t.Fatalf("driver.rules missing or wrong type: %T", driver["rules"])
	}
	if len(rules) < 100 {
		t.Errorf("driver.rules: want >=100 (full catalog), got %d", len(rules))
	}
	rule0 := rules[0].(map[string]any)
	sd, ok := rule0["shortDescription"].(map[string]any)
	if !ok || sd["text"] == "" {
		t.Errorf("rules[0].shortDescription should be populated, got %v", rule0["shortDescription"])
	}
	if rule0["helpUri"] == nil || rule0["helpUri"] == "" {
		t.Errorf("rules[0].helpUri should be populated, got %v", rule0["helpUri"])
	}
}
