//ff:func feature=migration type=util control=selection
//ff:what columnTokenizer.step — 한 바이트 처리 (따옴표/괄호/공백 분기)
package migration

// step consumes one byte at s[i] and returns the next index to read.
func (st *columnTokenizer) step(s string, i int) int {
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
	case isSpaceByte(c) && st.depth == 0:
		st.flush()
	default:
		st.sb.WriteByte(c)
	}
	return i
}
