//ff:func feature=agent type=test-helper control=iteration dimension=1
//ff:what equalStrings — 두 문자열 슬라이스의 순서까지 동일 여부 비교 헬퍼
package agent

// equalStrings reports whether a and b have equal elements in order.
func equalStrings(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
