//ff:func feature=migration type=util control=iteration dimension=1
//ff:what sortedMapKeys — 문자열 키 맵의 키 슬라이스를 정렬해 반환 (제네릭)
package migration

import "sort"

// sortedMapKeys returns the keys of a string-keyed map in sorted order.
func sortedMapKeys[V any](m map[string]V) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}
