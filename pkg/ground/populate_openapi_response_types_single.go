//ff:func feature=rule type=loader control=iteration dimension=3
//ff:what populateOpenAPIResponseTypesSingle — 단일 OpenAPI doc 의 2xx response schema 각 field → 타입 이름 등록
package ground

import (
	"strings"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/rule"
)

// populateOpenAPIResponseTypesSingle takes, per operation, the first 2xx
// response's application/json schema and maps each top-level property via
// resolveSchemaType ($ref normalized, primitives normalized to Go types):
//
//	Types["OpenAPI.response.<opID>.<fieldName>"] = <Type>
func populateOpenAPIResponseTypesSingle(g *rule.Ground, doc *openapi3.T) {
	if doc == nil || doc.Paths == nil {
		return
	}
	for _, item := range doc.Paths.Map() {
		if item == nil {
			continue
		}
		for _, op := range item.Operations() {
			if op == nil || op.OperationID == "" || op.Responses == nil {
				continue
			}
			for code, resp := range op.Responses.Map() {
				if !strings.HasPrefix(code, "2") || resp.Value == nil || resp.Value.Content == nil {
					continue
				}
				ct := resp.Value.Content.Get("application/json")
				if ct == nil || ct.Schema == nil || ct.Schema.Value == nil {
					continue
				}
				registerOpenAPIResponseProps(g, op.OperationID, ct.Schema.Value)
				break // first 2xx only
			}
		}
	}
}
