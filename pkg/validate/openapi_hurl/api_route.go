//ff:type feature=validate type=util topic=scenario-check
//ff:what apiRoute — OpenAPI 경로의 정규화 라우트 타입 (Hurl 매칭용)

package openapi_hurl

// apiRoute represents a normalized OpenAPI route for hurl matching.
type apiRoute struct {
	Path      string
	Method    string
	Segments  []string
	Responses map[string]bool
}
