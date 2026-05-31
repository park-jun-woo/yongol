//ff:func feature=migration type=test-helper control=sequence
//ff:what assertSplitStatementsCount — splitStatements 의 비어있지 않은 문장 개수 검증 헬퍼
package migration

import "testing"

// assertSplitStatementsCount asserts splitStatements(in) yields want non-empty
// trimmed statements.
func assertSplitStatementsCount(t *testing.T, in string, want int) {
	t.Helper()
	got := splitStatements(in)
	if n := countNonEmptyStmts(got); n != want {
		t.Errorf("splitStatements(%q) -> %d non-empty stmts, want %d: %#v", in, n, want, got)
	}
}
