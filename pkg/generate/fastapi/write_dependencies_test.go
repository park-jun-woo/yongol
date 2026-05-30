//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestWriteDependencies — FastAPI 의존성 주입 모듈 파일 기록 검증

package fastapi

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteDependencies(t *testing.T) {
	t.Run("WritesDependencyModules", func(t *testing.T) {
		appDir := t.TempDir()
		if err := writeDependencies(appDir); err != nil {
			t.Fatalf("writeDependencies error: %v", err)
		}
		for _, name := range []string{"__init__.py", "database.py", "auth.py"} {
			if _, err := os.Stat(filepath.Join(appDir, "dependencies", name)); err != nil {
				t.Errorf("expected dependencies/%s: %v", name, err)
			}
		}
	})

	t.Run("MkdirFails", func(t *testing.T) {
		base := t.TempDir()
		filePath := filepath.Join(base, "file")
		if err := os.WriteFile(filePath, []byte("x"), 0o644); err != nil {
			t.Fatalf("setup: %v", err)
		}
		if err := writeDependencies(filePath); err == nil {
			t.Errorf("expected MkdirAll error, got nil")
		}
	})
}
