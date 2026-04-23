//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestGenerateBearerAuth_JwtDefaultBearer — manifest.auth.type=jwt → defaultAuthMode="bearer" embed

package auth

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestGenerateBearerAuth_JwtDefaultBearer pins BUG-014 Phase002: when
// manifest.backend.auth.type=jwt and auth.mode is empty, prepared.Auth.Mode
// resolves to "bearer" and generateBearerAuth embeds that value as the
// defaultAuthMode const. No BACKEND_AUTH_MODE env → authMode() returns
// "bearer".
func TestGenerateBearerAuth_JwtDefaultBearer(t *testing.T) {
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
	if strings.Contains(got, `return "cookie"`) {
		t.Errorf("jwt project must not emit cookie as authMode fallback, got:\n%s", got)
	}
}
