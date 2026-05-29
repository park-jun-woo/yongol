//ff:func feature=validate type=test control=sequence topic=design-manifest
//ff:what TestXnv02_Positive_NoDiagWhenDeclared — 선언된 파일 시 진단 0건

package design_manifest

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXnv02_Positive_NoDiagWhenDeclared(t *testing.T) {
	tmp := t.TempDir()
	frontendDir := filepath.Join(tmp, "frontend")
	if err := os.MkdirAll(frontendDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(frontendDir, "DESIGN.md"), []byte("# d"), 0o644); err != nil {
		t.Fatal(err)
	}

	fs := &yongol.Fullstack{
		SpecsDir: tmp,
		Manifest: &manifest.ProjectConfig{
			Frontend: manifest.Frontend{Design: "frontend/DESIGN.md"},
		},
	}
	diags := xnv02Undeclared(fs)
	if len(diags) != 0 {
		t.Fatalf("want 0 diags (declared), got %+v", diags)
	}
}
