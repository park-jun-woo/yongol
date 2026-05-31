//ff:func feature=rule type=test control=sequence
//ff:what resolveOAPIGoType test — type×format×shape×context 매트릭스 (oapi-codegen ground truth 대조)
package ground

import (
	"github.com/getkin/kin-openapi/openapi3"
)

func objSchema() *openapi3.SchemaRef {
	return &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"object"}}}
}
