//ff:func feature=rule type=loader control=iteration dimension=1
//ff:what populateOpenAPIParamsSingle — 단일 OpenAPI doc 의 각 operation param/request 를 Ground 에 등록
package ground

import (
	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/rule"
)

// populateOpenAPIParamsSingle registers param and request-field sets for every
// operation of one OpenAPI document. opID-keyed and loop-safe.
func populateOpenAPIParamsSingle(g *rule.Ground, doc *openapi3.T) {
	if doc == nil {
		return
	}
	for _, item := range doc.Paths.Map() {
		populatePathOpsParams(g, item.Operations())
	}
}
