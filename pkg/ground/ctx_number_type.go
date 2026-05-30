//ff:func feature=rule type=util control=selection topic=openapi
//ff:what ctxNumberType — number(float/double) 의 맥락별 Go 타입 (response float64·float32 / param 미생성)

package ground

// ctxNumberType returns the Go type oapi-codegen generates for an OpenAPI
// `number`, per context. In a response body, format=float → float32 and the
// default → float64. In the parameter context the prior resolver produced no
// Go type for number (it was never registered), so "" is returned to preserve
// that behaviour.
func ctxNumberType(format string, ctx OAPIContext) string {
	switch {
	case ctx == CtxParam:
		return ""
	case format == "float":
		return "float32"
	default:
		return "float64"
	}
}
