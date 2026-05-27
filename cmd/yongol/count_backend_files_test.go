//ff:func feature=cli type=test control=sequence
//ff:what countBackendFiles test — .go 파일 카운트, vendor/hidden 제외 검증

package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCountBackendFiles(t *testing.T) {
	t.Run("Normal", func(t *testing.T) {
		arts := t.TempDir()
		backend := filepath.Join(arts, "backend")
		if err := os.MkdirAll(filepath.Join(backend, "internal", "service"), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.MkdirAll(filepath.Join(backend, "vendor", "foo"), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.MkdirAll(filepath.Join(backend, ".hidden"), 0o755); err != nil {
			t.Fatal(err)
		}
		mustWriteEmpty(t, filepath.Join(backend, "main.go"))
		mustWriteEmpty(t, filepath.Join(backend, "internal", "service", "a.go"))
		mustWriteEmpty(t, filepath.Join(backend, "internal", "service", "b.go"))
		mustWriteEmpty(t, filepath.Join(backend, "vendor", "foo", "vendor.go"))
		mustWriteEmpty(t, filepath.Join(backend, ".hidden", "hidden.go"))
		mustWriteEmpty(t, filepath.Join(backend, "README.md"))
		got := countBackendFiles(arts)
		if got != 3 {
			t.Fatalf("expected 3 .go files, got %d", got)
		}
	})
	t.Run("MissingDir", func(t *testing.T) {
		got := countBackendFiles("/tmp/this-dir-should-not-exist-yongol-count")
		if got != 0 {
			t.Errorf("expected 0 for missing dir, got %d", got)
		}
	})
	t.Run("NodeModules", func(t *testing.T) {
		arts := t.TempDir()
		backend := filepath.Join(arts, "backend")
		if err := os.MkdirAll(filepath.Join(backend, "node_modules", "pkg"), 0o755); err != nil {
			t.Fatal(err)
		}
		mustWriteEmpty(t, filepath.Join(backend, "node_modules", "pkg", "skip.go"))
		mustWriteEmpty(t, filepath.Join(backend, "app.go"))
		got := countBackendFiles(arts)
		if got != 1 {
			t.Errorf("expected 1, got %d", got)
		}
	})
}
