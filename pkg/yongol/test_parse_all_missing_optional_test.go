//ff:func feature=orchestrator type=test control=sequence
//ff:what ParseAll — 선택적 SSOT 누락 시 diag 발생 없이 graceful skip
package yongol

import (
	"path/filepath"
	"testing"
)

// TestParseAllMissingOptionalSSOTs drives ParseAll against a specs root that
// contains only manifest.yaml. All other SSOTs are optional at the parser
// layer (the `*_if_present` contract), so ParseAll must succeed with an empty
// ParseDiagnostics list. This guards against a common regression class:
// parse_X_if_present blindly dereferencing a nil DetectedSSOT when KindX is
// absent, which would either panic or produce spurious diagnostics.
func TestParseAllMissingOptionalSSOTs(t *testing.T) {
	tmp := newTmpSpecsDir(t)
	writeFile(t, filepath.Join(tmp, "manifest.yaml"), minimalManifest)

	detected, err := DetectSSOTs(tmp)
	if err != nil {
		t.Fatalf("DetectSSOTs: %v", err)
	}

	fs := ParseAll(tmp, detected)
	if fs == nil {
		t.Fatalf("ParseAll returned nil")
	}
	if len(fs.ParseDiagnostics) != 0 {
		t.Fatalf("expected no parse diagnostics with only manifest present, got %d: %+v",
			len(fs.ParseDiagnostics), fs.ParseDiagnostics)
	}
	if fs.Manifest == nil {
		t.Fatalf("manifest.yaml present but Manifest not wired")
	}
	if fs.OpenAPIDoc != nil {
		t.Errorf("OpenAPIDoc should stay nil when api/openapi.yaml missing")
	}
	if len(fs.DDLResults) != 0 {
		t.Errorf("DDLResults should be empty when db/ missing")
	}
	if len(fs.ServiceFuncs) != 0 {
		t.Errorf("ServiceFuncs should be empty when service/ missing")
	}
	if fs.Presences[KindConfig] != SSOTPopulated {
		t.Errorf("Presences[KindConfig] = %v; want SSOTPopulated", fs.Presences[KindConfig])
	}
}
