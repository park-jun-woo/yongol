//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestGenerateBodyLimit — body-limit 미들웨어 5파일 기록 success/error 검증

package middleware

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGenerateBodyLimit(t *testing.T) {
	t.Run("WritesFiles", func(t *testing.T) {
		arts := t.TempDir()
		if err := GenerateBodyLimit(arts); err != nil {
			t.Fatalf("GenerateBodyLimit error: %v", err)
		}
		mwDir := filepath.Join(arts, "backend", "internal", "middleware")
		for _, name := range []string{
			"body_limit.go", "multipart_limit.go", "override_body_limit.go",
			"apply_override.go", "respond_if_body_too_large.go",
		} {
			if _, err := os.Stat(filepath.Join(mwDir, name)); err != nil {
				t.Errorf("expected %s: %v", name, err)
			}
		}
	})

	t.Run("WriteError", func(t *testing.T) {
		arts := t.TempDir()
		// backend/internal is a file -> middleware MkdirAll (in writeFiles) fails.
		internal := filepath.Join(arts, "backend", "internal")
		if err := os.MkdirAll(filepath.Dir(internal), 0o755); err != nil {
			t.Fatalf("setup: %v", err)
		}
		if err := os.WriteFile(internal, []byte("x"), 0o644); err != nil {
			t.Fatalf("setup: %v", err)
		}
		if err := GenerateBodyLimit(arts); err == nil {
			t.Errorf("expected write_body_limit error, got nil")
		}
	})
}
