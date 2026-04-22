//ff:func feature=orchestrator type=util control=iteration dimension=1
//ff:what sortedKeys — map[string]bool 의 key 를 정렬된 슬라이스로 반환
package sqlc

import "sort"

// sortedKeys returns the keys of m sorted lexicographically.
func sortedKeys(m map[string]bool) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
