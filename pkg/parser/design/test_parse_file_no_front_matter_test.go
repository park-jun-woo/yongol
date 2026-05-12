//ff:func feature=design-parse type=test control=sequence
//ff:what TestParseFile_NoFrontMatter — front matter 없는 DESIGN.md 에러 검증

package design

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseFile_NoFrontMatter(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "DESIGN.md")
	if err := os.WriteFile(path, []byte("# No front matter\n"), 0644); err != nil {
		t.Fatal(err)
	}
	_, diags := ParseFile(path)
	if len(diags) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d", len(diags))
	}
	if diags[0].Level != "ERROR" {
		t.Errorf("expected ERROR level, got %q", diags[0].Level)
	}
}
