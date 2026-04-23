//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestGenerateBearerAuth_CookieDefaultCookie — mode=cookie → defaultAuthMode="cookie" embed

package auth

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestGenerateBearerAuth_CookieDefaultCookie pins BUG-014 Phase002 for
// the cookie-mode path: prepared.Auth.Mode="cookie" (explicit) must be
// embedded verbatim as defaultAuthMode; no Phase002 regression drops
// cookie-mode projects into bearer.
func TestGenerateBearerAuth_CookieDefaultCookie(t *testing.T) {
	dir := t.TempDir()
	if err := generateBearerAuth(dir, "example.com/proj", nil, "cookie"); err != nil {
		t.Fatalf("generateBearerAuth: %v", err)
	}
	path := filepath.Join(dir, "backend", "internal", "middleware", "bearerauth.go")
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read emitted file: %v", err)
	}
	got := string(raw)
	if !strings.Contains(got, `return "cookie"`) {
		t.Errorf("expected authMode fallback to return \"cookie\", got:\n%s", got)
	}
}
