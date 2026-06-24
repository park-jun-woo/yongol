//ff:func feature=rule type=loader control=iteration dimension=1
//ff:what populateOpenAPISingle — 단일 OpenAPI doc 에서 operationId/path/security 를 Ground 에 MERGE (도메인 루프 안전, ASSIGN 아님)
package ground

import (
	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/rule"
)

// populateOpenAPISingle extracts operationId, path, per-path methods, and
// security schemes from one OpenAPI document. The operationId/path/security
// sets are MERGED into any existing set on the Ground (union, not assign) so a
// per-domain loop accumulates across documents instead of the last domain
// overwriting the rest. Per-path "OpenAPI.method.<path>" stays an assign — path
// is the key and XDO-90 forbids cross-domain path duplicates.
func populateOpenAPISingle(g *rule.Ground, doc *openapi3.T) {
	if doc == nil {
		return
	}
	opIDs := mergedSet(g, "OpenAPI.operationId")
	paths := mergedSet(g, "OpenAPI.path")
	security := mergedSet(g, "OpenAPI.security")

	for path, item := range doc.Paths.Map() {
		paths[path] = true
		methods := make(rule.StringSet)
		populatePathOps(g, opIDs, methods, item.Operations())
		g.Lookup["OpenAPI.method."+path] = methods
	}
	g.Lookup["OpenAPI.operationId"] = opIDs
	g.Lookup["OpenAPI.path"] = paths

	if doc.Components != nil {
		for name := range doc.Components.SecuritySchemes {
			security[name] = true
		}
	}
	g.Lookup["OpenAPI.security"] = security
}
