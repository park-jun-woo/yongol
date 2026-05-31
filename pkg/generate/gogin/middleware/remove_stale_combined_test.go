//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestMiddlewareExtras — csrf/prometheus/rate-limit/request-id/writeFiles 유틸 검증
package middleware

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRemoveStaleCombined(t *testing.T) {
	t.Run("RemovesAndIgnoresMissing", func(t *testing.T) {
		dir := t.TempDir()
		// create one stale file; the others are missing (ignored).
		if err := os.WriteFile(filepath.Join(dir, "prometheus.go"), []byte("x"), 0o644); err != nil {
			t.Fatalf("setup: %v", err)
		}
		if err := removeStaleCombined(dir); err != nil {
			t.Fatalf("error: %v", err)
		}
		if _, err := os.Stat(filepath.Join(dir, "prometheus.go")); !os.IsNotExist(err) {
			t.Errorf("expected prometheus.go removed")
		}
	})

	t.Run("RemoveError", func(t *testing.T) {
		dir := t.TempDir()
		if err := os.WriteFile(filepath.Join(dir, "prometheus.go"), []byte("x"), 0o644); err != nil {
			t.Fatalf("setup: %v", err)
		}
		if err := os.Chmod(dir, 0o555); err != nil {
			t.Fatalf("setup: %v", err)
		}
		defer os.Chmod(dir, 0o755)
		if err := removeStaleCombined(dir); err == nil {
			t.Skip("Remove did not fail (likely root)")
		}
	})
}
