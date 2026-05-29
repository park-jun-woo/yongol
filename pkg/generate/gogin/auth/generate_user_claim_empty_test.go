//ff:func feature=gen-gogin type=test control=sequence
//ff:what test: TestGenerateUserClaimEmptyFields — 빈 claim list 시 파일 미생성 확인

package auth

import (
	"os"
	"path/filepath"
	"testing"
)

// TestGenerateUserClaimEmptyFields verifies that an empty claim list is a
// no-op: no file is written. Aligns with the guard in auth.Generate which
// skips claim emission entirely when manifest.backend.auth.claims is empty.
func TestGenerateUserClaimEmptyFields(t *testing.T) {
	dir := t.TempDir()
	if err := generateUserClaim(dir, nil); err != nil {
		t.Fatalf("generateUserClaim empty: %v", err)
	}
	path := filepath.Join(dir, "backend", "internal", "model", "user_claim.go")
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		t.Fatalf("expected no file for empty fields, stat err=%v", err)
	}
}
