//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestWriteAdapterFile — postgres.go adapter 기록 success + mkdir 에러 검증

package infra

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteAdapterFile(t *testing.T) {
	t.Run("WritesPostgresGo", func(t *testing.T) {
		arts := t.TempDir()
		if err := writeAdapterFile(arts, "cache", []byte("package cache\n")); err != nil {
			t.Fatalf("writeAdapterFile error: %v", err)
		}
		path := filepath.Join(arts, "backend", "internal", "infra", "cache", "postgres.go")
		raw, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("read: %v", err)
		}
		if string(raw) != "package cache\n" {
			t.Errorf("content mismatch: %q", raw)
		}
	})

	t.Run("MkdirError", func(t *testing.T) {
		arts := t.TempDir()
		// backend is a regular file -> MkdirAll fails.
		if err := os.WriteFile(filepath.Join(arts, "backend"), []byte("x"), 0o644); err != nil {
			t.Fatalf("setup: %v", err)
		}
		if err := writeAdapterFile(arts, "cache", []byte("x")); err == nil {
			t.Errorf("expected MkdirAll error, got nil")
		}
	})
}
