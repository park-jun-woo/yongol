//ff:func feature=validate type=test-helper control=sequence topic=manifest-auth
//ff:what runSec401LiteralSecretRejected — SEC-401 literal secret 거부 케이스 검증 헬퍼

package manifest

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	pm "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func runSec401LiteralSecretRejected(t *testing.T) {
	t.Helper()
	dir := t.TempDir()
	yaml := `apiVersion: yongol/v1
kind: Project
metadata:
  name: test
backend:
  auth:
    type: jwt
    secret: my-super-secret-key
`
	os.WriteFile(filepath.Join(dir, "manifest.yaml"), []byte(yaml), 0o644)
	fs := &yongol.Fullstack{
		SpecsDir: dir,
		Manifest: &pm.ProjectConfig{Backend: pm.Backend{Auth: &pm.Auth{}}},
	}
	diags := sec401JWTSecretEnvRequired(fs)
	if len(diags) != 1 {
		t.Fatalf("expected 1 diag, got %d: %v", len(diags), diags)
	}
	if !strings.Contains(diags[0].Message, "[SEC-401]") {
		t.Errorf("expected [SEC-401], got %q", diags[0].Message)
	}
	if diags[0].Level != diagnostic.LevelError {
		t.Errorf("expected LevelError, got %q", diags[0].Level)
	}
}
