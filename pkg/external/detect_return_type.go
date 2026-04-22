//ff:func feature=external type=generator control=sequence
//ff:what 오퍼레이션의 200 응답에서 반환 타입을 감지한다
package external

import (
	"github.com/getkin/kin-openapi/openapi3"
)

func detectReturnType(op *openapi3.Operation) string {
	resp, ok := op.Responses.Map()["200"]
	if !ok || resp.Value == nil {
		return ""
	}
	ct := resp.Value.Content.Get("application/json")
	if ct == nil || ct.Schema == nil || ct.Schema.Value == nil {
		return ""
	}
	return toPascalCase(op.OperationID) + "Response"
}
