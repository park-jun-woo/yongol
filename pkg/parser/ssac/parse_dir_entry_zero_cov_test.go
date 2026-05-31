//ff:func feature=ssac-parse type=test control=sequence
//ff:what TestParseBatch_ZeroCov — ssac 파서 헬퍼를 이름으로 직접 호출해 커버 귀속
package ssac

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseDirEntry_ZeroCov(t *testing.T) {
	dir := t.TempDir()
	feature := filepath.Join(dir, "course")
	if err := os.MkdirAll(feature, 0o755); err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(feature, "get_course.ssac")
	if err := os.WriteFile(path, []byte(sampleSSaC), 0o644); err != nil {
		t.Fatal(err)
	}
	sfs, diags := parseDirEntry(dir, path, "get_course.ssac")
	if len(diags) != 0 {
		t.Fatalf("unexpected diags: %v", diags)
	}
	if len(sfs) == 0 || sfs[0].Feature != "course" {
		t.Fatalf("expected feature course, got %+v", sfs)
	}
	// file directly in dir → error diag
	flat := filepath.Join(dir, "flat.ssac")
	if err := os.WriteFile(flat, []byte(sampleSSaC), 0o644); err != nil {
		t.Fatal(err)
	}
	_, diags2 := parseDirEntry(dir, flat, "flat.ssac")
	if len(diags2) == 0 {
		t.Error("expected error diag for flat file")
	}
}
