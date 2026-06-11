//ff:func feature=manifest type=test control=sequence
//ff:what Load — frontend.auth.role_field 파싱(role_field 전용 블록)과 오타 키 즉시 ERROR(엄격 디코딩) 검증

package manifest

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoad_FrontendAuthRoleField(t *testing.T) {
	t.Run("role_field-only block parses", func(t *testing.T) {
		dir := t.TempDir()
		content := `apiVersion: yongol/v1
kind: Project
metadata:
  name: roled
backend:
  module: github.com/test/roled
frontend:
  framework: react
  auth:
    role_field: role
`
		if err := os.WriteFile(filepath.Join(dir, "manifest.yaml"), []byte(content), 0644); err != nil {
			t.Fatal(err)
		}
		cfg, diags := Load(dir)
		if len(diags) > 0 {
			t.Fatalf("Load() diagnostics: %v", diags)
		}
		if cfg.Frontend.Auth == nil || cfg.Frontend.Auth.RoleField != "role" {
			t.Fatalf("role_field: want \"role\", got %+v", cfg.Frontend.Auth)
		}
		if !cfg.Frontend.Auth.RoleFieldOnly() {
			t.Errorf("RoleFieldOnly() = false, want true for a role_field-only block")
		}
	})

	t.Run("typo key is rejected by strict decoding", func(t *testing.T) {
		dir := t.TempDir()
		content := `apiVersion: yongol/v1
kind: Project
metadata:
  name: roled
backend:
  module: github.com/test/roled
frontend:
  framework: react
  auth:
    rolefield: role
`
		if err := os.WriteFile(filepath.Join(dir, "manifest.yaml"), []byte(content), 0644); err != nil {
			t.Fatal(err)
		}
		if _, diags := Load(dir); len(diags) == 0 {
			t.Errorf("expected strict-decoding diagnostics for the typo key, got none")
		}
	})
}
