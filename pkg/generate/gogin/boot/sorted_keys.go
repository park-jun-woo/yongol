//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what sortedKeys — map[string]int64 의 key 리스트를 문자열 순으로 정렬

package boot

import "sort"

// sortedKeys returns the keys of m sorted deterministically for codegen.
func sortedKeys(m map[string]int64) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
