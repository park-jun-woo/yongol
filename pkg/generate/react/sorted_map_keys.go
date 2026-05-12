//ff:func feature=gen-react type=util control=iteration dimension=1
//ff:what string 맵의 키를 정렬된 슬라이스로 반환한다

package react

import "sort"

// sortedMapKeys returns the keys of a string map in sorted order.
func sortedMapKeys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
