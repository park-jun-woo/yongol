//ff:func feature=migration type=parser control=iteration dimension=1
//ff:what tokenizeColumnDef — 컬럼 정의를 토큰화 (괄호/따옴표는 한 토큰으로 보존)
package migration

// tokenizeColumnDef splits a column definition into whitespace-separated
// tokens while keeping parenthesised groups and quoted strings as a
// single token each.
func tokenizeColumnDef(s string) []string {
	st := newColumnTokenizer()
	for i := 0; i < len(s); i++ {
		i = st.step(s, i)
	}
	return st.finish()
}
