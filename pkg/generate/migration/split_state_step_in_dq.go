//ff:func feature=migration type=util control=selection
//ff:what splitState.stepInDQ — double-quote 식별자 안 처리
package migration

// stepInDQ handles one byte while inside a "..." identifier.
func (st *splitState) stepInDQ(i int, c byte) int {
	st.sb.WriteByte(c)
	if c == '"' {
		st.inDQ = false
	}
	return i
}
