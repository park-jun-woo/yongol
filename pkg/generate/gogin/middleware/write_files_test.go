//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestMiddlewareExtras — csrf/prometheus/rate-limit/request-id/writeFiles 유틸 검증
package middleware

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteFiles(t *testing.T) {
	t.Run("WritesGoAndPlain", func(t *testing.T) {
		dir := filepath.Join(t.TempDir(), "out")
		files := map[string]string{
			"a.go":  "package x\n\nfunc A() {}\n",
			"b.txt": "plain",
		}
		if err := writeFiles(dir, files); err != nil {
			t.Fatalf("error: %v", err)
		}
		if _, err := os.Stat(filepath.Join(dir, "a.go")); err != nil {
			t.Errorf("expected a.go: %v", err)
		}
		if _, err := os.Stat(filepath.Join(dir, "b.txt")); err != nil {
			t.Errorf("expected b.txt: %v", err)
		}
	})

	t.Run("MkdirError", func(t *testing.T) {
		base := t.TempDir()
		fp := filepath.Join(base, "file")
		if err := os.WriteFile(fp, []byte("x"), 0o644); err != nil {
			t.Fatalf("setup: %v", err)
		}
		if err := writeFiles(filepath.Join(fp, "sub"), map[string]string{"a.go": "package x"}); err == nil {
			t.Errorf("expected mkdir error, got nil")
		}
	})

	t.Run("WriteFileError", func(t *testing.T) {
		dir := t.TempDir()
		// target file name pre-exists as a directory -> WriteIfNotPreserved fails.
		if err := os.MkdirAll(filepath.Join(dir, "a.go"), 0o755); err != nil {
			t.Fatalf("setup: %v", err)
		}
		if err := writeFiles(dir, map[string]string{"a.go": "package x\n\nfunc A() {}\n"}); err == nil {
			t.Errorf("expected write error, got nil")
		}
	})
}
