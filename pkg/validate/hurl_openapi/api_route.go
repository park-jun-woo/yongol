//ff:type feature=validate type=util topic=hurl-openapi
//ff:what apiRoute — Hurl 매칭용 정규화된 OpenAPI route 타입

package hurl_openapi

import "github.com/getkin/kin-openapi/openapi3"

// apiRoute represents a normalized OpenAPI operation for hurl matching.
// The Op back-reference lets rules walk into the operation's request /
// response schemas without re-indexing paths.
type apiRoute struct {
	Path      string
	Method    string
	Segments  []string
	Responses map[string]bool
	Op        *openapi3.Operation
}
