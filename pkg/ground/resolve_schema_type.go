//ff:func feature=rule type=util control=sequence
//ff:what resolveSchemaType — OpenAPI schema ref 를 타입 이름 문자열로 정규화
package ground

import (
	"github.com/getkin/kin-openapi/openapi3"
)

// resolveSchemaType converts an OpenAPI schema reference into a canonical
// type name for type-compatibility comparisons. Primitive-type dispatch lives
// in resolvePrimitiveType.
func resolveSchemaType(ref *openapi3.SchemaRef) string {
	if ref == nil {
		return ""
	}
	if ref.Ref != "" {
		return refName(ref.Ref)
	}
	return resolvePrimitiveType(ref.Value)
}
