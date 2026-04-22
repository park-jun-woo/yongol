//ff:func feature=cli type=test control=iteration dimension=1 topic=format
//ff:what test: yongol validate --format md|sarif|<bad> end-to-end cases (md default, sarif JSON, usage error)
package main

import (
	"encoding/json"
	"strings"
	"testing"
)

// TestIntegrationValidate_FormatMD runs `yongol validate <specs> -f md` and
// expects the GFM-lite header plus the preserved "0 errors, 0 warnings"
// footer. Protects the md dispatcher from regressions.
func TestIntegrationValidate_FormatMD(t *testing.T) {
	specs := zenflowSpecsDir(t)
	stdout, _, err := runCmd(t, "validate", specs, "-f", "md")
	if err != nil {
		t.Fatalf("unexpected err: %v\nstdout:\n%s", err, stdout)
	}
	if !strings.HasPrefix(stdout, "## Validation") {
		t.Errorf("expected md output to begin with `## Validation`, got:\n%s", stdout)
	}
	if !strings.Contains(stdout, "0 errors, 0 warnings") {
		t.Errorf("expected `0 errors, 0 warnings` footer, got:\n%s", stdout)
	}
}

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

// TestIntegrationValidate_FormatUnknown verifies `-f foo` surfaces as a
// *usageError (exit 2) per PhaseC01 exit-code semantics.
func TestIntegrationValidate_FormatUnknown(t *testing.T) {
	specs := zenflowSpecsDir(t)
	_, _, err := runCmd(t, "validate", specs, "-f", "yaml")
	if err == nil {
		t.Fatal("expected error for unknown -f value, got nil")
	}
	if !isUsageError(err) {
		t.Fatalf("unknown -f value should be *usageError (exit 2), got %T: %v", err, err)
	}
	if !strings.Contains(err.Error(), "invalid --format") {
		t.Errorf("err should mention 'invalid --format'; got %q", err.Error())
	}
}
