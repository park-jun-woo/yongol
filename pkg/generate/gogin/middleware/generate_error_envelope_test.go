//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestGenerateErrorEnvelope — error envelope 미들웨어 파일 기록 success/error 검증

package middleware

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGenerateErrorEnvelope(t *testing.T) {
	t.Run("WritesFiles", func(t *testing.T) {
		arts := t.TempDir()
		if err := GenerateErrorEnvelope(arts); err != nil {
			t.Fatalf("GenerateErrorEnvelope error: %v", err)
		}
		mwDir := filepath.Join(arts, "backend", "internal", "middleware")
		for _, name := range []string{
			"error_envelope.go", "default_code_for.go", "default_message_for.go",
			"write_envelope.go", "write_envelope_with_context.go", "error_envelope_middleware.go",
		} {
			if _, err := os.Stat(filepath.Join(mwDir, name)); err != nil {
				t.Errorf("expected %s: %v", name, err)
			}
		}
	})

	t.Run("WriteError", func(t *testing.T) {
		arts := t.TempDir()
		internal := filepath.Join(arts, "backend", "internal")
		if err := os.MkdirAll(filepath.Dir(internal), 0o755); err != nil {
			t.Fatalf("setup: %v", err)
		}
		if err := os.WriteFile(internal, []byte("x"), 0o644); err != nil {
			t.Fatalf("setup: %v", err)
		}
		if err := GenerateErrorEnvelope(arts); err == nil {
			t.Errorf("expected error, got nil")
		}
	})
}
