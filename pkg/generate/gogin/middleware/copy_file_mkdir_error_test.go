//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestCopyFile — 성공 + mkdir/open/create 에러 분기 검증
package middleware

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCopyFile_MkdirError(t *testing.T) {
	dir := t.TempDir()
	// A regular file occupies the path that would need to be a parent dir.
	blocker := filepath.Join(dir, "blocker")
	if err := os.WriteFile(blocker, []byte("x"), 0o644); err != nil {
		t.Fatalf("setup: %v", err)
	}
	dst := filepath.Join(blocker, "child", "out.txt")
	err := copyFile(filepath.Join(dir, "src.txt"), dst)
	if err == nil || !strings.Contains(err.Error(), "mkdir") {
		t.Errorf("expected mkdir error, got: %v", err)
	}
}
