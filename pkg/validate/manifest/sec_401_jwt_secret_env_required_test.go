//ff:func feature=validate type=test control=sequence topic=manifest-auth
//ff:what sec401JWTSecretEnvRequired — literal secret 금지, secret_env만 허용 검증

package manifest

import (
	"os"
	"path/filepath"
	"testing"

	pm "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestSec401JWTSecretEnvRequired(t *testing.T) {
	t.Run("nil_fs", func(t *testing.T) {
		diags := sec401JWTSecretEnvRequired(nil)
		if len(diags) != 0 {
			t.Errorf("expected 0 diags, got %d", len(diags))
		}
	})

	t.Run("no_specs_dir", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		diags := sec401JWTSecretEnvRequired(fs)
		if len(diags) != 0 {
			t.Errorf("expected 0 diags, got %d", len(diags))
		}
	})

	t.Run("no_auth", func(t *testing.T) {
		fs := &yongol.Fullstack{SpecsDir: "/tmp", Manifest: &pm.ProjectConfig{}}
		diags := sec401JWTSecretEnvRequired(fs)
		if len(diags) != 0 {
			t.Errorf("expected 0 diags, got %d", len(diags))
		}
	})

	t.Run("manifest_file_not_found", func(t *testing.T) {
		fs := &yongol.Fullstack{
			SpecsDir: "/nonexistent",
			Manifest: &pm.ProjectConfig{Backend: pm.Backend{Auth: &pm.Auth{}}},
		}
		diags := sec401JWTSecretEnvRequired(fs)
		if len(diags) != 0 {
			t.Errorf("expected 0 diags, got %d", len(diags))
		}
	})

	t.Run("manifest_without_auth_node", func(t *testing.T) {
		dir := t.TempDir()
		yamlContent := `apiVersion: yongol/v1
kind: Project
metadata:
  name: test
backend:
  module: github.com/org/test
`
		os.WriteFile(filepath.Join(dir, "manifest.yaml"), []byte(yamlContent), 0o644)
		fs := &yongol.Fullstack{
			SpecsDir: dir,
			Manifest: &pm.ProjectConfig{Backend: pm.Backend{Auth: &pm.Auth{}}},
		}
		diags := sec401JWTSecretEnvRequired(fs)
		if len(diags) != 0 {
			t.Errorf("expected 0 diags, got %d", len(diags))
		}
	})

	t.Run("secret_env_allowed", func(t *testing.T) {
		dir := t.TempDir()
		yaml := `apiVersion: yongol/v1
kind: Project
metadata:
  name: test
backend:
  auth:
    type: jwt
    secret_env: JWT_SECRET
`
		os.WriteFile(filepath.Join(dir, "manifest.yaml"), []byte(yaml), 0o644)
		fs := &yongol.Fullstack{
			SpecsDir: dir,
			Manifest: &pm.ProjectConfig{Backend: pm.Backend{Auth: &pm.Auth{SecretEnv: "JWT_SECRET"}}},
		}
		diags := sec401JWTSecretEnvRequired(fs)
		if len(diags) != 0 {
			t.Errorf("expected 0 diags, got %d: %v", len(diags), diags)
		}
	})

	t.Run("literal_secret_produces_error", func(t *testing.T) {
		runSec401LiteralSecretRejected(t)
	})
}
