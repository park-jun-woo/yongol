//ff:func feature=orchestrator type=test control=sequence
//ff:what DetectSSOTs — tests/scenario-*.hurl 분기
package yongol

import (
	"path/filepath"
	"testing"
)

// TestDetectSSOTsScenarioHappy verifies the scenario-*.hurl glob branch.
func TestDetectSSOTsScenarioHappy(t *testing.T) {
	tmp := newTmpSpecsDir(t)
	writeFile(t, filepath.Join(tmp, "tests", "scenario-login.hurl"),
		"GET {{host}}/ping\nHTTP 200\n")

	detected, err := DetectSSOTs(tmp)
	if err != nil {
		t.Fatalf("DetectSSOTs: %v", err)
	}
	d, ok := hasKind(detected, KindScenario)
	if !ok {
		t.Fatalf("KindScenario not detected; detected=%+v", detected)
	}
	if d.Presence != SSOTPopulated {
		t.Fatalf("expected SSOTPopulated, got %v", d.Presence)
	}
}
