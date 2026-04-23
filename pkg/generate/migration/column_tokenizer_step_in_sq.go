//ff:func feature=migration type=util control=selection
//ff:what columnTokenizer.stepInSQ — single-quoted 문자열 안 처리 (''escape 지원)
package migration

func (st *columnTokenizer) stepInSQ(s string, i int, c byte) int {
	st.sb.WriteByte(c)
	if c != '\'' {
		return i
	}
	if i+1 < len(s) && s[i+1] == '\'' {
		st.sb.WriteByte('\'')
		return i + 1
	}
	st.inSQ = false
	return i
}
