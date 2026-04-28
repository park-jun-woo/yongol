//ff:func feature=validate type=util control=iteration dimension=2 topic=hurl-openapi
//ff:what collectOpenAPIRoutes — OpenAPI Doc 의 전체 정규화 route 목록 생성

package hurl_openapi

import (
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// collectOpenAPIRoutes builds normalized routes from an OpenAPI document.
// Every operation becomes one route; the method is upper-cased so hurl
// lookups need not re-normalise.
func collectOpenAPIRoutes(doc *openapi3.T) []apiRoute {
	var routes []apiRoute
	if doc == nil || doc.Paths == nil {
		return routes
	}
	for path, pi := range doc.Paths.Map() {
		segs := normalizeOpenAPIPath(path)
		for method, op := range pi.Operations() {
			routes = append(routes, apiRoute{
				Path:      path,
				Method:    strings.ToUpper(method),
				Segments:  segs,
				Responses: collectResponseCodes(op),
				Op:        op,
			})
		}
	}
	return routes
}
