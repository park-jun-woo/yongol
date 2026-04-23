//ff:func feature=manifest type=test control=sequence topic=http-config
//ff:what Load — backend.http 블록이 없으면 HTTP=nil

package manifest

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoad_HTTPConfigMissing(t *testing.T) {
	dir := t.TempDir()
	content := `apiVersion: yongol/v1
kind: Project
metadata:
  name: nohttp
backend:
  module: github.com/test/nohttp
`
	if err := os.WriteFile(filepath.Join(dir, "manifest.yaml"), []byte(content), 0644); err != nil {
		t.Fatalf("write: %v", err)
	}
	cfg, diags := Load(dir)
	if len(diags) > 0 {
		t.Fatalf("Load() diagnostics: %v", diags)
	}
	if cfg.Backend.HTTP != nil {
		t.Errorf("expected nil HTTP config, got %+v", cfg.Backend.HTTP)
	}
}
