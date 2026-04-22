//ff:func feature=validate type=util control=iteration dimension=1 topic=openapi-structural
//ff:what diffStringSets — 집합 a 에는 있고 b 에는 없는 원소를 정렬해 반환

package openapi

import "sort"

// diffStringSets returns sorted keys present in `a` but absent from `b`.
// Used by O-3 to compute missing/extra path parameters.
func diffStringSets(a, b map[string]bool) []string {
	var out []string
	for k := range a {
		if !b[k] {
			out = append(out, k)
		}
	}
	sort.Strings(out)
	return out
}
