//ff:func feature=cli type=test control=sequence
//ff:what test: TestCountBackendFiles — arts/backend .go 파일 집계, vendor/hidden 제외

package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCountBackendFiles(t *testing.T) {
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
	missing := countBackendFiles(filepath.Join(arts, "no-such-dir"))
	if missing != 0 {
		t.Errorf("expected 0 for missing dir, got %d", missing)
	}
}
