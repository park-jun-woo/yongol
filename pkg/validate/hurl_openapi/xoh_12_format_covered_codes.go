//ff:func feature=validate type=util control=iteration dimension=1 topic=hurl-openapi
//ff:what formatCoveredCodes — covered status code set 을 CSV 문자열로 포맷

package hurl_openapi

import (
	"sort"
	"strings"
)

func formatCoveredCodes(coveredSet map[string]bool) string {
	if len(coveredSet) == 0 {
		return "none"
	}
	var list []string
	for code := range coveredSet {
		list = append(list, code)
	}
	sort.Strings(list)
	return strings.Join(list, ", ")
}
