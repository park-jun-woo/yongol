//ff:func feature=gen-hurl type=test-helper control=iteration dimension=1
//ff:what indexOfString — []string 슬라이스에서 want 의 0-based index (없으면 -1) 반환

package hurl

// indexOfString returns the 0-based index of want in xs, or -1 when
// absent. Tests use it to assert ordering between two operationIds
// without worrying about total length.
func indexOfString(xs []string, want string) int {
	for i, x := range xs {
		if x == want {
			return i
		}
	}
	return -1
}
