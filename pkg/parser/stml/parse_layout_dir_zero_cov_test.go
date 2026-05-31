//ff:func feature=stml-parse type=test control=sequence
//ff:what TestParseLayoutDir_ZeroCov — ParseLayoutDir 정상/에러 경로
package stml

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseLayoutDir_ZeroCov(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "app.html"), []byte(layoutHTML), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "skip.md"), []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.Mkdir(filepath.Join(dir, "nested"), 0o755); err != nil {
		t.Fatal(err)
	}
	layouts, diags := ParseLayoutDir(dir)
	if len(diags) > 0 {
		t.Fatalf("unexpected diags: %v", diags)
	}
	if len(layouts) != 1 {
		t.Errorf("layouts = %d, want 1", len(layouts))
	}

	_, diags = ParseLayoutDir(filepath.Join(dir, "missing"))
	if len(diags) == 0 {
		t.Errorf("expected diag for missing layout dir")
	}
}
