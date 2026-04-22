//ff:func feature=orchestrator type=test control=sequence
//ff:what ParseAll — zenflow 전체 SSOT 회귀 (OpenAPI/DDL/SSaC/Scenario/State 적재)
package yongol

import (
	"testing"
)

// TestParseAllZenflow drives ParseAll over the shared zenflow fixture to keep
// the full orchestrator integration honest. Failure here means one of the
// parse_*_if_present helpers silently stopped wiring results onto Fullstack
// — which is the class of regression PhaseP04 / P05 previously hit.
func TestParseAllZenflow(t *testing.T) {
	if testing.Short() {
		t.Skip("zenflow fixture is heavy; skipped in -short mode")
	}
	specsDir := findZenflowSpecsAbs(t)

	detected, err := DetectSSOTs(specsDir)
	if err != nil {
		t.Fatalf("DetectSSOTs: %v", err)
	}
	if len(detected) == 0 {
		t.Fatalf("expected zenflow fixture to yield detected SSOTs, got 0")
	}

	fs := ParseAll(specsDir, detected)
	if fs == nil {
		t.Fatalf("ParseAll returned nil Fullstack")
	}
	if fs.SpecsDir != specsDir {
		t.Fatalf("SpecsDir = %q; want %q", fs.SpecsDir, specsDir)
	}
	if fs.OpenAPIDoc == nil {
		t.Errorf("OpenAPIDoc is nil — parse_open_api_if_present did not wire the doc")
	}
	if len(fs.DDLResults) == 0 {
		t.Errorf("DDLResults empty — parse_ddl_if_present did not wire results")
	}
	if len(fs.ServiceFuncs) == 0 {
		t.Errorf("ServiceFuncs empty — parse_ssa_c_if_present did not wire service funcs")
	}
	if len(fs.HurlEntries) == 0 {
		t.Errorf("HurlEntries empty — parse_scenario_if_present did not wire entries")
	}
	if fs.Manifest == nil {
		t.Errorf("Manifest nil — parse_manifest_if_present did not wire config")
	}
	if fs.Presences[KindOpenAPI] != SSOTPopulated {
		t.Errorf("Presences[KindOpenAPI] = %v; want SSOTPopulated", fs.Presences[KindOpenAPI])
	}
}
