//ff:func feature=validate type=util control=sequence topic=stml-openapi
//ff:what schemaType — SchemaRef 에서 primary type 문자열 반환

package stml_openapi

import "github.com/getkin/kin-openapi/openapi3"

// schemaType returns the primary type string from a schema ref, or "" if unavailable.
func schemaType(ref *openapi3.SchemaRef) string {
	if ref == nil || ref.Value == nil || ref.Value.Type == nil {
		return ""
	}
	types := ref.Value.Type.Slice()
	if len(types) == 0 {
		return ""
	}
	return types[0]
}
