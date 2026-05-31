//ff:func feature=validate type=test control=sequence topic=stml-design
//ff:what TestOverride_StructureOnly_NoClass

package stml_design

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/design"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestOverride_StructureOnly_NoClass(t *testing.T) {
	tmpDir := t.TempDir()
	frontendDir := filepath.Join(tmpDir, "frontend")
	os.MkdirAll(frontendDir, 0o755)
	os.WriteFile(filepath.Join(frontendDir, "page.html"), []byte(
		`<main>
<!-- @override -->
<div>Hello</div>
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
				{Kind: "static", Static: &stml.StaticElement{Tag: "div"}},
			},
		}},
	}

	ovr := collectOverrides(fs)
	// Structure-only override should not add any class to override set
	if m, ok := ovr["page.html"]; ok && len(m) > 0 {
		t.Fatalf("expected no class overrides for structure-only @override, got %v", m)
	}
}
