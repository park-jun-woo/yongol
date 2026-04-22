//ff:func feature=external type=util control=selection
//ff:what OpenAPI 스키마를 Go 타입 문자열로 변환한다
package external

import (
	"github.com/getkin/kin-openapi/openapi3"
)

func schemaToGoType(ref *openapi3.SchemaRef) string {
	if ref == nil || ref.Value == nil {
		return "any"
	}
	s := ref.Value
	types := s.Type.Slice()
	if len(types) == 0 {
		// No type specified (e.g. anyOf, oneOf, allOf, or untyped)
		return "any"
	}
	switch types[0] {
	case "integer":
		switch s.Format {
		case "int32":
			return "int32"
		case "int64":
			return "int64"
		default:
			return "int"
		}
	case "number":
		switch s.Format {
		case "float":
			return "float32"
		default:
			return "float64"
		}
	case "string":
		if s.Format == "date-time" {
			return "time.Time"
		}
		return "string"
	case "boolean":
		return "bool"
	case "array":
		if s.Items != nil {
			return "[]" + schemaToGoType(s.Items)
		}
		return "[]any"
	case "object":
		return "map[string]any"
	default:
		return "any"
	}
}
