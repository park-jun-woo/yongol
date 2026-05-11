//ff:func feature=validate type=test control=sequence topic=design-manifest
//ff:what TestXnv01 — manifest.frontend.design 경로 존재 여부 검증 테스트

package design_manifest

import (
	"os"
	"path/filepath"
	"strings"
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

func TestXnv01_Skip_NilManifest(t *testing.T) {
	fs := &yongol.Fullstack{}
	diags := xnv01PathExists(fs)
	if len(diags) != 0 {
		t.Fatalf("want 0 diags (nil manifest), got %+v", diags)
	}
}
