//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what TestRemoveStaleGoModFiles — 기존 파일 제거 / 부재 무시 / 제거 실패 에러 분기 검증
package gogin

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRemoveStaleGoModFiles_RemovesExisting(t *testing.T) {
	dir := t.TempDir()
	for _, n := range []string{"go.mod", "go.sum"} {
		if err := os.WriteFile(filepath.Join(dir, n), []byte("x"), 0o644); err != nil {
			t.Fatalf("setup: %v", err)
		}
	}
	if err := removeStaleGoModFiles(dir); err != nil {
		t.Fatalf("removeStaleGoModFiles: %v", err)
	}
	for _, n := range []string{"go.mod", "go.sum"} {
		if _, err := os.Stat(filepath.Join(dir, n)); !os.IsNotExist(err) {
			t.Errorf("expected %s removed, stat err = %v", n, err)
		}
	}
}
