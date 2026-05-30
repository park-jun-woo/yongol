//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestWriteFuncStubs — 외부 패키지 stub Python 모듈 기록 검증

package fastapi

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestWriteFuncStubs(t *testing.T) {
	t.Run("WritesStubPerPackage", func(t *testing.T) {
		appDir := t.TempDir()
		pkgs := []externalPackage{
			{Name: "mailer", Methods: []string{"send"}},
		}
		if err := writeFuncStubs(appDir, pkgs); err != nil {
			t.Fatalf("writeFuncStubs error: %v", err)
		}
		if _, err := os.Stat(filepath.Join(appDir, "services", "mailer.py")); err != nil {
			t.Errorf("expected services/mailer.py: %v", err)
		}
	})

	t.Run("MkdirServicesFails", func(t *testing.T) {
		appDir := t.TempDir()
		if err := os.WriteFile(filepath.Join(appDir, "services"), []byte("x"), 0o644); err != nil {
			t.Fatalf("setup: %v", err)
		}
		err := writeFuncStubs(appDir, nil)
		if err == nil || !strings.Contains(err.Error(), "mkdir services") {
			t.Errorf("expected mkdir services error, got: %v", err)
		}
	})

	t.Run("WriteStubFails", func(t *testing.T) {
		// services dir is created, but the stub file path collides with a dir.
		appDir := t.TempDir()
		servicesDir := filepath.Join(appDir, "services")
		if err := os.MkdirAll(filepath.Join(servicesDir, "mailer.py"), 0o755); err != nil {
			t.Fatalf("setup: %v", err)
		}
		pkgs := []externalPackage{{Name: "mailer", Methods: []string{"send"}}}
		err := writeFuncStubs(appDir, pkgs)
		if err == nil || !strings.Contains(err.Error(), "write mailer stub") {
			t.Errorf("expected write stub error, got: %v", err)
		}
	})
}
