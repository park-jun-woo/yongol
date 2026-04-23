//ff:func feature=migration type=util control=selection
//ff:what updateParenDepth — '(' → +1, ')' → -1, 그 외 → depth 유지
package migration

// updateParenDepth returns the new paren depth after consuming c.
func updateParenDepth(c byte, depth int) int {
	switch c {
	case '(':
		return depth + 1
	case ')':
		return depth - 1
	}
	return depth
}
