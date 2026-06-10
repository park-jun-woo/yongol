//ff:func feature=validate type=util control=iteration dimension=1 topic=config-check
//ff:what propCandidates — XON-60 advice 용 정렬된 후보 응답 필드명 목록 렌더링

package openapi_manifest

import (
	"sort"
	"strings"
)

// propCandidates renders a sorted candidate field list for XON-60 advice,
// or a hint that no 2xx response declares object properties at all.
func propCandidates(props map[string]bool) string {
	if len(props) == 0 {
		return "No OpenAPI 2xx response declares object properties"
	}
	names := make([]string, 0, len(props))
	for name := range props {
		names = append(names, name)
	}
	sort.Strings(names)
	return "Available 2xx response fields: " + strings.Join(names, ", ")
}
