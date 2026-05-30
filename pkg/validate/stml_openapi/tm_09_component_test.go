//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestTM09Component — 컴포넌트 .tsx 존재/미존재 분기 검증

package stml_openapi

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestTM09Component(t *testing.T) {
	tmp := t.TempDir()
	fs := &yongol.Fullstack{SpecsDir: tmp}

	// Missing component → TM-09 diagnostic.
	diags := tm09Component("Missing", "p.html", fs)
	if len(diags) != 1 {
		t.Fatalf("missing: expected 1 diagnostic, got %d: %+v", len(diags), diags)
	}

	// Existing component → nil.
	compDir := filepath.Join(tmp, "frontend", "components")
	if err := os.MkdirAll(compDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(compDir, "Modal.tsx"), []byte("export {}"), 0o644); err != nil {
		t.Fatal(err)
	}
	if got := tm09Component("Modal", "p.html", fs); got != nil {
		t.Fatalf("existing: expected nil, got %+v", got)
	}
}
