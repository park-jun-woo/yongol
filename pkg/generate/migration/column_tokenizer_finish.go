//ff:func feature=migration type=util control=sequence
//ff:what columnTokenizer.finish — 마지막 토큰을 flush 한 뒤 out 반환
package migration

func (st *columnTokenizer) finish() []string {
	st.flush()
	return st.out
}
