//ff:func feature=orchestrator type=loader control=iteration dimension=1
//ff:what 멀티 도메인 프로젝트의 도메인별 OpenAPI·STML·sitemap·layout 을 단일 패스로 적재하고 같은 패스에서 presence 파생
package yongol

import (
	"path/filepath"

	"github.com/getkin/kin-openapi/openapi3"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// parseDomainsIfPresent loads per-domain SSOTs for a multi-domain project. It is
// a no-op for single-site projects (fs.IsDomained() == false), leaving the
// singular OpenAPIDoc/STMLPages/... fields as the only data source.
//
// Each domain is processed exactly once: the per-SSOT loadDomain* helpers load
// the OpenAPI contract, STML pages, sitemap and layouts, and the OpenAPI/STML
// presence is derived in the SAME iteration from those loads (Decision G) so the
// presence map never drifts from the data actually loaded. Shared SSOTs
// (DDL/sqlc/SSaC/Rego/states/Hurl/FuncSpec/features) are loaded once by the caller
// and are intentionally NOT repeated per domain.
func parseDomainsIfPresent(fs *Fullstack, root string) {
	if !fs.IsDomained() {
		return
	}
	fs.DomainOpenAPIDocs = make(map[string]*openapi3.T)
	fs.DomainOpenAPILines = make(map[string]*oapiparser.LineIndex)
	fs.DomainSTMLPages = make(map[string][]stml.PageSpec)
	fs.DomainSitemaps = make(map[string]*stml.SitemapSpec)
	fs.DomainLayouts = make(map[string][]stml.LayoutSpec)
	fs.DomainPresences = make(map[string]map[SSOTKind]SSOTPresence)

	for name, cfg := range fs.Manifest.Domains {
		oapiPath := filepath.Join(root, cfg.OpenAPI)
		frontDir := filepath.Join(root, cfg.Frontend)
		presence := map[SSOTKind]SSOTPresence{}
		presence[KindOpenAPI] = loadDomainOpenAPI(fs, name, oapiPath)
		presence[KindSTML] = loadDomainSTML(fs, name, frontDir)
		loadDomainSitemap(fs, name, frontDir)
		loadDomainLayouts(fs, name, frontDir)
		fs.DomainPresences[name] = presence
	}
}
