//ff:func feature=validate type=test control=sequence topic=manifest-auth
//ff:what SEC-401 테스트 — secret_env 필드는 통과

package manifest

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

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
