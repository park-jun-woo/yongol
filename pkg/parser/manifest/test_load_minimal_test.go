//ff:func feature=manifest type=parser control=sequence
//ff:what 최소 manifest.yaml 로드 시 Auth nil 검증
package manifest

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoad_Minimal(t *testing.T) {
	dir := t.TempDir()
	content := `apiVersion: yongol/v1
kind: Project
metadata:
  name: minimal
backend:
  module: github.com/test/minimal
`
	os.WriteFile(filepath.Join(dir, "manifest.yaml"), []byte(content), 0644)

	cfg, diags := Load(dir)
	if len(diags) > 0 {
		t.Fatalf("Load() diagnostics: %v", diags)
	}
	if cfg.Backend.Auth != nil {
		t.Errorf("expected nil Auth, got %+v", cfg.Backend.Auth)
	}
}
