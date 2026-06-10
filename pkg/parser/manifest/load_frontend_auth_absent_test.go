//ff:func feature=manifest type=test control=sequence
//ff:what Load — frontend.auth 블록이 없으면 Auth 는 nil

package manifest

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoad_FrontendAuthAbsent(t *testing.T) {
	dir := t.TempDir()
	content := `apiVersion: yongol/v1
kind: Project
metadata:
  name: plain
backend:
  module: github.com/test/plain
frontend:
  framework: react
`
	if err := os.WriteFile(filepath.Join(dir, "manifest.yaml"), []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
	cfg, diags := Load(dir)
	if len(diags) > 0 {
		t.Fatalf("Load() diagnostics: %v", diags)
	}
	if cfg.Frontend.Auth != nil {
		t.Fatalf("expected nil Auth when frontend.auth absent, got %+v", cfg.Frontend.Auth)
	}
}
