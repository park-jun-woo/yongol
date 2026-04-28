//ff:func feature=validate type=util control=iteration dimension=1 topic=hurl-openapi
//ff:what sortedKeys — bool map 의 key 를 정렬된 slice 로 반환

package hurl_openapi

import "sort"

// sortedKeys returns the keys of a bool map in deterministic order so
// diagnostics are stable in tests and CI diff.
func sortedKeys(m map[string]bool) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}
