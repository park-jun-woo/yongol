//ff:func feature=validate type=util control=iteration dimension=1 topic=ssac-structural
//ff:what s71ScopeList — scope map 을 정렬된 콤마 구분 문자열로 변환

package ssac

import (
	"sort"
	"strings"
)

func s71ScopeList(scope map[string]bool) string {
	keys := make([]string, 0, len(scope))
	for k := range scope {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return strings.Join(keys, ", ")
}
