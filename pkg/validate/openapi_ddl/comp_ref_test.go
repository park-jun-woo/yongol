//ff:func feature=validate type=test-helper control=sequence topic=openapi-ddl
//ff:what compRef — component $ref 와 resolved Value 를 동시 보유하는 SchemaRef 생성

package openapi_ddl

import "github.com/getkin/kin-openapi/openapi3"

// compRef returns a SchemaRef that $refs a component while also carrying the
// resolved Value (as kin-openapi does after loading).
func compRef(name string, props ...string) *openapi3.SchemaRef {
	return &openapi3.SchemaRef{Ref: "#/components/schemas/" + name, Value: schemaOf(props...)}
}
