//ff:func feature=validate type=test-helper control=sequence topic=stml-openapi
//ff:what typedSchema — 주어진 OpenAPI type 을 선언한 테스트용 SchemaRef 생성

package stml_openapi

import "github.com/getkin/kin-openapi/openapi3"

// typedSchema builds a SchemaRef whose value declares the given OpenAPI type.
func typedSchema(typ string) *openapi3.SchemaRef {
	return &openapi3.SchemaRef{Value: &openapi3.Schema{
		Type: &openapi3.Types{typ},
	}}
}
