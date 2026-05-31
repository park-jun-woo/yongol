//ff:func feature=rule type=loader control=iteration dimension=3
//ff:what populateOpenAPIResponseTypes — OpenAPI 2xx response schema 의 각 field → 타입 이름 등록
package ground

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// populateOpenAPIResponseTypes registers per-operation response field types.
// Takes the first 2xx response's application/json schema; each top-level
// property is mapped via resolveSchemaType ($ref normalized, primitives
// normalized to Go types).
//
//	Types["OpenAPI.response.<opID>.<fieldName>"] = <Type>
//
// XOS-67 consumes this to validate @response field value types.
func populateOpenAPIResponseTypes(g *rule.Ground, fs *yongol.Fullstack) {
	if fs.OpenAPIDoc == nil || fs.OpenAPIDoc.Paths == nil {
		return
	}
	for _, item := range fs.OpenAPIDoc.Paths.Map() {
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
