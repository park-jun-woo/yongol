//ff:func feature=rule type=util control=selection
//ff:what resolveOAPIParamGoType — OpenAPI param type+format → oapi-codegen Go 타입 매핑

package ground

// resolveOAPIParamGoType maps an OpenAPI parameter's type and format to the
// Go type that oapi-codegen would generate. Returns "" for unrecognised
// combinations.
//
// Mapping:
//
//	string + uuid   → openapi_types.UUID
//	string + email  → openapi_types.Email
//	string (no fmt) → string
//	integer + int32 → int32
//	integer + int64 → int64
//	integer (no fmt)→ int32
//	boolean         → bool
func resolveOAPIParamGoType(baseType, format string) string {
	switch baseType {
	case "string":
		switch format {
		case "uuid":
			return "openapi_types.UUID"
		case "email":
			return "openapi_types.Email"
		default:
			return "string"
		}
	case "integer":
		switch format {
		case "int64":
			return "int64"
		case "int32":
			return "int32"
		default:
			return "int32"
		}
	case "boolean":
		return "bool"
	}
	return ""
}
