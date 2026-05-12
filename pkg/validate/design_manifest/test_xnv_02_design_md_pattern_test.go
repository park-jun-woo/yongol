//ff:func feature=validate type=test control=sequence topic=design-manifest
//ff:what TestXnv02_Negative_DesignMdPattern — *.design.md 미선언 시 XNV-02 진단 1건

package design_manifest

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXnv02_Negative_DesignMdPattern(t *testing.T) {
	tmp := t.TempDir()
	frontendDir := filepath.Join(tmp, "frontend")
	if err := os.MkdirAll(frontendDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(frontendDir, "mobile.design.md"), []byte("# mobile"), 0o644); err != nil {
		t.Fatal(err)
	}

	fs := &yongol.Fullstack{
		SpecsDir: tmp,
		Manifest: &manifest.ProjectConfig{
			Frontend: manifest.Frontend{Design: ""},
		},
	}
	diags := xnv02Undeclared(fs)
	if len(diags) != 1 {
		t.Fatalf("want 1 diag for undeclared *.design.md, got %d: %+v", len(diags), diags)
	}
	if !strings.Contains(diags[0].Message, "mobile.design.md") {
		t.Fatalf("want reference to mobile.design.md, got %q", diags[0].Message)
	}
}
