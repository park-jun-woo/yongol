//ff:func feature=rule type=test control=sequence
//ff:what resolveOAPIGoType(CtxResponseBody) test — 응답 본문 필드 type+format → Go 타입 (array items format-aware, BUG-102)
package ground

import (
	"github.com/getkin/kin-openapi/openapi3"
)

func strSchema(format string) *openapi3.SchemaRef {
	return &openapi3.SchemaRef{Value: &openapi3.Schema{
		Type:   &openapi3.Types{"string"},
		Format: format,
	}}
}
