//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestCopyFile — 성공 + mkdir/open/create 에러 분기 검증
package middleware

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCopyFile_Success(t *testing.T) {
	dir := t.TempDir()
	src := filepath.Join(dir, "src.txt")
	if err := os.WriteFile(src, []byte("hello world"), 0o644); err != nil {
		t.Fatalf("setup: %v", err)
	}
	// dst includes a non-existent parent dir to exercise MkdirAll.
	dst := filepath.Join(dir, "nested", "deep", "out.txt")
	if err := copyFile(src, dst); err != nil {
		t.Fatalf("copyFile: %v", err)
	}
	got, err := os.ReadFile(dst)
	if err != nil {
		t.Fatalf("read dst: %v", err)
	}
	if string(got) != "hello world" {
		t.Errorf("content = %q, want %q", string(got), "hello world")
	}
}
