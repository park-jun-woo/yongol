//ff:func feature=validate type=util control=iteration dimension=1 topic=ssac-openapi
//ff:what XOS-80/82 helper — int 키 맵을 오름차순으로 정렬해 결정적 출력 제공

package openapi_ssac

import "sort"

// sortedKeys returns the int keys of m in ascending order — used only by
// diagnostic messages so the output is deterministic across runs.
func sortedKeys(m map[int]bool) []int {
	out := make([]int, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	sort.Ints(out)
	return out
}
