//ff:func feature=external type=generator control=iteration dimension=1
//ff:what 오퍼레이션의 요청 바디에서 파라미터를 추출한다
package external

import (
	"github.com/getkin/kin-openapi/openapi3"
)

func extractBodyParams(op *openapi3.Operation) []paramInfo {
	if op.RequestBody == nil || op.RequestBody.Value == nil {
		return nil
	}
	ct := op.RequestBody.Value.Content.Get("application/json")
	if ct == nil || ct.Schema == nil || ct.Schema.Value == nil {
		return nil
	}
	var params []paramInfo
	propNames := sortedKeys(ct.Schema.Value.Properties)
	for _, name := range propNames {
		ref := ct.Schema.Value.Properties[name]
		params = append(params, paramInfo{
			Name:   toCamelCase(name),
			GoType: schemaToGoType(ref),
			In:     "body",
		})
	}
	return params
}
