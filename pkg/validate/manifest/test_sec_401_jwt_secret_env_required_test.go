//ff:func feature=validate type=test control=sequence topic=manifest-auth
//ff:what SEC-401 테스트 — secret 리터럴 감지 + 정상 secret_env 통과

package manifest

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

// writeManifest writes manifest.yaml into a tmp dir and returns the dir.
func writeManifest(t *testing.T, body string) string {
	t.Helper()
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "manifest.yaml"), []byte(body), 0o644); err != nil {
		t.Fatalf("write manifest: %v", err)
	}
	return dir
}

func TestSEC401_LiteralSecretRejected(t *testing.T) {
	body := `apiVersion: yongol/v1
kind: Project
metadata:
  name: test
backend:
  module: example.com/app
  auth:
    type: jwt
    secret: "hardcoded-secret-hardcoded-secret-xx"
    claims:
      ID: "user_id:int64"
`
	dir := writeManifest(t, body)
	fs := &yongol.Fullstack{
		SpecsDir: dir,
		Manifest: &pmanifest.ProjectConfig{Backend: pmanifest.Backend{Auth: &pmanifest.Auth{Type: "jwt"}}},
	}
	diags := sec401JWTSecretEnvRequired(fs)
	if len(diags) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d: %+v", len(diags), diags)
	}
	if diags[0].Level != diagnostic.LevelError {
		t.Fatalf("expected ERROR, got %v", diags[0].Level)
	}
}

func TestSEC401_SecretEnvPasses(t *testing.T) {
	body := `apiVersion: yongol/v1
kind: Project
metadata:
  name: test
backend:
  module: example.com/app
  auth:
    type: jwt
    secret_env: JWT_SECRET
    claims:
      ID: "user_id:int64"
`
	dir := writeManifest(t, body)
	fs := &yongol.Fullstack{
		SpecsDir: dir,
		Manifest: &pmanifest.ProjectConfig{Backend: pmanifest.Backend{Auth: &pmanifest.Auth{Type: "jwt", SecretEnv: "JWT_SECRET"}}},
	}
	if diags := sec401JWTSecretEnvRequired(fs); len(diags) != 0 {
		t.Fatalf("expected 0 diagnostics, got %d: %+v", len(diags), diags)
	}
}

func TestSEC401_NoAuthSkipped(t *testing.T) {
	fs := &yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{}}
	if diags := sec401JWTSecretEnvRequired(fs); len(diags) != 0 {
		t.Fatalf("expected 0 diagnostics when auth is nil, got %d", len(diags))
	}
}
