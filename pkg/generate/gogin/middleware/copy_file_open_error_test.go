//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestCopyFile — 성공 + mkdir/open/create 에러 분기 검증
package middleware

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestCopyFile_OpenError(t *testing.T) {
	dir := t.TempDir()
	src := filepath.Join(dir, "does-not-exist.txt")
	dst := filepath.Join(dir, "out.txt")
	err := copyFile(src, dst)
	if err == nil || !strings.Contains(err.Error(), "open") {
		t.Errorf("expected open error, got: %v", err)
	}
}
