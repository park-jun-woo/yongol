//ff:func feature=orchestrator type=test control=sequence
//ff:what TestIsDir/dirPresence/isYongolRoot — 디렉토리 판별·presence 매핑·루트 판별 검증
package yongol

import (
	"os"
	"path/filepath"
	"testing"
)

func TestIsDir(t *testing.T) {
	dir := t.TempDir()
	if !isDir(dir) {
		t.Error("expected isDir(tempdir) = true")
	}

	file := filepath.Join(dir, "f.txt")
	if err := os.WriteFile(file, []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	if isDir(file) {
		t.Error("expected isDir(file) = false")
	}
	if isDir(filepath.Join(dir, "missing")) {
		t.Error("expected isDir(missing) = false")
	}
}
