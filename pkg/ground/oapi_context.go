//ff:type feature=rule type=model topic=openapi
//ff:what OAPIContext — OpenAPI→Go 타입 렌더링 맥락(응답 본문 / 파라미터) enum

package ground

// OAPIContext selects the rendering context for an OpenAPI→Go type mapping.
// The same OpenAPI schema can map to a different Go type depending on where
// oapi-codegen renders it: a { type: string, format: date-time } field becomes
// time.Time in a response-body struct but a plain string as a path/query
// parameter. Threading the context lets a single recursive resolver honour
// those differences instead of fragmenting the mapping per shape and site.
type OAPIContext int

const (
	// CtxResponseBody is an oapi-codegen response struct field
	// (`-generate types,strict-server,gin`).
	CtxResponseBody OAPIContext = iota
	// CtxParam is a path/query parameter.
	CtxParam
	// (CtxRequestBody can be added later — request bodies currently register
	// only field names, so no Go-type context is needed for them.)
)
