//ff:func feature=manifest type=test control=sequence topic=http-config
//ff:what Load — backend.http 블록이 있으면 body/multipart/overrides 파싱

package manifest

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoad_HTTPConfig(t *testing.T) {
	dir := t.TempDir()
	content := `apiVersion: yongol/v1
kind: Project
metadata:
  name: httpcfg
backend:
  module: github.com/test/httpcfg
  http:
    body_limit: 2MiB
    multipart_limit: 64MiB
    overrides:
      UploadAvatar:
        body_limit: 5MiB
      UploadDocument:
        multipart_limit: 100MiB
`
	if err := os.WriteFile(filepath.Join(dir, "manifest.yaml"), []byte(content), 0644); err != nil {
		t.Fatalf("write: %v", err)
	}

	cfg, diags := Load(dir)
	if len(diags) > 0 {
		t.Fatalf("Load() diagnostics: %v", diags)
	}
	if cfg.Backend.HTTP == nil {
		t.Fatalf("expected non-nil HTTP config")
	}
	if cfg.Backend.HTTP.BodyLimit != "2MiB" {
		t.Errorf("BodyLimit = %q, want 2MiB", cfg.Backend.HTTP.BodyLimit)
	}
	if cfg.Backend.HTTP.MultipartLimit != "64MiB" {
		t.Errorf("MultipartLimit = %q, want 64MiB", cfg.Backend.HTTP.MultipartLimit)
	}
	if len(cfg.Backend.HTTP.Overrides) != 2 {
		t.Fatalf("Overrides len = %d, want 2", len(cfg.Backend.HTTP.Overrides))
	}
	if cfg.Backend.HTTP.Overrides["UploadAvatar"].BodyLimit != "5MiB" {
		t.Errorf("UploadAvatar body_limit = %q, want 5MiB", cfg.Backend.HTTP.Overrides["UploadAvatar"].BodyLimit)
	}
}
