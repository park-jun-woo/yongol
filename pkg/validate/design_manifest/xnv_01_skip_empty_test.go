//ff:func feature=validate type=test control=sequence topic=design-manifest
//ff:what TestXnv01_Skip_EmptyDesign — design 경로 비어있을 때 skip

package design_manifest

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXnv01_Skip_EmptyDesign(t *testing.T) {
	tmp := t.TempDir()

	fs := &yongol.Fullstack{
		SpecsDir: tmp,
		Manifest: &manifest.ProjectConfig{
			Frontend: manifest.Frontend{Design: ""},
		},
	}
	diags := xnv01PathExists(fs)
	if len(diags) != 0 {
		t.Fatalf("want 0 diags (skip), got %+v", diags)
	}
}
