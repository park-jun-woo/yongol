//ff:func feature=openapi-parse type=test control=sequence
//ff:what BuildLineIndex — top-level scalar 는 빈 Paths 반환

package openapi

import (
	"os"
	"path/filepath"
	"testing"
)

func TestBuildLineIndex_ScalarRoot(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "scalar.yaml")
	// Top-level is a scalar, not a mapping → BuildLineIndex must return empty index.
	if err := os.WriteFile(path, []byte("hello\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	idx, err := BuildLineIndex(path)
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if len(idx.Paths) != 0 {
		t.Errorf("scalar root should yield empty Paths")
	}
}
