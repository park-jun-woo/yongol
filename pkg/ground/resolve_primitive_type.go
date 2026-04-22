//ff:func feature=rule type=util control=selection
//ff:what resolvePrimitiveType — OpenAPI 기본 타입(Schema.Type) → Go 타입명 변환

package ground

import (
	"github.com/getkin/kin-openapi/openapi3"
)

// resolvePrimitiveType converts a primitive OpenAPI schema type into a Go
// type name. Returns "" when the type is unrecognised or the items element
// of an array is unresolvable.
func resolvePrimitiveType(s *openapi3.Schema) string {
	if s == nil || len(s.Type.Slice()) == 0 {
		return ""
	}
	t := s.Type.Slice()[0]
	switch t {
	case "integer":
		if s.Format == "int64" {
			return "int64"
		}
		return "int"
	case "number":
		if s.Format == "float" {
			return "float32"
		}
		return "float64"
	case "string":
		return "string"
	case "boolean":
		return "bool"
	case "array":
		inner := resolveSchemaType(s.Items)
		if inner == "" {
			return ""
		}
		return "[]" + inner
	case "object":
		return "object"
	}
	return ""
}
