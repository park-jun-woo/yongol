//ff:func feature=validate type=test-helper control=sequence topic=openapi-ddl
//ff:what inlineRef — inline 스키마($ref 없음) SchemaRef 생성

package openapi_ddl

import "github.com/getkin/kin-openapi/openapi3"

// inlineRef returns a SchemaRef holding an inline schema (no $ref).
func inlineRef(props ...string) *openapi3.SchemaRef {
	return &openapi3.SchemaRef{Value: schemaOf(props...)}
}
