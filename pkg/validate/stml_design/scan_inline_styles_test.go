//ff:func feature=validate type=test control=sequence topic=stml-design
//ff:what TestScanInlineStyles — scanInlineStyles inline style 하드코딩 색상 검사 분기 검증
package stml_design

import (
	"os"
	"path/filepath"
	"testing"
)

func TestScanInlineStyles(t *testing.T) {
	hexToToken := map[string]string{"#6366f1": "primary"}

	t.Run("missing file returns nil", func(t *testing.T) {
		d := scanInlineStyles(filepath.Join(t.TempDir(), "nope.html"), "p.html", hexToToken, overrideSet{})
		if d != nil {
			t.Errorf("expected nil, got %+v", d)
		}
	})

	t.Run("hardcoded token color fires", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "page.html")
		os.WriteFile(path, []byte(
			`<main><div style="color: #6366F1">x</div></main>`), 0o644)

		diags := scanInlineStyles(path, "page.html", hexToToken, overrideSet{})
		if len(diags) != 1 {
			t.Fatalf("expected 1 diagnostic, got %d (%+v)", len(diags), diags)
		}
	})
}
