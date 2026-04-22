//ff:func feature=external type=generator control=iteration dimension=1
//ff:what OpenAPI 문서에서 메서드 정보를 추출한다
package external

import (
	"github.com/getkin/kin-openapi/openapi3"
)

func extractMethods(doc *openapi3.T) []methodInfo {
	var methods []methodInfo
	if doc.Paths == nil {
		return methods
	}

	paths := sortedPathKeys(doc)
	for _, path := range paths {
		pi := doc.Paths.Map()[path]
		methods = append(methods, extractPathMethods(pi, path)...)
	}

	return methods
}
