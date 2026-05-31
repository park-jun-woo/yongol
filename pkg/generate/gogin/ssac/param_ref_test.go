//ff:func feature=gen-gogin type=test control=sequence
//ff:what methodGen.addParam 단위 테스트 (path → PathParams, query → QueryParams, 그 외 무시)
package ssac

import (
	"github.com/getkin/kin-openapi/openapi3"
)

func paramRef(name, in string, required bool, typ string) *openapi3.ParameterRef {
	return &openapi3.ParameterRef{Value: &openapi3.Parameter{
		Name:     name,
		In:       in,
		Required: required,
		Schema:   openapi3.NewSchemaRef("", &openapi3.Schema{Type: &openapi3.Types{typ}}),
	}}
}
