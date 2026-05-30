//ff:func feature=rule type=util control=selection topic=openapi
//ff:what ctxIntDefault — format 없는 integer 의 맥락별 기본 Go 타입 (param int32 / response int)

package ground

// ctxIntDefault returns the Go type oapi-codegen generates for an OpenAPI
// integer with no (or a non-int64) format. oapi-codegen renders a formatless
// integer parameter as int32 but a formatless response-body integer field as
// int, so the default differs by context.
func ctxIntDefault(ctx OAPIContext) string {
	switch ctx {
	case CtxParam:
		return "int32"
	default:
		return "int"
	}
}
