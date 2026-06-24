//ff:func feature=gen-react type=generator control=iteration dimension=1
//ff:what generateDomainFrontends — 도메인마다 frontend/<name> 아래 독립 React 앱 스캐폴드 방출 (per-domain view + spec 경로)

package react

import (
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// generateDomainFrontends emits one independent React app per domain under
// <artifactsDir>/frontend/<domain>/. Each iteration reuses the single-site
// codepath with fs.DomainView(name) — the singular OpenAPIDoc/STMLPages/...
// fields are swapped to the domain's parsed data (Decision A), so no per-reader
// rerouting is needed. The per-domain spec PATH is threaded explicitly from
// cfg.OpenAPI (Decision N): DomainView keeps SpecsDir shared and does NOT
// redirect SpecsDir-based path synthesis, so filepath.Join(fs.SpecsDir,
// cfg.OpenAPI) (matching parse_domains_if_present.go:36) is the only
// domain-aware spec path.
func generateDomainFrontends(fs *yongol.Fullstack, artifactsDir string) error {
	for _, name := range fs.DomainNames() {
		cfg := fs.Manifest.Domains[name]
		view := fs.DomainView(name)
		frontendDir := filepath.Join(artifactsDir, "frontend", name)
		domainSpecPath := filepath.Join(fs.SpecsDir, cfg.OpenAPI)
		if err := generateFrontendSetup(view, frontendDir, domainSpecPath); err != nil {
			return err
		}
	}
	return nil
}
