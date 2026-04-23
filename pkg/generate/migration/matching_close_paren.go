//ff:func feature=migration type=util control=iteration dimension=1
//ff:what matchingCloseParen — 첫 "(" 의 짝 맞는 ")" 인덱스, 없으면 -1
package migration

// matchingCloseParen assumes s starts with `(` and returns the index of
// the paired `)`, or -1 if none found.
func matchingCloseParen(s string) int {
	depth := 0
	for i := 0; i < len(s); i++ {
		depth = updateParenDepth(s[i], depth)
		if s[i] == ')' && depth == 0 {
			return i
		}
	}
	return -1
}
