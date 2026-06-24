//ff:func feature=rule type=loader control=sequence
//ff:what populateOpenAPI — 단일/도메인 모드 분기: 도메인 모드는 모든 도메인 doc 을 MERGE 누적, 단일 사이트는 fs.OpenAPIDoc 하나만 적재
package ground

import (
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// populateOpenAPI registers operationId/path/method/security from OpenAPI. In
// domain mode it iterates every per-domain document, merging into the shared
// operationId/path/security sets (populateOpenAPISingle unions rather than
// assigns) so no domain overwrites another; in single-site mode it loads the
// singular fs.OpenAPIDoc. Branching here keeps Build's call order intact.
func populateOpenAPI(g *rule.Ground, fs *yongol.Fullstack) {
	if fs.IsDomained() {
		for _, doc := range fs.DomainOpenAPIDocs {
			populateOpenAPISingle(g, doc)
		}
		return
	}
	populateOpenAPISingle(g, fs.OpenAPIDoc)
}
