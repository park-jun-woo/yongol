//ff:func feature=migration type=test-helper control=iteration dimension=1
//ff:what countNonEmptyStmts — trim 후 비어있지 않은 문장 개수 카운트 헬퍼
package migration

// countNonEmptyStmts returns the number of statements in stmts whose trimmed
// form is non-empty.
func countNonEmptyStmts(stmts []string) int {
	n := 0
	for _, s := range stmts {
		if len(trimSpaceSimple(s)) > 0 {
			n++
		}
	}
	return n
}
