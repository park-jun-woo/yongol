//ff:func feature=validate type=util control=iteration dimension=1 topic=scenario-check
//ff:what collectOpenAPIRoutes — OpenAPI Doc의 전체 정규화 라우트 목록 생성

package openapi_hurl

import (
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// collectOpenAPIRoutes builds normalized routes from an OpenAPI document.
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
			})
		}
	}
	return routes
}
