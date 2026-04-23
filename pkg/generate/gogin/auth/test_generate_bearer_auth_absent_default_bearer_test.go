//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestGenerateBearerAuth_AbsentDefaultBearer — auth 미선언 경로에서도 bearer 기본값 유지 가능

package auth

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestGenerateBearerAuth_AbsentDefaultBearer pins BUG-014 Phase002 for
// the auth-undeclared path: when the manifest has no backend.auth block
// auth.Generate short-circuits and emits nothing, so BearerAuthStrict is
// never wired. The test asserts prepared.New observes Auth.Present=false
// (so no bearerauth.go gets written), and that generateBearerAuth called
// directly with defaultMode="bearer" still emits a well-formed file —
// defending against a future regression where the bearer literal is lost
// from the template.
func TestGenerateBearerAuth_AbsentDefaultBearer(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{Backend: pmanifest.Backend{}},
	}
	p := prepared.New(fs)
	if p.Auth.Present {
		t.Fatalf("prepared.Auth.Present=true when manifest has no auth; want false")
	}

	dir := t.TempDir()
	if err := generateBearerAuth(dir, "example.com/proj", nil, "bearer"); err != nil {
		t.Fatalf("generateBearerAuth: %v", err)
	}
	path := filepath.Join(dir, "backend", "internal", "middleware", "bearerauth.go")
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read emitted file: %v", err)
	}
	got := string(raw)
	if !strings.Contains(got, `return "bearer"`) {
		t.Errorf("expected authMode fallback to return \"bearer\", got:\n%s", got)
	}
}
