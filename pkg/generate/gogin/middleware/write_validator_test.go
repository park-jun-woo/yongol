//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestMiddlewareExtras — csrf/prometheus/rate-limit/request-id/writeFiles 유틸 검증
package middleware

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteValidator(t *testing.T) {
	t.Run("Writes", func(t *testing.T) {
		dir := t.TempDir()
		if err := writeValidator(dir); err != nil {
			t.Fatalf("error: %v", err)
		}
		if _, err := os.Stat(filepath.Join(dir, "request_validator.go")); err != nil {
			t.Errorf("expected request_validator.go: %v", err)
		}
	})
	t.Run("WriteError", func(t *testing.T) {
		dir := t.TempDir()
		// target is a directory -> WriteFile fails.
		if err := os.MkdirAll(filepath.Join(dir, "request_validator.go"), 0o755); err != nil {
			t.Fatalf("setup: %v", err)
		}
		if err := writeValidator(dir); err == nil {
			t.Errorf("expected write error, got nil")
		}
	})
}
