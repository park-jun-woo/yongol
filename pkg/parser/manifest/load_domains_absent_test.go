//ff:func feature=manifest type=test control=sequence
//ff:what Load domains 미선언 시 Domains nil 유지 검증 (단일 사이트 후방 호환)

package manifest

import (
	"os"
	"path/filepath"
	"testing"
)

// TestLoadDomainsAbsent verifies that omitting the domains block leaves
// Domains nil — the single-site path must not regress.
func TestLoadDomainsAbsent(t *testing.T) {
	dir := t.TempDir()
	content := `apiVersion: yongol/v1
kind: Project
metadata:
  name: single
backend:
  module: github.com/test/single
  auth:
    type: jwt
`
	if err := os.WriteFile(filepath.Join(dir, "manifest.yaml"), []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
	cfg, diags := Load(dir)
	if len(diags) > 0 {
		t.Fatalf("Load() diagnostics: %v", diags)
	}
	if cfg.Domains != nil {
		t.Errorf("Domains = %v, want nil", cfg.Domains)
	}
}
