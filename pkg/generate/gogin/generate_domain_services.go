//ff:func feature=gen-gogin type=command control=iteration dimension=1
//ff:what generateDomainServices — 도메인마다 fs.DomainView 로 internal/service 메서드 생성 (alias+prefix)

package gogin

import (
	"fmt"

	ssacgen "github.com/park-jun-woo/yongol/pkg/generate/gogin/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// generateDomainServices runs the shared internal/service codegen once per
// manifest domain over fs.DomainView(name). Each pass aliases the domain's
// oapi-codegen package to `api` (sanitizeDomainName(name)) and prefixes
// converters with domainTitle(name), so the single internal/service package
// compiles against per-domain types without collisions (Phase007). The
// operationId membership filter inside ssac.Generate ensures each method is
// emitted exactly once, by its owning domain.
func generateDomainServices(fs *yongol.Fullstack, artifactsDir string) error {
	for _, name := range fs.DomainNames() {
		view := fs.DomainView(name)
		if err := ssacgen.Generate(view, artifactsDir, sanitizeDomainName(name), domainTitle(name)); err != nil {
			return fmt.Errorf("ssac service [%s]: %w", name, err)
		}
	}
	return nil
}
