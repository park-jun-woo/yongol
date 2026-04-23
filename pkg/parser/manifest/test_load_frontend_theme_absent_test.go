//ff:func feature=manifest type=test control=sequence
//ff:what Load — frontend.theme 블록이 없으면 Theme=nil

package manifest

import (
	"os"
	"path/filepath"
	"testing"
)

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
