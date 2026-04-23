//ff:func feature=migration type=util control=selection
//ff:what splitState.stepInBC — 블록 코멘트 안에서 `*/` 감지하면 종료
package migration

// stepInBC handles one byte while inside a /* block comment */.
func (st *splitState) stepInBC(s string, i int, c byte) int {
	st.sb.WriteByte(c)
	if c == '*' && i+1 < len(s) && s[i+1] == '/' {
		st.sb.WriteByte('/')
		st.inBC = false
		return i + 1
	}
	return i
}
