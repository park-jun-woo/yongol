//ff:func feature=gen-gogin type=test-helper control=iteration dimension=1
//ff:what equalStrings — 두 []string의 순서 포함 동치 비교 (길이 + 각 원소 일치)
package boot

// equalStrings returns true when a and b have the same length and every
// element at the same index is equal. Used by boot tests for order-sensitive
// slice comparison.
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
