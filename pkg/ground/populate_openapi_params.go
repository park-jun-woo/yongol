//ff:func feature=rule type=loader control=sequence
//ff:what populateOpenAPIParams — 단일/도메인 모드 분기: operationId별 param/request 등록 (opID-keyed, 루프 안전)
package ground

import (
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// populateOpenAPIParams registers per-operation path/query params and request
// fields. Keys are operationId-scoped and operationIds are globally unique
// under XDO-90, so the per-domain loop never collides. Single-site mode loads
// the singular fs.OpenAPIDoc.
func populateOpenAPIParams(g *rule.Ground, fs *yongol.Fullstack) {
	if fs.IsDomained() {
		for _, doc := range fs.DomainOpenAPIDocs {
			populateOpenAPIParamsSingle(g, doc)
		}
		return
	}
	populateOpenAPIParamsSingle(g, fs.OpenAPIDoc)
}
