//ff:func feature=external type=util control=iteration dimension=1
//ff:what OpenAPI 문서의 경로 키를 정렬하여 반환한다
package external

import (
	"sort"

	"github.com/getkin/kin-openapi/openapi3"
)

func sortedPathKeys(doc *openapi3.T) []string {
	paths := make([]string, 0, len(doc.Paths.Map()))
	for p := range doc.Paths.Map() {
		paths = append(paths, p)
	}
	sort.Strings(paths)
	return paths
}
