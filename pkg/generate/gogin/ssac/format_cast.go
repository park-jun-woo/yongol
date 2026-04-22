//ff:func feature=gen-gogin type=util control=sequence
//ff:what formatPrimitiveCast — OpenAPI format이 oapi-codegen 래퍼 타입을 만들 때 원시 캐스트 반환
package ssac

// openAPIFormatToPrimitive maps OpenAPI `format` values that oapi-codegen
// emits as named wrapper types (e.g. openapi_types.Email) back to the Go
// primitive accepted by downstream primitives like sqlc-generated functions.
//
// Only formats that (a) produce a wrapper type AND (b) have a trivial cast
// to a primitive are listed. date/date-time are omitted — oapi-codegen
// emits time.Time which needs formatting, not a simple cast.
var openAPIFormatToPrimitive = map[string]string{
	"email": "string",
	"uuid":  "string",
}

// formatPrimitiveCast returns the Go primitive type (e.g. "string") when the
// given OpenAPI format requires a wrapper-to-primitive cast. Empty string
// means no cast is needed.
func formatPrimitiveCast(format string) string {
	return openAPIFormatToPrimitive[format]
}
