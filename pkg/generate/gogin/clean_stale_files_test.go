//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestCleanStaleFiles — stale .go 제거 + keep/ext/subdir/missing/error 분기 검증

package gogin

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCleanStaleFiles(t *testing.T) {
	t.Run("RemovesStaleKeepsRest", func(t *testing.T) {
		dir := t.TempDir()
		mustWrite(t, filepath.Join(dir, "keep.go"), "x")
		mustWrite(t, filepath.Join(dir, "stale.go"), "x")
		mustWrite(t, filepath.Join(dir, "README.md"), "x") // non-.go -> untouched
		if err := os.MkdirAll(filepath.Join(dir, "sub.go"), 0o755); err != nil {
			t.Fatalf("setup: %v", err) // a directory named *.go -> ignored
		}

		keep := map[string]bool{"keep.go": true}
		if err := CleanStaleFiles(dir, ".go", keep); err != nil {
			t.Fatalf("CleanStaleFiles error: %v", err)
		}
		if _, err := os.Stat(filepath.Join(dir, "keep.go")); err != nil {
			t.Errorf("keep.go should survive: %v", err)
		}
		if _, err := os.Stat(filepath.Join(dir, "stale.go")); !os.IsNotExist(err) {
			t.Errorf("stale.go should be removed, stat err: %v", err)
		}
		if _, err := os.Stat(filepath.Join(dir, "README.md")); err != nil {
			t.Errorf("README.md should survive: %v", err)
		}
		if _, err := os.Stat(filepath.Join(dir, "sub.go")); err != nil {
			t.Errorf("directory sub.go should survive: %v", err)
		}
	})

	t.Run("MissingDirIsNil", func(t *testing.T) {
		if err := CleanStaleFiles(filepath.Join(t.TempDir(), "nope"), ".go", nil); err != nil {
			t.Errorf("expected nil for missing dir, got: %v", err)
		}
	})

	t.Run("ReadDirNonNotExistError", func(t *testing.T) {
		// dir is a regular file -> ReadDir returns ENOTDIR (not IsNotExist).
		base := t.TempDir()
		fp := filepath.Join(base, "afile")
		mustWrite(t, fp, "x")
		if err := CleanStaleFiles(fp, ".go", nil); err == nil {
			t.Errorf("expected ReadDir error for non-dir path, got nil")
		}
	})

	t.Run("RemoveError", func(t *testing.T) {
		dir := t.TempDir()
		mustWrite(t, filepath.Join(dir, "stale.go"), "x")
		// Make the dir read-only so Remove fails (EACCES).
		if err := os.Chmod(dir, 0o555); err != nil {
			t.Fatalf("setup: %v", err)
		}
		defer os.Chmod(dir, 0o755)
		err := CleanStaleFiles(dir, ".go", nil)
		if err == nil {
			t.Skip("Remove did not fail (likely running as root)")
		}
	})
}

//ff:func feature=gen-gogin type=test-helper control=sequence
//ff:what mustWrite — 테스트용 파일 기록 헬퍼
func mustWrite(t *testing.T, path, content string) {
	t.Helper()
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("write %s: %v", path, err)
	}
}
