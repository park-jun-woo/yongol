//ff:func feature=cli type=test-helper control=iteration dimension=1
//ff:what repoRoot — 테스트 파일에서 위로 go.mod 찾아 yongol 모듈 root 반환

package main

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

// repoRoot walks up from this test file to the yongol module root. The
// anchor is go.mod — every caller expects to see it at the top of the repo.
func repoRoot(t *testing.T) string {
	t.Helper()
	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("runtime.Caller(0) failed")
	}
	dir := filepath.Dir(thisFile)
	for i := 0; i < 10; i++ {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	t.Fatal("repo root (go.mod) not found by walking up from test file")
	return ""
}
