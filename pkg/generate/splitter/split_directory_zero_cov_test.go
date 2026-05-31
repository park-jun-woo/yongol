//ff:func feature=gen-splitter type=test control=sequence
//ff:what zz_zerocov_test — splitter 패키지의 0% 커버리지 함수(cleanOriginal/preserveComments/isPreservedFile/SplitDirectory) 단위 테스트
package splitter

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSplitDirectory_ZeroCov(t *testing.T) {
	dir := t.TempDir()
	// A real sqlc-style source file to split, plus a preserved querier.go,
	// plus a nested dir to exercise the IsDir skip.
	writeFileZeroCov(t, dir, "models.go", "package db\n\ntype User struct {\n\tID int64\n}\n")
	writeFileZeroCov(t, dir, "querier.go", "package db\n\ntype Querier interface{}\n")
	if err := os.Mkdir(filepath.Join(dir, "sub"), 0o755); err != nil {
		t.Fatal(err)
	}

	if err := SplitDirectory(dir, ToolSQLC); err != nil {
		t.Fatalf("SplitDirectory: %v", err)
	}

	// models.go should be split away (removed) and querier.go preserved.
	if _, err := os.Stat(filepath.Join(dir, "models.go")); !os.IsNotExist(err) {
		t.Error("models.go should have been split and removed")
	}
	if _, err := os.Stat(filepath.Join(dir, "querier.go")); err != nil {
		t.Errorf("querier.go should be preserved: %v", err)
	}
}
