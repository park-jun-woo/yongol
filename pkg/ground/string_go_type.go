//ff:func feature=rule type=util control=selection topic=openapi
//ff:what stringGoType — OpenAPI string format → 맥락별(응답/파라미터) oapi-codegen Go 타입

package ground

// stringGoType maps an OpenAPI string `format` to the Go type oapi-codegen
// generates, per rendering context. Ground truth (oapi-codegen实측):
//
//	format     | ResponseBody          | Param
//	-----------+-----------------------+----------------------
//	uuid       | openapi_types.UUID    | openapi_types.UUID
//	email      | openapi_types.Email   | openapi_types.Email
//	date-time  | time.Time             | string
//	(none)/etc | string                | string
//
// The date-time divergence is the only context-sensitive row: a response-body
// date-time field is rendered as time.Time, whereas a path/query date-time
// parameter is rendered as a plain string.
func stringGoType(format string, ctx OAPIContext) string {
	switch format {
	case "uuid":
		return "openapi_types.UUID"
	case "email":
		return "openapi_types.Email"
	case "date-time":
		if ctx == CtxResponseBody {
			return "time.Time"
		}
		return "string"
	default:
		return "string"
	}
}
