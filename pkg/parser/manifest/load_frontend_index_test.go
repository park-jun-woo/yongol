//ff:func feature=manifest type=test control=sequence
//ff:what Load — frontend.index 가 엄격 디코딩(KnownFields) 하에서 합법 키로 수용되는지 검증

package manifest

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoad_FrontendIndexIsKnownField(t *testing.T) {
	dir := t.TempDir()
	// frontend.index (page-flow Phase009) is a Frontend struct field, so
	// strict decoding (KnownFields(true), BUG-115) must accept the key —
	// adding the field is what legalizes it (no separate allowlist).
	content := `apiVersion: yongol/v1
kind: Project
metadata:
  name: strict
backend:
  module: github.com/test/strict
  auth:
    type: jwt
frontend:
  lang: typescript
  framework: react
  index: dashboard
`
	if err := os.WriteFile(filepath.Join(dir, "manifest.yaml"), []byte(content), 0o644); err != nil {
		t.Fatalf("write fixture: %v", err)
	}

	cfg, diags := Load(dir)
	if len(diags) != 0 {
		t.Fatalf("expected no diagnostics for frontend.index, got %+v", diags)
	}
	if cfg == nil || cfg.Frontend.Index != "dashboard" {
		t.Errorf("Frontend.Index = %+v, want %q", cfg, "dashboard")
	}
}
