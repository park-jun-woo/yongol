//ff:func feature=rule type=test control=sequence
//ff:what registerParamGoType test — 파라미터 schema → OpenAPI.paramType.<op>.<name> 등록 (맥락·array·skip)
package ground

import (
	"github.com/getkin/kin-openapi/openapi3"
)

func strParam(name, format string) *openapi3.Parameter {
	return &openapi3.Parameter{Name: name, In: "query", Schema: strSchema(format)}
}
