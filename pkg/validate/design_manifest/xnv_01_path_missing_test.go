//ff:func feature=validate type=test control=sequence topic=design-manifest
//ff:what TestXnv01_Negative_PathMissing — 경로 미존재 시 XNV-01 진단 1건

package design_manifest

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXnv01_Negative_PathMissing(t *testing.T) {
	tmp := t.TempDir()

	fs := &yongol.Fullstack{
		SpecsDir: tmp,
		Manifest: &manifest.ProjectConfig{
			Frontend: manifest.Frontend{Design: "frontend/DESIGN.md"},
		},
	}
	diags := xnv01PathExists(fs)
	if len(diags) != 1 {
		t.Fatalf("want 1 diag, got %d", len(diags))
	}
	if !strings.Contains(diags[0].Message, "[XNV-01]") {
		t.Fatalf("want XNV-01 message, got %q", diags[0].Message)
	}
}
