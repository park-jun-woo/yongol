//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestRemoveStaleGoModFiles — 기존 파일 제거 / 부재 무시 / 제거 실패 에러 분기 검증
package gogin

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRemoveStaleGoModFiles_RemoveError(t *testing.T) {
	dir := t.TempDir()
	// go.mod is a non-empty directory -> os.Remove fails (not IsNotExist).
	goModDir := filepath.Join(dir, "go.mod")
	if err := os.MkdirAll(filepath.Join(goModDir, "child"), 0o755); err != nil {
		t.Fatalf("setup: %v", err)
	}
	err := removeStaleGoModFiles(dir)
	if err == nil || !strings.Contains(err.Error(), "remove stale go.mod") {
		t.Errorf("expected remove stale go.mod error, got: %v", err)
	}
}
