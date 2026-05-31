//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestMiddlewareExtras — csrf/prometheus/rate-limit/request-id/writeFiles 유틸 검증
package middleware

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGenerateRequestID(t *testing.T) {
	t.Run("WritesFiles", func(t *testing.T) {
		arts := t.TempDir()
		if err := GenerateRequestID(arts); err != nil {
			t.Fatalf("error: %v", err)
		}
		if _, err := os.Stat(filepath.Join(arts, "backend", "internal", "middleware", "request_id.go")); err != nil {
			t.Errorf("expected request_id.go: %v", err)
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
		if err := GenerateRequestID(arts); err == nil {
			t.Errorf("expected write error, got nil")
		}
	})
}
