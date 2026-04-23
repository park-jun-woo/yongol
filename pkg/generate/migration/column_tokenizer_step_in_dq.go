//ff:func feature=migration type=util control=selection
//ff:what columnTokenizer.stepInDQ — double-quoted 식별자 안 처리
package migration

func (st *columnTokenizer) stepInDQ(i int, c byte) int {
	st.sb.WriteByte(c)
	if c == '"' {
		st.inDQ = false
	}
	return i
}
