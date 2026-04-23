//ff:func feature=cli type=test control=sequence topic=format
//ff:what validate -f sarif — SARIF 2.1.0 driver metadata + single run 검증

package main

import (
	"encoding/json"
	"testing"
)

// TestIntegrationValidate_FormatSARIF verifies the SARIF branch produces a
// valid SARIF 2.1.0 document with the hardcoded driver name + informationUri.
func TestIntegrationValidate_FormatSARIF(t *testing.T) {
	specs := zenflowSpecsDir(t)
	stdout, _, err := runCmd(t, "validate", specs, "-f", "sarif")
	if err != nil {
		t.Fatalf("unexpected err: %v\nstdout:\n%s", err, stdout)
	}
	var doc map[string]any
	if uerr := json.Unmarshal([]byte(stdout), &doc); uerr != nil {
		t.Fatalf("sarif output is not valid JSON: %v\n%s", uerr, stdout)
	}
	if doc["version"] != "2.1.0" {
		t.Errorf("sarif version: got %v, want 2.1.0", doc["version"])
	}
	runs, ok := doc["runs"].([]any)
	if !ok || len(runs) != 1 {
		t.Fatalf("runs: want 1 entry, got %T=%v", doc["runs"], doc["runs"])
	}
	run0 := runs[0].(map[string]any)
	tool := run0["tool"].(map[string]any)
	driver := tool["driver"].(map[string]any)
	if driver["name"] != "yongol" {
		t.Errorf("driver.name: got %v, want yongol", driver["name"])
	}
	if driver["informationUri"] != "https://github.com/park-jun-woo/yongol" {
		t.Errorf("driver.informationUri: got %v", driver["informationUri"])
	}
}
