//ff:func feature=validate type=util control=iteration dimension=1 topic=stml-design
//ff:what sortedKeys — map[string]string의 키를 정렬된 슬라이스로 반환
package stml_design

import (
	"sort"
)

// sortedKeys returns the keys of a map[string]string in sorted order.
func sortedKeys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
