//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestGenerateBearerAuth_HybridDefaultHybrid — mode=hybrid → defaultAuthMode="hybrid" embed

package auth

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestGenerateBearerAuth_HybridDefaultHybrid pins BUG-014 Phase002 for
// the hybrid-mode path: prepared.Auth.Mode="hybrid" must be embedded
// verbatim as defaultAuthMode so the emitted authMode() falls back to
// hybrid (header first, cookie second) when BACKEND_AUTH_MODE is unset.
func TestGenerateBearerAuth_HybridDefaultHybrid(t *testing.T) {
	dir := t.TempDir()
	if err := generateBearerAuth(dir, "example.com/proj", nil, "hybrid"); err != nil {
		t.Fatalf("generateBearerAuth: %v", err)
	}
	path := filepath.Join(dir, "backend", "internal", "middleware", "bearerauth.go")
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read emitted file: %v", err)
	}
	got := string(raw)
	if !strings.Contains(got, `return "hybrid"`) {
		t.Errorf("expected authMode fallback to return \"hybrid\", got:\n%s", got)
	}
}
