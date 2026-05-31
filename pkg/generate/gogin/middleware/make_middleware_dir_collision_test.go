//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestMiddlewareExtras — csrf/prometheus/rate-limit/request-id/writeFiles 유틸 검증
package middleware

import (
	"os"
	"path/filepath"
	"testing"
)

func makeMiddlewareDirCollision(t *testing.T) string {
	t.Helper()
	arts := t.TempDir()
	internal := filepath.Join(arts, "backend", "internal")
	if err := os.MkdirAll(filepath.Dir(internal), 0o755); err != nil {
		t.Fatalf("setup: %v", err)
	}
	if err := os.WriteFile(internal, []byte("x"), 0o644); err != nil {
		t.Fatalf("setup: %v", err)
	}
	return arts
}
