//ff:func feature=validate type=test control=sequence topic=design-manifest
//ff:what TestXnv02_Skip_NoFrontendDir — frontend 디렉토리 없을 때 skip

package design_manifest

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXnv02_Skip_NoFrontendDir(t *testing.T) {
	tmp := t.TempDir()
	// No frontend/ directory at all

	fs := &yongol.Fullstack{
		SpecsDir: tmp,
		Manifest: &manifest.ProjectConfig{
			Frontend: manifest.Frontend{Design: ""},
		},
	}
	diags := xnv02Undeclared(fs)
	if len(diags) != 0 {
		t.Fatalf("want 0 diags (no frontend dir), got %+v", diags)
	}
}
