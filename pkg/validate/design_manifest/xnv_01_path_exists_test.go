//ff:func feature=validate type=test control=sequence topic=design-manifest
//ff:what TestXnv01_Positive_PathExists — 경로 존재 시 진단 0건

package design_manifest

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXnv01_Positive_PathExists(t *testing.T) {
	tmp := t.TempDir()
	designDir := filepath.Join(tmp, "frontend")
	if err := os.MkdirAll(designDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(designDir, "DESIGN.md"), []byte("# design"), 0o644); err != nil {
		t.Fatal(err)
	}

	fs := &yongol.Fullstack{
		SpecsDir: tmp,
		Manifest: &manifest.ProjectConfig{
			Frontend: manifest.Frontend{Design: "frontend/DESIGN.md"},
		},
	}
	diags := xnv01PathExists(fs)
	if len(diags) != 0 {
		t.Fatalf("want 0 diags, got %+v", diags)
	}
}
