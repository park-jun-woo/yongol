//ff:func feature=design-parse type=test control=sequence
//ff:what TestParseFile_InvalidYAML — 잘못된 YAML front matter 에러 검증

package design

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseFile_InvalidYAML(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "DESIGN.md")
	content := "---\n: invalid: yaml: [broken\n---\n"
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
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
