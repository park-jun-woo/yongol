//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestGenerateBearerAuthError — middleware 경로 충돌 시 bearer_auth 에러
package auth

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGenerateBearerAuthError(t *testing.T) {
	dir := t.TempDir()
	// model dir creatable, but middleware path collides with a file so
	// generateBearerAuth's MkdirAll fails after user_claim succeeds.
	mwParent := filepath.Join(dir, "backend", "internal")
	if err := os.MkdirAll(mwParent, 0o755); err != nil {
		t.Fatalf("setup: %v", err)
	}
	if err := os.WriteFile(filepath.Join(mwParent, "middleware"), []byte("x"), 0o644); err != nil {
		t.Fatalf("setup: %v", err)
	}
	err := Generate(authFSWithModule(), authClaimsState(), dir)
	if err == nil || !strings.Contains(err.Error(), "bearer_auth") {
		t.Errorf("expected bearer_auth error, got: %v", err)
	}
}
