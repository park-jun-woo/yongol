//ff:func feature=gen-react type=util control=iteration dimension=1
//ff:what OpenAPI 문서에서 엔드포인트 목록을 수집하여 정렬 반환한다

package react

import (
	"sort"

	"github.com/getkin/kin-openapi/openapi3"
)

// collectEndpoints extracts sorted endpoint list from OpenAPI document.
func collectEndpoints(doc *openapi3.T) []endpoint {
	var endpoints []endpoint

	for path, pi := range doc.Paths.Map() {
		endpoints = appendOperations(endpoints, path, pi)
	}
	sort.Slice(endpoints, func(i, j int) bool {
		return endpoints[i].opID < endpoints[j].opID
	})

	return endpoints
}
