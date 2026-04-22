//ff:func feature=orchestrator type=test control=sequence
//ff:what DetectSSOTs — api/openapi.yaml presence 회귀
package yongol

import (
	"path/filepath"
	"testing"
)

// TestDetectSSOTsOpenAPIPresent asserts that when api/openapi.yaml exists the
// detector returns a KindOpenAPI entry pointing at the file with
// SSOTPopulated presence. Single-file SSOTs never produce SSOTDeclared.
func TestDetectSSOTsOpenAPIPresent(t *testing.T) {
	tmp := newTmpSpecsDir(t)
	writeFile(t, filepath.Join(tmp, "api", "openapi.yaml"), "openapi: 3.0.0\n")

	detected, err := DetectSSOTs(tmp)
	if err != nil {
		t.Fatalf("DetectSSOTs: %v", err)
	}
	d, ok := hasKind(detected, KindOpenAPI)
	if !ok {
		t.Fatalf("KindOpenAPI not detected; detected=%+v", detected)
	}
	if d.Presence != SSOTPopulated {
		t.Fatalf("expected SSOTPopulated for single-file OpenAPI, got %v", d.Presence)
	}
	if filepath.Base(d.Path) != "openapi.yaml" {
		t.Fatalf("unexpected Path %q", d.Path)
	}
}

// TestDetectSSOTsOpenAPIAbsent confirms that a bare specs root with no api/
// directory produces no KindOpenAPI entry at all (SSOTAbsent is omitted).
func TestDetectSSOTsOpenAPIAbsent(t *testing.T) {
	tmp := newTmpSpecsDir(t)

	detected, err := DetectSSOTs(tmp)
	if err != nil {
		t.Fatalf("DetectSSOTs: %v", err)
	}
	if _, ok := hasKind(detected, KindOpenAPI); ok {
		t.Fatalf("KindOpenAPI unexpectedly detected; detected=%+v", detected)
	}
}
