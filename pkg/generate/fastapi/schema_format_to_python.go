//ff:func feature=gen-fastapi type=util control=selection
//ff:what schemaFormatToPython — OpenAPI format 문자열 → Python 타입 매핑

package fastapi

// schemaFormatToPython maps an OpenAPI format string to a Python type.
func schemaFormatToPython(format string) string {
	switch format {
	case "email", "uuid", "uri", "url", "":
		return "str"
	case "date-time", "date":
		return "str"
	case "int32", "int64":
		return "int"
	case "float", "double":
		return "float"
	case "boolean":
		return "bool"
	default:
		return "str"
	}
}
