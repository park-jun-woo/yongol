//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestWriteInitFiles — Python 패키지 __init__.py 일괄 기록 검증

package fastapi

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteInitFiles(t *testing.T) {
	t.Run("WritesAllInitFiles", func(t *testing.T) {
		appDir := t.TempDir()
		if err := writeInitFiles(appDir, []string{"workflow"}); err != nil {
			t.Fatalf("writeInitFiles error: %v", err)
		}
		paths := []string{
			"__init__.py",
			filepath.Join("services", "__init__.py"),
			filepath.Join("routers", "__init__.py"),
			filepath.Join("schemas", "__init__.py"),
		}
		for _, p := range paths {
			if _, err := os.Stat(filepath.Join(appDir, p)); err != nil {
				t.Errorf("expected %s: %v", p, err)
			}
		}
	})

	t.Run("AppInitFails", func(t *testing.T) {
		// appDir itself is a regular file -> WriteFile(app/__init__.py) fails.
		base := t.TempDir()
		filePath := filepath.Join(base, "file")
		if err := os.WriteFile(filePath, []byte("x"), 0o644); err != nil {
			t.Fatalf("setup: %v", err)
		}
		if err := writeInitFiles(filePath, nil); err == nil {
			t.Errorf("expected write error, got nil")
		}
	})

	t.Run("MkdirSubdirFails", func(t *testing.T) {
		// app/__init__.py writes ok, but services path collides with a file.
		appDir := t.TempDir()
		if err := os.WriteFile(filepath.Join(appDir, "services"), []byte("x"), 0o644); err != nil {
			t.Fatalf("setup: %v", err)
		}
		if err := writeInitFiles(appDir, nil); err == nil {
			t.Errorf("expected MkdirAll error, got nil")
		}
	})
}
