//ff:func feature=manifest type=parser control=sequence
//ff:what frontend.theme 블록 파싱 검증 (Phase003)

package manifest

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoad_FrontendTheme(t *testing.T) {
	dir := t.TempDir()
	content := `apiVersion: yongol/v1
kind: Project
metadata:
  name: themed
backend:
  module: github.com/test/themed
frontend:
  framework: react
  theme:
    primary: "#3b82f6"
    destructive: "#ef4444"
    radius: "0.75rem"
`
	if err := os.WriteFile(filepath.Join(dir, "manifest.yaml"), []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
	cfg, diags := Load(dir)
	if len(diags) > 0 {
		t.Fatalf("Load() diagnostics: %v", diags)
	}
	if cfg.Frontend.Theme == nil {
		t.Fatalf("expected non-nil Theme")
	}
	if cfg.Frontend.Theme.Primary != "#3b82f6" {
		t.Errorf("primary: want #3b82f6, got %q", cfg.Frontend.Theme.Primary)
	}
	if cfg.Frontend.Theme.Destructive != "#ef4444" {
		t.Errorf("destructive: want #ef4444, got %q", cfg.Frontend.Theme.Destructive)
	}
	if cfg.Frontend.Theme.Radius != "0.75rem" {
		t.Errorf("radius: want 0.75rem, got %q", cfg.Frontend.Theme.Radius)
	}
}

func TestLoad_FrontendTheme_Absent(t *testing.T) {
	dir := t.TempDir()
	content := `apiVersion: yongol/v1
kind: Project
metadata:
  name: no-theme
backend:
  module: github.com/test/no-theme
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
	if cfg.Frontend.Theme != nil {
		t.Errorf("expected nil Theme when block omitted, got %+v", cfg.Frontend.Theme)
	}
}
