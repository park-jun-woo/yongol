//ff:func feature=stml-parse type=test control=sequence
//ff:what TestParseFile_ZeroCov — ParseFile 정상/에러 경로
package stml

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseFile_ZeroCov(t *testing.T) {
	dir := t.TempDir()
	good := filepath.Join(dir, "page.html")
	if err := os.WriteFile(good, []byte(richPageHTML), 0o644); err != nil {
		t.Fatal(err)
	}
	page, diags := ParseFile(good)
	if len(diags) > 0 {
		t.Fatalf("unexpected diags: %v", diags)
	}
	if page.Route != "/things/:id" {
		t.Errorf("Route = %q", page.Route)
	}

	// Error path: missing file.
	_, diags = ParseFile(filepath.Join(dir, "nope.html"))
	if len(diags) == 0 {
		t.Errorf("expected diag for missing file")
	}
}
