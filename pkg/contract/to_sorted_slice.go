//ff:func feature=contract type=util control=iteration dimension=1
//ff:what toSortedSlice — map 키를 정렬된 슬라이스로 반환 (결정적 출력)

package contract

import "sort"

// toSortedSlice returns the keys of m in lexical order so output is
// deterministic and diff-friendly. A nil or empty map yields nil so
// callers can JSON-serialize without emitting `[]` placeholders.
func toSortedSlice(m map[string]struct{}) []string {
	if len(m) == 0 {
		return nil
	}
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}
