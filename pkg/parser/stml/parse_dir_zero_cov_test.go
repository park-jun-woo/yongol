//ff:func feature=stml-parse type=test control=sequence
//ff:what TestParseDir_ZeroCov — ParseDir 정상/에러 경로
package stml

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseDir_ZeroCov(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "a.html"), []byte(richPageHTML), 0o644); err != nil {
		t.Fatal(err)
	}
	// Non-html and a subdir should be skipped.
	if err := os.WriteFile(filepath.Join(dir, "ignore.txt"), []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.Mkdir(filepath.Join(dir, "sub"), 0o755); err != nil {
		t.Fatal(err)
	}
	pages, diags := ParseDir(dir)
	if len(diags) > 0 {
		t.Fatalf("unexpected diags: %v", diags)
	}
	if len(pages) != 1 {
		t.Errorf("pages = %d, want 1", len(pages))
	}

	// Error path: missing dir.
	_, diags = ParseDir(filepath.Join(dir, "missing"))
	if len(diags) == 0 {
		t.Errorf("expected diag for missing dir")
	}
}
