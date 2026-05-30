//ff:func feature=gen-gogin type=test control=branch topic=go-mod
//ff:what TestRemoveStaleGoModFiles — 기존 파일 제거 / 부재 무시 / 제거 실패 에러 분기 검증

package gogin

import (
	"os"
	"path/filepath"
	"strings"
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

func TestRemoveStaleGoModFiles_NoFiles(t *testing.T) {
	// Empty dir -> os.Remove returns IsNotExist which is tolerated.
	if err := removeStaleGoModFiles(t.TempDir()); err != nil {
		t.Errorf("expected nil when no stale files, got: %v", err)
	}
}

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
