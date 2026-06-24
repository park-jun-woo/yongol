//ff:func feature=rule type=loader control=sequence
//ff:what populateOpenAPIResponseTypes — 단일/도메인 모드 분기: operationId별 2xx response field 타입 등록 (opID-keyed, 루프 안전)
package ground

import (
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// populateOpenAPIResponseTypes registers per-operation response field types.
// Keys are operationId-scoped (globally unique under XDO-90), so the per-domain
// loop is collision-free. Single-site mode loads the singular fs.OpenAPIDoc.
// XOS-67 consumes this to validate @response field value types.
func populateOpenAPIResponseTypes(g *rule.Ground, fs *yongol.Fullstack) {
	if fs.IsDomained() {
		for _, doc := range fs.DomainOpenAPIDocs {
			populateOpenAPIResponseTypesSingle(g, doc)
		}
		return
	}
	populateOpenAPIResponseTypesSingle(g, fs.OpenAPIDoc)
}
