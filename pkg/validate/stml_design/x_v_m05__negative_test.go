//ff:func feature=validate type=test control=iteration dimension=1 topic=stml-design
//ff:what TestXVM05_Negative

package stml_design

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/design"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXVM05_Negative(t *testing.T) {
	tmpDir := t.TempDir()
	frontendDir := filepath.Join(tmpDir, "frontend")
	os.MkdirAll(frontendDir, 0o755)
	os.WriteFile(filepath.Join(frontendDir, "page.html"), []byte(`<div style="color: #6366F1; background: #f59e0b">Hello</div>`), 0o644)

	fs := &yongol.Fullstack{
		SpecsDir: tmpDir,
		DesignSpec: &design.DesignSpec{
			File:   "DESIGN.md",
			Colors: map[string]string{"primary": "#6366F1", "accent": "#F59E0B"},
		},
		STMLPages: []stml.PageSpec{{
			Name:     "page",
			FileName: "page.html",
		}},
	}
	diags := xvm05Inline(fs, nil)
	if len(diags) != 2 {
		t.Fatalf("expected 2, got %d: %+v", len(diags), diags)
	}
	for _, d := range diags {
		if !strings.Contains(d.Message, "[XVM-05]") {
			t.Fatalf("expected [XVM-05], got %q", d.Message)
		}
	}
}
