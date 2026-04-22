//ff:func feature=orchestrator type=test control=sequence
//ff:what ParseAll — 잘못된 openapi.yaml 이 fs.ParseDiagnostics 로 전파되는지 (P04 게이트 회귀)
package yongol

import (
	"path/filepath"
	"strings"
	"testing"
)

// TestParseAllBrokenOpenAPIEmitsDiagnostic is the PhaseP04 gate-bypass
// regression test. A syntactically broken api/openapi.yaml must surface as a
// parser diagnostic inside fs.ParseDiagnostics so the CLI-level orchestrator
// (which gates validate on len(ParseDiagnostics) > 0) halts. If this test ever
// returns zero diagnostics, ParseAll is silently swallowing a parse failure
// and downstream validate/gen will run on a nil OpenAPIDoc.
func TestParseAllBrokenOpenAPIEmitsDiagnostic(t *testing.T) {
	tmp := newTmpSpecsDir(t)
	writeFile(t, filepath.Join(tmp, "api", "openapi.yaml"), "::: not yaml :::\n")

	detected, err := DetectSSOTs(tmp)
	if err != nil {
		t.Fatalf("DetectSSOTs: %v", err)
	}
	if _, ok := hasKind(detected, KindOpenAPI); !ok {
		t.Fatalf("KindOpenAPI not detected — DetectSSOTs broken independent of gate test")
	}

	fs := ParseAll(tmp, detected)
	if fs == nil {
		t.Fatalf("ParseAll returned nil")
	}
	if len(fs.ParseDiagnostics) == 0 {
		t.Fatalf("PhaseP04 gate regression — broken openapi should emit ParseDiagnostics, got 0")
	}

	// Specifically the OpenAPI loader must be the source (not an unrelated parser).
	found := false
	for _, d := range fs.ParseDiagnostics {
		if strings.Contains(d.Message, "OpenAPI load error") {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("expected an 'OpenAPI load error' diagnostic; got %+v", fs.ParseDiagnostics)
	}

	if fs.OpenAPIDoc != nil {
		t.Errorf("OpenAPIDoc should be nil after parse failure, got %+v", fs.OpenAPIDoc)
	}
}
