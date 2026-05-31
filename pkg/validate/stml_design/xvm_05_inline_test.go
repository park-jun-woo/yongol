//ff:func feature=validate type=test control=sequence topic=stml-design
//ff:what TestXVM05_Golden

package stml_design

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/design"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXVM05_Golden(t *testing.T) {
	// No inline style — no diagnostics
	tmpDir := t.TempDir()
	frontendDir := filepath.Join(tmpDir, "frontend")
	os.MkdirAll(frontendDir, 0o755)
	os.WriteFile(filepath.Join(frontendDir, "page.html"), []byte(`<div class="bg-primary">Hello</div>`), 0o644)

	fs := &yongol.Fullstack{
		SpecsDir: tmpDir,
		DesignSpec: &design.DesignSpec{
			File:   "DESIGN.md",
			Colors: map[string]string{"primary": "#6366F1"},
		},
		STMLPages: []stml.PageSpec{{
			Name:     "page",
			FileName: "page.html",
		}},
	}
	diags := xvm05Inline(fs, nil)
	if len(diags) != 0 {
		t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
	}
}
