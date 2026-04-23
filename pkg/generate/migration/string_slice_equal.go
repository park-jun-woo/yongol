//ff:func feature=migration type=util control=iteration dimension=1
//ff:what stringSliceEqual — 두 문자열 슬라이스의 요소별 동등 비교
package migration

// stringSliceEqual reports whether a and b have the same length and
// pairwise-equal elements.
func stringSliceEqual(a, b []string) bool {
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
