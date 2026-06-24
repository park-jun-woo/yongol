//ff:func feature=generate type=test control=sequence
//ff:what TestRunDomainFrontendPipelines — 도메인 루프 경로 파생(STML 무·소스 무·tsc skip) 무에러 검증
package generate

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestRunDomainFrontendPipelines exercises the per-domain loop path derivation:
// for each domain it must (1) run STML codegen against the domain view (no STML
// pages → no-op), (2) copy components from filepath.Join(SpecsDir, cfg.Frontend)
// (missing dir → no-op), and (3) run the tsc gate on <artifacts>/frontend/<name>
// (no node_modules → graceful skip). No domain has pages/sources, so the loop
// must complete without error while visiting every domain.
func TestRunDomainFrontendPipelines(t *testing.T) {
	specs := t.TempDir()
	arts := t.TempDir()
	fs := &yongol.Fullstack{
		SpecsDir: specs,
		Manifest: &manifest.ProjectConfig{
			Domains: map[string]manifest.DomainConfig{
				"public": {Frontend: "frontend"},
				"admin":  {Frontend: "admin/frontend"},
			},
		},
	}
	if !fs.IsDomained() {
		t.Fatal("fixture must be domained")
	}
	if err := runDomainFrontendPipelines(fs, arts); err != nil {
		t.Fatalf("runDomainFrontendPipelines: %v", err)
	}
}
