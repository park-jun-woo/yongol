//ff:func feature=gen-filefunc type=test control=sequence
//ff:what TestInsertDirEntry — 디렉토리만 feature 맵에 삽입, 파일/중복 제외 검증

package filefunc

import (
	"os"
	"path/filepath"
	"testing"
)

func TestInsertDirEntry(t *testing.T) {
	root := t.TempDir()
	if err := os.MkdirAll(filepath.Join(root, "auth"), 0o755); err != nil {
		t.Fatalf("setup: %v", err)
	}
	if err := os.MkdirAll(filepath.Join(root, "existing"), 0o755); err != nil {
		t.Fatalf("setup: %v", err)
	}
	if err := os.WriteFile(filepath.Join(root, "file.txt"), []byte("x"), 0o644); err != nil {
		t.Fatalf("setup: %v", err)
	}
	entries, err := os.ReadDir(root)
	if err != nil {
		t.Fatalf("ReadDir: %v", err)
	}

	dst := map[string]string{"existing": "kept"}
	for _, e := range entries {
		insertDirEntry(dst, e)
	}

	// directory with infra description -> resolved.
	if got := dst["auth"]; got != "JWT issuance and verification" {
		t.Errorf("auth: expected infra description, got %q", got)
	}
	// existing key preserved.
	if got := dst["existing"]; got != "kept" {
		t.Errorf("existing: expected preserved value, got %q", got)
	}
	// file skipped.
	if _, ok := dst["file.txt"]; ok {
		t.Errorf("file should not be inserted")
	}
}
