//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestCopyFile — 성공 + mkdir/open/create 에러 분기 검증
package middleware

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCopyFile_CreateError(t *testing.T) {
	dir := t.TempDir()
	src := filepath.Join(dir, "src.txt")
	if err := os.WriteFile(src, []byte("data"), 0o644); err != nil {
		t.Fatalf("setup: %v", err)
	}
	// dst is an existing directory -> os.Create fails.
	dst := filepath.Join(dir, "adir")
	if err := os.MkdirAll(dst, 0o755); err != nil {
		t.Fatalf("setup: %v", err)
	}
	err := copyFile(src, dst)
	if err == nil || !strings.Contains(err.Error(), "create") {
		t.Errorf("expected create error, got: %v", err)
	}
}
