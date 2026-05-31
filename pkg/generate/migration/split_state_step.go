//ff:func feature=migration type=util control=selection
//ff:what splitState.step — 한 바이트 처리해 i 의 다음 인덱스 반환 (lookahead 1)
package migration

// step consumes one byte at position i of s and returns the index to
// read next (may skip ahead for multi-char tokens like '/*' or ”').
func (st *splitState) step(s string, i int) int {
	c := s[i]
	switch {
	case st.inBC:
		return st.stepInBC(s, i, c)
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
	case c == '/' && i+1 < len(s) && s[i+1] == '*':
		st.inBC = true
		st.sb.WriteByte(c)
		st.sb.WriteByte('*')
		return i + 1
	case c == '(':
		st.depth++
		st.sb.WriteByte(c)
	case c == ')':
		if st.depth > 0 {
			st.depth--
		}
		st.sb.WriteByte(c)
	case c == ';' && st.depth == 0:
		st.out = append(st.out, st.sb.String())
		st.sb.Reset()
	default:
		st.sb.WriteByte(c)
	}
	return i
}
