//ff:func feature=orchestrator type=test control=sequence
//ff:what DetectSSOTs — tests/scenario-*.hurl, invariant-*.hurl 둘 다 감지
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

// TestDetectSSOTsScenarioInvariant verifies the invariant-*.hurl glob branch.
// Mirrors the scenario-*.hurl branch so both inputs stay covered.
func TestDetectSSOTsScenarioInvariant(t *testing.T) {
	tmp := newTmpSpecsDir(t)
	writeFile(t, filepath.Join(tmp, "tests", "invariant-tenant-breach.hurl"),
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
		t.Fatalf("expected SSOTPopulated (invariant glob), got %v", d.Presence)
	}
}
