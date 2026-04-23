//ff:func feature=migration type=util control=selection
//ff:what stepTopLevel — splitTopLevel 한 바이트 처리 (sep 감지 시 분할)
package migration

// stepTopLevel consumes one byte of s at i and, when c == sep at
// depth 0, splits the accumulator. Returns the next index.
func stepTopLevel(st *splitState, s string, i int, sep byte) int {
	c := s[i]
	switch {
	case st.inSQ:
		return st.stepInSQ(s, i, c)
	case st.inDQ:
		return st.stepInDQ(i, c)
	case c == '\'':
		st.inSQ = true
		st.sb.WriteByte(c)
	case c == '"':
		st.inDQ = true
		st.sb.WriteByte(c)
	case c == '(':
		st.depth++
		st.sb.WriteByte(c)
	case c == ')':
		st.depth--
		st.sb.WriteByte(c)
	case c == sep && st.depth == 0:
		st.out = append(st.out, st.sb.String())
		st.sb.Reset()
	default:
		st.sb.WriteByte(c)
	}
	return i
}
