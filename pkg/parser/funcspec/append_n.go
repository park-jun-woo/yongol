//ff:func feature=funcspec type=util control=iteration dimension=1
//ff:what appendN — 같은 문자열을 n 회 반복 append (extractReturnTypes 보조)

package funcspec

// appendN appends value n times to dst. Used by extractReturnTypes to expand
// a named-result group like `(a, b int)` into one entry per name without
// nesting another loop inside the caller.
func appendN(dst []string, value string, n int) []string {
	for i := 0; i < n; i++ {
		dst = append(dst, value)
	}
	return dst
}
