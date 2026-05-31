//ff:func feature=agent type=test-helper control=iteration dimension=1
//ff:what errorResponseCodes — 응답 맵의 정렬된 status 코드 키 목록 반환 헬퍼
package agent

import "sort"

// errorResponseCodes returns the sorted keys of m.
func errorResponseCodes(m map[string]any) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}
