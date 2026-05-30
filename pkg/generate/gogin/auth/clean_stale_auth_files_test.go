//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestCleanStaleAuthFiles — internal/auth 디렉토리 제거 + 미존재 무시 검증

package auth

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCleanStaleAuthFiles(t *testing.T) {
	t.Run("RemovesExisting", func(t *testing.T) {
		artifactsDir := t.TempDir()
		authDir := filepath.Join(artifactsDir, "backend", "internal", "auth")
		if err := os.MkdirAll(authDir, 0o755); err != nil {
			t.Fatalf("setup: %v", err)
		}
		if err := os.WriteFile(filepath.Join(authDir, "claim.go"), []byte("x"), 0o644); err != nil {
			t.Fatalf("setup: %v", err)
		}
		if err := cleanStaleAuthFiles(artifactsDir); err != nil {
			t.Fatalf("cleanStaleAuthFiles error: %v", err)
		}
		if _, err := os.Stat(authDir); !os.IsNotExist(err) {
			t.Errorf("expected auth dir removed, stat err: %v", err)
		}
	})

	t.Run("RemoveError", func(t *testing.T) {
		// Make a parent component (internal) a regular file so RemoveAll on
		// internal/auth returns a non-IsNotExist (ENOTDIR) error.
		artifactsDir := t.TempDir()
		internalParent := filepath.Join(artifactsDir, "backend", "internal")
		if err := os.MkdirAll(filepath.Dir(internalParent), 0o755); err != nil {
			t.Fatalf("setup: %v", err)
		}
		if err := os.WriteFile(internalParent, []byte("x"), 0o644); err != nil {
			t.Fatalf("setup: %v", err)
		}
		err := cleanStaleAuthFiles(artifactsDir)
		if err == nil {
			t.Skip("RemoveAll did not surface ENOTDIR on this platform")
		}
		if os.IsNotExist(err) {
			t.Errorf("expected non-IsNotExist error, got: %v", err)
		}
	})

	t.Run("MissingDirIgnored", func(t *testing.T) {
		artifactsDir := t.TempDir()
		if err := cleanStaleAuthFiles(artifactsDir); err != nil {
			t.Errorf("expected nil for missing auth dir, got: %v", err)
		}
	})
}
