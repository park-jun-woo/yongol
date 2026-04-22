//ff:func feature=external type=generator control=iteration dimension=1
//ff:what 오퍼레이션에서 경로 파라미터를 추출한다
package external

import (
	"github.com/getkin/kin-openapi/openapi3"
)

func extractPathParams(op *openapi3.Operation) []paramInfo {
	var params []paramInfo
	for _, p := range op.Parameters {
		if p.Value != nil && p.Value.In == "path" {
			params = append(params, paramInfo{
				Name:   toCamelCase(p.Value.Name),
				GoType: schemaToGoType(p.Value.Schema),
				In:     "path",
			})
		}
	}
	return params
}
