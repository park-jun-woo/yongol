//ff:func feature=validate type=test control=sequence topic=design-manifest
//ff:what TestXnv02 — specs/frontend/ 내 미선언 DESIGN.md 경고 테스트

package design_manifest

import (
	"os"
	"path/filepath"
	"strings"
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

func TestXnv02_Negative_UndeclaredFile(t *testing.T) {
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
			Frontend: manifest.Frontend{Design: ""},
		},
	}
	diags := xnv02Undeclared(fs)
	if len(diags) != 1 {
		t.Fatalf("want 1 diag, got %d: %+v", len(diags), diags)
	}
	if !strings.Contains(diags[0].Message, "[XNV-02]") {
		t.Fatalf("want XNV-02 message, got %q", diags[0].Message)
	}
}

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

func TestXnv02_Skip_NilManifest(t *testing.T) {
	fs := &yongol.Fullstack{}
	diags := xnv02Undeclared(fs)
	if len(diags) != 0 {
		t.Fatalf("want 0 diags (nil manifest), got %+v", diags)
	}
}
