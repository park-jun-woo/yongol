//ff:func feature=gen-fastapi type=util control=selection
//ff:what openAPITypeToPython — OpenAPI 타입 문자열 → Python 타입 어노테이션 매핑

package ssac

// openAPITypeToPython maps an OpenAPI type string to a Python type annotation.
func openAPITypeToPython(t string) string {
	switch t {
	case "integer":
		return "int"
	case "number":
		return "float"
	case "boolean":
		return "bool"
	default:
		return "str"
	}
}
