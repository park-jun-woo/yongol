//ff:func feature=migration type=util control=sequence
//ff:what columnTokenizer.flush — 누적 바이트 버퍼를 out 에 추가 후 리셋
package migration

func (st *columnTokenizer) flush() {
	if st.sb.Len() > 0 {
		st.out = append(st.out, st.sb.String())
		st.sb.Reset()
	}
}
