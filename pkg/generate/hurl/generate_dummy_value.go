//ff:func feature=gen-hurl type=util control=selection
//ff:what generateDummyValue — OpenAPI type+format에서 dummy 리터럴 생성
package hurl

import (
	"github.com/getkin/kin-openapi/openapi3"
)

// generateDummyValue returns a Go value suitable for JSON marshaling
// based on the OpenAPI type and format.
func generateDummyValue(types *openapi3.Types, format string) any {
	typ := ""
	if types != nil && len(*types) > 0 {
		typ = (*types)[0]
	}
	switch typ {
	case "string":
		return dummyString(format)
	case "integer":
		return 1
	case "number":
		return 1.0
	case "boolean":
		return true
	case "array":
		return []any{}
	case "object":
		return map[string]any{}
	default:
		return "test"
	}
}
