//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestCopyFile — 파일 복사 success + mkdir/open/create 에러 경로 검증

package gogin

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCopyFile(t *testing.T) {
	t.Run("CopiesContent", func(t *testing.T) {
		dir := t.TempDir()
		src := filepath.Join(dir, "src.txt")
		mustWrite(t, src, "hello world")
		dst := filepath.Join(dir, "nested", "dst.txt")
		if err := copyFile(src, dst); err != nil {
			t.Fatalf("copyFile error: %v", err)
		}
		got, err := os.ReadFile(dst)
		if err != nil {
			t.Fatalf("read dst: %v", err)
		}
		if string(got) != "hello world" {
			t.Errorf("content mismatch: %q", got)
		}
	})

	t.Run("MkdirError", func(t *testing.T) {
		dir := t.TempDir()
		// Parent of dst is a regular file -> MkdirAll fails.
		blocker := filepath.Join(dir, "blocker")
		mustWrite(t, blocker, "x")
		err := copyFile(filepath.Join(dir, "src.txt"), filepath.Join(blocker, "sub", "dst.txt"))
		if err == nil {
			t.Errorf("expected mkdir error, got nil")
		}
	})

	t.Run("OpenSrcError", func(t *testing.T) {
		dir := t.TempDir()
		err := copyFile(filepath.Join(dir, "missing.txt"), filepath.Join(dir, "dst.txt"))
		if err == nil {
			t.Errorf("expected open error for missing src, got nil")
		}
	})

	t.Run("CreateDstError", func(t *testing.T) {
		dir := t.TempDir()
		src := filepath.Join(dir, "src.txt")
		mustWrite(t, src, "x")
		// dst path is an existing directory -> os.Create fails.
		dst := filepath.Join(dir, "out")
		if err := os.MkdirAll(dst, 0o755); err != nil {
			t.Fatalf("setup: %v", err)
		}
		if err := copyFile(src, dst); err == nil {
			t.Errorf("expected create error when dst is a dir, got nil")
		}
	})
}
