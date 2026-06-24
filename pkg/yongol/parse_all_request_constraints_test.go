//ff:func feature=orchestrator type=test control=sequence
//ff:what ParseAll — 단일 사이트 OpenAPI 적재 시 Request/ResponseConstraints 가 여전히 추출되는지 (Phase004 추출 회귀)
package yongol

import (
	"path/filepath"
	"testing"
)

// TestParseAllSingleSiteConstraintsPreserved guards the Phase004 parseOpenAPI
// extraction: after delegating doc/lines loading to parseOpenAPI, the single-site
// loader (parseOpenAPIIfPresent) MUST still call ExtractRequest/ResponseConstraints.
// 19 readers depend on fs.RequestConstraints / fs.ResponseConstraints, so a
// regression here would silently break request/response validation.
func TestParseAllSingleSiteConstraintsPreserved(t *testing.T) {
	tmp := newTmpSpecsDir(t)
	writeFile(t, filepath.Join(tmp, "manifest.yaml"), minimalManifest)
	writeFile(t, filepath.Join(tmp, "api", "openapi.yaml"), validDomainOpenAPI)

	detected, err := DetectSSOTs(tmp)
	if err != nil {
		t.Fatalf("DetectSSOTs: %v", err)
	}
	if _, ok := hasKind(detected, KindOpenAPI); !ok {
		t.Fatalf("KindOpenAPI not detected — fixture wiring broken")
	}

	fs := ParseAll(tmp, detected)
	if fs == nil {
		t.Fatalf("ParseAll returned nil")
	}
	if fs.OpenAPIDoc == nil {
		t.Fatalf("OpenAPIDoc nil — single-site OpenAPI not loaded")
	}
	// Single-site is NOT domained: the singular fields are the data source.
	if fs.IsDomained() {
		t.Fatalf("IsDomained() = true for a manifest without a domains block")
	}
	if fs.RequestConstraints == nil {
		t.Fatalf("RequestConstraints nil — ExtractRequestConstraints no longer called after parseOpenAPI extraction")
	}
	if fs.ResponseConstraints == nil {
		t.Fatalf("ResponseConstraints nil — ExtractResponseConstraints no longer called after parseOpenAPI extraction")
	}
	// The crafted CreateThing operation has a maxLength request-body constraint,
	// so extraction must actually yield it (not just a non-nil empty map).
	if _, ok := fs.RequestConstraints["CreateThing"]; !ok {
		t.Errorf("RequestConstraints[CreateThing] missing — constraint extraction regressed; got %+v", fs.RequestConstraints)
	}
}
