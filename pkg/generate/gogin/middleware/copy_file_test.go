//ff:func feature=gen-gogin type=test control=branch topic=file-copy
//ff:what TestCopyFile — 성공 + mkdir/open/create 에러 분기 검증

package middleware

import (
	"os"
	"path/filepath"
	"strings"
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

func TestCopyFile_OpenError(t *testing.T) {
	dir := t.TempDir()
	src := filepath.Join(dir, "does-not-exist.txt")
	dst := filepath.Join(dir, "out.txt")
	err := copyFile(src, dst)
	if err == nil || !strings.Contains(err.Error(), "open") {
		t.Errorf("expected open error, got: %v", err)
	}
}

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
