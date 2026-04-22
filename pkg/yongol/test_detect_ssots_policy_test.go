//ff:func feature=orchestrator type=test control=sequence
//ff:what DetectSSOTs — policy/*.rego presence 감지
package yongol

import (
	"path/filepath"
	"testing"
)

// TestDetectSSOTsPolicyPopulated asserts that a *.rego file under policy/
// marks KindPolicy as SSOTPopulated.
func TestDetectSSOTsPolicyPopulated(t *testing.T) {
	tmp := newTmpSpecsDir(t)
	writeFile(t, filepath.Join(tmp, "policy", "authz.rego"),
		"package authz\n\ndefault allow := false\n")

	detected, err := DetectSSOTs(tmp)
	if err != nil {
		t.Fatalf("DetectSSOTs: %v", err)
	}
	d, ok := hasKind(detected, KindPolicy)
	if !ok {
		t.Fatalf("KindPolicy not detected; detected=%+v", detected)
	}
	if d.Presence != SSOTPopulated {
		t.Fatalf("expected SSOTPopulated, got %v", d.Presence)
	}
}
