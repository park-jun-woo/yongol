//ff:func feature=generate type=generator control=iteration dimension=1
//ff:what runDomainFrontendPipelines — 도메인마다 STML 코드젠·컴포넌트 복사·tsc 게이트를 도메인 frontend 디렉토리에서 수행
package generate

import (
	"fmt"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/generate/react"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// runDomainFrontendPipelines runs the post-scaffold frontend steps (STML page
// codegen, user-component copy, tsc gate) once per domain, mirroring the
// single-site pipeline against each domain directory <artifacts>/frontend/<name>.
//
// Each iteration uses fs.DomainView(name) so the existing single-site codepaths
// see the domain's parsed STML/OpenAPI data (Decision A). The component/STML
// SOURCE directory is threaded explicitly as filepath.Join(fs.SpecsDir,
// cfg.Frontend) — DomainView keeps SpecsDir shared, so any path synthesized from
// SpecsDir+convention would bypass the view and not be domain-aware (Decision N).
func runDomainFrontendPipelines(fs *yongol.Fullstack, artifactsDir string) error {
	for _, name := range fs.DomainNames() {
		cfg := fs.Manifest.Domains[name]
		view := fs.DomainView(name)
		domainDir := filepath.Join(artifactsDir, "frontend", name)
		srcFrontendDir := filepath.Join(fs.SpecsDir, cfg.Frontend)

		if err := runSTMLCodegen(view, srcFrontendDir, domainDir); err != nil {
			return fmt.Errorf("stml %s: %w", name, err)
		}
		if err := copyFrontendComponents(srcFrontendDir, filepath.Join(domainDir, "src")); err != nil {
			return fmt.Errorf("components %s: %w", name, err)
		}
		if err := react.RunTscCheck(domainDir); err != nil {
			return fmt.Errorf("tsc gate %s: %w", name, err)
		}
	}
	return nil
}
