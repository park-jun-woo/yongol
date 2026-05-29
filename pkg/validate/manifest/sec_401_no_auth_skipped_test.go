//ff:func feature=validate type=test control=sequence topic=manifest-auth
//ff:what SEC-401 테스트 — auth 블록이 없으면 미발화

package manifest

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestSEC401_NoAuthSkipped(t *testing.T) {
	fs := &yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{}}
	if diags := sec401JWTSecretEnvRequired(fs); len(diags) != 0 {
		t.Fatalf("expected 0 diagnostics when auth is nil, got %d", len(diags))
	}
}
