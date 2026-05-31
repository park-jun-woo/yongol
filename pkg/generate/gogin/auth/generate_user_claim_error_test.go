//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestGenerateUserClaimError — model 디렉토리 생성 실패 시 user_claim 에러
package auth

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGenerateUserClaimError(t *testing.T) {
	dir := t.TempDir()
	// Make backend/internal a regular file so model dir MkdirAll fails.
	internal := filepath.Join(dir, "backend", "internal")
	if err := os.MkdirAll(filepath.Dir(internal), 0o755); err != nil {
		t.Fatalf("setup: %v", err)
	}
	if err := os.WriteFile(internal, []byte("x"), 0o644); err != nil {
		t.Fatalf("setup: %v", err)
	}
	err := Generate(authFSWithModule(), authClaimsState(), dir)
	if err == nil || !strings.Contains(err.Error(), "user_claim") {
		t.Errorf("expected user_claim error, got: %v", err)
	}
}
