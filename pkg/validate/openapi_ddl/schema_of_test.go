//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=openapi-ddl
//ff:what schemaOf — top-level 프로퍼티 이름만 갖는 inline object 스키마 빌더

package openapi_ddl

import "github.com/getkin/kin-openapi/openapi3"

// schemaOf builds an inline object schema with the given top-level property
// names. Property value schemas are empty (shape comparison only needs keys).
func schemaOf(props ...string) *openapi3.Schema {
	s := &openapi3.Schema{Properties: openapi3.Schemas{}}
	for _, p := range props {
		s.Properties[p] = &openapi3.SchemaRef{Value: &openapi3.Schema{}}
	}
	return s
}
