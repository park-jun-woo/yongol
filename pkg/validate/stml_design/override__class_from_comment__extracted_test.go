//ff:func feature=validate type=test control=sequence topic=stml-design
//ff:what TestOverride_ClassFromComment_Extracted

package stml_design

import (
	"os"
	"path/filepath"
	"testing"
	"github.com/park-jun-woo/yongol/pkg/parser/design"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestOverride_ClassFromComment_Extracted(t *testing.T) {
	tmpDir := t.TempDir()
	frontendDir := filepath.Join(tmpDir, "frontend")
	os.MkdirAll(frontendDir, 0o755)
	os.WriteFile(filepath.Join(frontendDir, "page.html"), []byte(
		`<main>
<!-- @override class="bg-neon-green p-8" -->
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
		}},
	}

	ovr := collectOverrides(fs)
	m, ok := ovr["page.html"]
	if !ok {
		t.Fatal("expected override set for page.html")
	}
	if !m["bg-neon-green p-8"] {
		t.Fatalf("expected class 'bg-neon-green p-8' in override set, got %v", m)
	}
}
