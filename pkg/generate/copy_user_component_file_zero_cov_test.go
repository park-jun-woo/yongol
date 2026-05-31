//ff:func feature=generate type=test control=sequence
//ff:what TestZeroCov — 0% util 함수 (isCopiedExtension / isYongolManaged / mergeFieldlessOps / ResolveBackendType / WithMigration / appendChildNodeFormActions) 회귀
package generate

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCopyUserComponentFile_ZeroCov(t *testing.T) {
	dir := t.TempDir()
	src := filepath.Join(dir, "src.txt")
	dst := filepath.Join(dir, "out", "dst.txt")
	if err := os.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(src, []byte("hello"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := copyUserComponentFile(src, dst); err != nil {
		t.Fatalf("copy: %v", err)
	}
	got, err := os.ReadFile(dst)
	if err != nil || string(got) != "hello" {
		t.Errorf("dst content = %q err=%v", got, err)
	}

	// open error: missing src.
	if err := copyUserComponentFile(filepath.Join(dir, "nope"), dst); err == nil {
		t.Error("expected open error")
	}
	// create error: dst dir does not exist.
	if err := copyUserComponentFile(src, filepath.Join(dir, "missing-dir", "x.txt")); err == nil {
		t.Error("expected create error")
	}
}
