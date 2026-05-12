//ff:func feature=gen-react type=util control=iteration dimension=1
//ff:what string 맵의 키를 정렬된 슬라이스로 반환한다

package react

import "sort"

// sortedKeys returns map keys in sorted order for deterministic output.
func sortedKeys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
