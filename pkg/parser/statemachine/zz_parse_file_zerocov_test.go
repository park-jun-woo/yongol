//ff:func feature=statemachine type=test control=sequence
//ff:what TestParseFile_ZeroCov — ParseFile 정상 파싱 + 읽기 실패 진단 직접 호출

package statemachine

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseFile_ZeroCov(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "course.md")
	content := "# CourseState\n\n```mermaid\nstateDiagram-v2\n    [*] --> unpublished\n    unpublished --> published: PublishCourse\n```\n"
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
	d, diags := ParseFile(path)
	if len(diags) > 0 {
		t.Fatalf("ParseFile diags: %v", diags)
	}
	if d == nil || d.ID != "course" {
		t.Fatalf("ParseFile = %+v", d)
	}

	// missing file → read error diagnostic.
	_, diags = ParseFile(filepath.Join(dir, "absent.md"))
	if len(diags) == 0 {
		t.Errorf("expected read-error diagnostic for missing file")
	}
}
