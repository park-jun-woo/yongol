//ff:func feature=rule type=util control=sequence topic=openapi
//ff:what ctxIntDefault — format 없는 integer 의 기본 Go 타입 (응답·파라미터 공통 int)

package ground

// ctxIntDefault returns the Go type oapi-codegen generates for an OpenAPI
// integer with no (or a non-int32/int64) format. oapi-codegen v2.6.0 renders a
// formatless integer as int in both response-body and parameter contexts, so
// the default is context-independent.
func ctxIntDefault() string {
	return "int"
}
