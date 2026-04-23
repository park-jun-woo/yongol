//ff:func feature=openapi-parse type=test control=sequence
//ff:what BuildLineIndex — 존재하지 않는 파일은 err 반환하면서도 non-nil 인덱스 보장

package openapi

import "testing"

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
