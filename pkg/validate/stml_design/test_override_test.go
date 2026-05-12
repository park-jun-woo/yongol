//ff:func feature=validate type=test control=sequence topic=stml-design
//ff:what TestOverride_SuppressesXVM01

package stml_design

import (
	"os"
	"path/filepath"
	"testing"
	"github.com/park-jun-woo/yongol/pkg/parser/design"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestOverride_SuppressesXVM01(t *testing.T) {
	tmpDir := t.TempDir()
	frontendDir := filepath.Join(tmpDir, "frontend")
	os.MkdirAll(frontendDir, 0o755)
	os.WriteFile(filepath.Join(frontendDir, "page.html"), []byte(
		`<main>
<!-- @override class="bg-custom" -->
<div>Hello</div>
<div>World</div>
</main>`), 0o644)

	fs := &yongol.Fullstack{
		SpecsDir: tmpDir,
		DesignSpec: &design.DesignSpec{
			File:   "DESIGN.md",
			Colors: map[string]string{"primary": "#6366F1"},
		},
		STMLPages: []stml.PageSpec{{
			Name:     "page",
			FileName: "page.html",
			Children: []stml.ChildNode{
				{Kind: "static", Static: &stml.StaticElement{ClassName: "bg-custom"}},
				{Kind: "static", Static: &stml.StaticElement{ClassName: "bg-another"}},
			},
		}},
	}

	tokens := extractAllTokens(fs)
	ovr := collectOverrides(fs)
	diags := xvm01Color(fs, tokens, ovr)

	// "bg-custom" should be suppressed, "bg-another" should fire
	if len(diags) != 1 {
		t.Fatalf("expected 1 (override suppresses bg-custom), got %d: %+v", len(diags), diags)
	}
	if diags[0].Message == "" {
		t.Fatal("expected non-empty message")
	}
}
