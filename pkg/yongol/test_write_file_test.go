//ff:func feature=orchestrator type=test-helper control=sequence
//ff:what writeFile — 테스트 전용 MkdirAll + WriteFile wrapper (에러는 t.Fatalf)
package yongol

import (
	"os"
	"path/filepath"
	"testing"
)

// writeFile is a tiny wrapper that fails the test on write error; keeps
// individual test bodies focused on the matrix they care about.
func writeFile(t *testing.T, path, body string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("mkdir %s: %v", filepath.Dir(path), err)
	}
	if err := os.WriteFile(path, []byte(body), 0o644); err != nil {
		t.Fatalf("write %s: %v", path, err)
	}
}
