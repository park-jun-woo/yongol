//ff:func feature=manifest type=test control=sequence
//ff:what BuildLineIndex — 존재하지 않는 파일 / 비어있는 파일에서 에러/빈 인덱스 반환

package openapi

import (
	"os"
	"path/filepath"
	"testing"
)

func TestBuildLineIndex_MissingFile(t *testing.T) {
	idx, err := BuildLineIndex("/nonexistent/oapi.yaml")
	if err == nil {
		t.Errorf("expected error, got nil")
	}
	if idx == nil {
		t.Fatalf("idx must be non-nil even on read error")
	}
	if idx.File != "/nonexistent/oapi.yaml" {
		t.Errorf("File = %q", idx.File)
	}
}

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
