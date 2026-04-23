//ff:func feature=openapi-parse type=test control=sequence
//ff:what BuildLineIndex — 빈 yaml 은 err 없이 빈 인덱스 반환

package openapi

import (
	"os"
	"path/filepath"
	"testing"
)

func TestBuildLineIndex_EmptyYaml(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "empty.yaml")
	if err := os.WriteFile(path, []byte(""), 0o644); err != nil {
		t.Fatal(err)
	}
	idx, err := BuildLineIndex(path)
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if idx == nil {
		t.Fatal("nil idx")
	}
	if len(idx.Paths) != 0 || len(idx.Schemas) != 0 {
		t.Errorf("empty yaml should yield empty maps, got paths=%d schemas=%d", len(idx.Paths), len(idx.Schemas))
	}
}
