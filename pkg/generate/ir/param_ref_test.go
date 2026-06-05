//ff:func feature=gen-ir type=test-helper control=sequence
//ff:what paramRef — 테스트 헬퍼

package ir

import (
	"github.com/getkin/kin-openapi/openapi3"
)

func paramRef(name, in string, required bool, typ string) *openapi3.ParameterRef {
	return &openapi3.ParameterRef{Value: &openapi3.Parameter{
		Name: name, In: in, Required: required,
		Schema: openapi3.NewSchemaRef("", &openapi3.Schema{Type: &openapi3.Types{typ}}),
	}}
}
