//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestGenerateBearerAuthErrors — mkdir/writefile 에러 경로 검증

package auth

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGenerateBearerAuthErrors(t *testing.T) {
	t.Run("MkdirFails", func(t *testing.T) {
		dir := t.TempDir()
		// backend/internal/middleware parent (internal) is a regular file.
		internal := filepath.Join(dir, "backend", "internal")
		if err := os.MkdirAll(filepath.Dir(internal), 0o755); err != nil {
			t.Fatalf("setup: %v", err)
		}
		if err := os.WriteFile(internal, []byte("x"), 0o644); err != nil {
			t.Fatalf("setup: %v", err)
		}
		if err := generateBearerAuth(dir, "github.com/test/app", nil, "bearer"); err == nil {
			t.Errorf("expected MkdirAll error, got nil")
		}
	})

	t.Run("WriteFileFails", func(t *testing.T) {
		dir := t.TempDir()
		mwDir := filepath.Join(dir, "backend", "internal", "middleware")
		if err := os.MkdirAll(mwDir, 0o755); err != nil {
			t.Fatalf("setup: %v", err)
		}
		// Pre-create one of the target files as a directory so WriteFile fails.
		for _, name := range []string{"auth_mode.go", "extract_token.go", "bearerauth.go"} {
			if err := os.MkdirAll(filepath.Join(mwDir, name), 0o755); err != nil {
				t.Fatalf("setup: %v", err)
			}
		}
		err := generateBearerAuth(dir, "github.com/test/app", nil, "bearer")
		if err == nil || !strings.Contains(err.Error(), "write ") {
			t.Errorf("expected write error, got: %v", err)
		}
	})
}
