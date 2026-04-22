//ff:func feature=external type=util control=iteration dimension=1
//ff:what OpenAPI 스키마 맵의 키를 정렬하여 반환한다
package external

import (
	"sort"

	"github.com/getkin/kin-openapi/openapi3"
)

func sortedKeys(m openapi3.Schemas) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
