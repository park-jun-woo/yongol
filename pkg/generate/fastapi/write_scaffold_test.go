//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestWriteScaffold — FastAPI scaffold 파일 일괄 기록 검증

package fastapi

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteScaffold(t *testing.T) {
	t.Run("WritesAllScaffoldFiles", func(t *testing.T) {
		backendDir := filepath.Join(t.TempDir(), "backend")
		if err := writeScaffold(backendDir, "myproject"); err != nil {
			t.Fatalf("writeScaffold error: %v", err)
		}
		for _, name := range []string{"pyproject.toml", "requirements.txt", ".env.example"} {
			info, err := os.Stat(filepath.Join(backendDir, name))
			if err != nil {
				t.Errorf("expected %s to exist: %v", name, err)
				continue
			}
			if info.Size() == 0 {
				t.Errorf("%s is empty", name)
			}
		}
	})

	t.Run("MkdirFails", func(t *testing.T) {
		// backendDir parent path is a regular file -> MkdirAll fails.
		base := t.TempDir()
		filePath := filepath.Join(base, "file")
		if err := os.WriteFile(filePath, []byte("x"), 0o644); err != nil {
			t.Fatalf("setup: %v", err)
		}
		if err := writeScaffold(filepath.Join(filePath, "backend"), "myproject"); err == nil {
			t.Errorf("expected MkdirAll error, got nil")
		}
	})

	t.Run("RenderPyprojectFailsOnEmptyProjectID", func(t *testing.T) {
		backendDir := filepath.Join(t.TempDir(), "backend")
		if err := writeScaffold(backendDir, ""); err == nil {
			t.Errorf("expected RenderPyproject error for empty projectID, got nil")
		}
	})

	t.Run("WritePyprojectFails", func(t *testing.T) {
		// backendDir exists; pyproject.toml path collides with a directory.
		backendDir := t.TempDir()
		if err := os.MkdirAll(filepath.Join(backendDir, "pyproject.toml"), 0o755); err != nil {
			t.Fatalf("setup: %v", err)
		}
		if err := writeScaffold(backendDir, "myproject"); err == nil {
			t.Errorf("expected write pyproject.toml error, got nil")
		}
	})
}
