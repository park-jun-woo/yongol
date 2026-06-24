//ff:func feature=generate type=test control=sequence
//ff:what TestRunDomainFrontendPipelines_CopiesPerDomainSource — 도메인 소스 dir=filepath.Join(SpecsDir,cfg.Frontend) 검증 (Decision N)
package generate

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestRunDomainFrontendPipelines_CopiesPerDomainSource confirms the SOURCE dir
// is filepath.Join(fs.SpecsDir, cfg.Frontend) (Decision N), not fs.SpecsDir nor
// a view-derived path: a .tsx authored under <specs>/<cfg.Frontend>/ lands under
// <artifacts>/frontend/<name>/src/.
func TestRunDomainFrontendPipelines_CopiesPerDomainSource(t *testing.T) {
	specs := t.TempDir()
	arts := t.TempDir()
	// admin source lives under specs/admin/frontend/components/Foo.tsx.
	comp := filepath.Join(specs, "admin", "frontend", "components")
	if err := os.MkdirAll(comp, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(comp, "Foo.tsx"), []byte("export const Foo = 1"), 0o644); err != nil {
		t.Fatal(err)
	}
	fs := &yongol.Fullstack{
		SpecsDir: specs,
		Manifest: &manifest.ProjectConfig{
			Domains: map[string]manifest.DomainConfig{
				"admin": {Frontend: "admin/frontend"},
			},
		},
	}
	if err := runDomainFrontendPipelines(fs, arts); err != nil {
		t.Fatalf("runDomainFrontendPipelines: %v", err)
	}
	want := filepath.Join(arts, "frontend", "admin", "src", "components", "Foo.tsx")
	if _, err := os.Stat(want); err != nil {
		t.Errorf("expected per-domain source copied to %s: %v", want, err)
	}
}
