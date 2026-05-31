//ff:func feature=stml-parse type=test control=sequence
//ff:what TestParseLayoutFile_ZeroCov — ParseLayoutFile 정상/에러 경로
package stml

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseLayoutFile_ZeroCov(t *testing.T) {
	dir := t.TempDir()
	good := filepath.Join(dir, "app.html")
	if err := os.WriteFile(good, []byte(layoutHTML), 0o644); err != nil {
		t.Fatal(err)
	}
	layout, diags := ParseLayoutFile(good)
	if len(diags) > 0 {
		t.Fatalf("unexpected diags: %v", diags)
	}
	if layout.Name != "app" {
		t.Errorf("Name = %q, want app", layout.Name)
	}

	_, diags = ParseLayoutFile(filepath.Join(dir, "nope.html"))
	if len(diags) == 0 {
		t.Errorf("expected diag for missing layout file")
	}
}
