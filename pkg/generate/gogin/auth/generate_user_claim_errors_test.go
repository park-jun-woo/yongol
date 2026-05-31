//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestGenerateUserClaimUUIDAndErrors — uuid 변환 + mkdir/writefile 에러 경로 검증
package auth

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGenerateUserClaimErrors(t *testing.T) {
	fields := []ClaimField{{Name: "ID", Key: "user_id", GoType: "int64"}}

	t.Run("MkdirFails", func(t *testing.T) {
		dir := t.TempDir()
		internal := filepath.Join(dir, "backend", "internal")
		if err := os.MkdirAll(filepath.Dir(internal), 0o755); err != nil {
			t.Fatalf("setup: %v", err)
		}
		if err := os.WriteFile(internal, []byte("x"), 0o644); err != nil {
			t.Fatalf("setup: %v", err)
		}
		if err := generateUserClaim(dir, fields); err == nil {
			t.Errorf("expected MkdirAll error, got nil")
		}
	})

	t.Run("WriteFileFails", func(t *testing.T) {
		dir := t.TempDir()
		modelDir := filepath.Join(dir, "backend", "internal", "model")
		if err := os.MkdirAll(filepath.Join(modelDir, "user_claim.go"), 0o755); err != nil {
			t.Fatalf("setup: %v", err)
		}
		if err := generateUserClaim(dir, fields); err == nil {
			t.Errorf("expected WriteFile error, got nil")
		}
	})
}
