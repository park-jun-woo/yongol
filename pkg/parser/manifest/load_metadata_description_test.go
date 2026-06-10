//ff:func feature=manifest type=test control=sequence
//ff:what Load — metadata.description(yongol init 가 방출하는 키)가 엄격 디코딩에서도 수용되는지 검증

package manifest

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoad_MetadataDescriptionAccepted(t *testing.T) {
	dir := t.TempDir()
	// `metadata.description` is emitted by `yongol init`; strict decoding must
	// still accept it (promoted to a formal field).
	content := `apiVersion: yongol/v1
kind: Project
metadata:
  name: described
  description: "a project with a description"
backend:
  module: github.com/test/described
`
	if err := os.WriteFile(filepath.Join(dir, "manifest.yaml"), []byte(content), 0o644); err != nil {
		t.Fatalf("write fixture: %v", err)
	}

	cfg, diags := Load(dir)
	if len(diags) > 0 {
		t.Fatalf("Load() diagnostics: %v", diags)
	}
	if cfg.Metadata.Description != "a project with a description" {
		t.Errorf("Description = %q, want %q", cfg.Metadata.Description, "a project with a description")
	}
}
