//ff:func feature=validate type=util control=iteration dimension=1 topic=config-check
//ff:what opCandidates — XON-60 advice 용 정렬된 후보 operationId 목록 렌더링

package openapi_manifest

import (
	"sort"
	"strings"
)

// opCandidates renders a sorted operationId candidate list for XON-60
// advice, or a hint that the document declares no operationIds.
func opCandidates(ids map[string]bool) string {
	if len(ids) == 0 {
		return "The OpenAPI document declares no operationIds"
	}
	names := make([]string, 0, len(ids))
	for name := range ids {
		names = append(names, name)
	}
	sort.Strings(names)
	return "Declared operationIds: " + strings.Join(names, ", ")
}
