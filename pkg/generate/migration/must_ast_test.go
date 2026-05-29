//ff:func feature=migration type=test-helper control=sequence
//ff:what mustAST — SQL 문자열을 테스트용 AST 로 파싱 (실패 시 t.Fatal)
package migration

import "testing"

func mustAST(t *testing.T, sql string) *Schema {
	t.Helper()
	s := NewSchema()
	if err := BuildASTFromSQL(s, sql); err != nil {
		t.Fatalf("BuildAST: %v", err)
	}
	return s
}
