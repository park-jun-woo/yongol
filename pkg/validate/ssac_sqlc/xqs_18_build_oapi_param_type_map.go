//ff:func feature=validate type=util control=iteration dimension=1 topic=ssac-sqlc
//ff:what buildXqs18OAPIParamTypeMap — OpenAPI Operation 의 parameter 이름 → 타입 맵

package ssac_sqlc

import (
	"github.com/getkin/kin-openapi/openapi3"
)

// buildXqs18OAPIParamTypeMap returns param name → OpenAPI type string for an operation.
func buildXqs18OAPIParamTypeMap(op *openapi3.Operation) map[string]string {
	result := make(map[string]string)
	for _, pRef := range op.Parameters {
		if pRef == nil || pRef.Value == nil {
			continue
		}
		p := pRef.Value
		if p.Schema == nil || p.Schema.Value == nil {
			continue
		}
		types := p.Schema.Value.Type
		if types == nil || len(*types) == 0 {
			continue
		}
		result[p.Name] = (*types)[0]
	}
	return result
}
