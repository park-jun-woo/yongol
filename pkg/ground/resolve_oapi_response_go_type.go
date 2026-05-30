//ff:func feature=rule type=util control=selection
//ff:what resolveOAPIResponseGoType — OpenAPI 응답 본문 type+format → oapi-codegen Go 타입 매핑

package ground

// resolveOAPIResponseGoType maps an OpenAPI response-body field's type and
// format to the Go type that oapi-codegen generates for the response struct
// field (`oapi-codegen -generate types,strict-server,gin`).
//
// Mapping:
//
//	string + uuid      → openapi_types.UUID
//	string + email     → openapi_types.Email
//	string + date-time → time.Time
//	string (no fmt)    → string
//
// This intentionally differs from resolveOAPIParamGoType: parameters have no
// date-time case (oapi-codegen renders a path/query date-time param as a plain
// string), whereas a response-body { type: string, format: date-time } field
// is generated as time.Time. Keeping the two resolvers separate prevents the
// param mapping from silently leaking the wrong type into response validation.
//
// Returns "" for non-string base types or unrecognised combinations; callers
// fall back to resolveSchemaType for $ref/integer/array/object handling.
func resolveOAPIResponseGoType(baseType, format string) string {
	if baseType != "string" {
		return ""
	}
	switch format {
	case "uuid":
		return "openapi_types.UUID"
	case "email":
		return "openapi_types.Email"
	case "date-time":
		return "time.Time"
	default:
		return "string"
	}
}
