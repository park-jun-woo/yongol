//ff:func feature=migration type=parser control=iteration dimension=1
//ff:what splitStatements — SQL 문자열을 top-level `;` 기준으로 분할 (따옴표·괄호·주석 고려)
package migration

// splitStatements divides SQL text on top-level `;` boundaries while
// respecting single-quoted strings, double-quoted identifiers and
// block comments. Line (`--`) comments are stripped before splitting.
func splitStatements(sql string) []string {
	clean := stripLineComments(sql)
	st := newSplitState()
	for i := 0; i < len(clean); i++ {
		i = st.step(clean, i)
	}
	return st.finish()
}
