//ff:func feature=rule type=util control=selection topic=openapi
//ff:what ctxNumberType — number 의 format별 Go 타입 (double→float64 / 그 외(formatless·float)→float32)

package ground

// ctxNumberType returns the Go type oapi-codegen generates for an OpenAPI
// `number`. oapi-codegen v2.6.0 renders a formatless number — and one with
// format=float — as float32, and only format=double as float64. The type is
// context-independent (response body and parameter alike), so no context
// branch is needed.
func ctxNumberType(format string) string {
	switch format {
	case "double":
		return "float64"
	default:
		return "float32"
	}
}
