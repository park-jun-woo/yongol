//ff:func feature=validate type=test control=sequence topic=manifest-auth
//ff:what SEC-401 테스트 — manifest 에 하드코딩된 secret 리터럴 ERROR

package manifest

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

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
