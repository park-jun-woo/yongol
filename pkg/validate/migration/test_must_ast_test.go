//ff:func feature=validate type=test-helper control=sequence topic=migration-hints
//ff:what 테스트 헬퍼 — SQL 문자열을 migration.Schema AST 로 파싱

package migration

import (
	"testing"

	migration "github.com/park-jun-woo/yongol/pkg/generate/migration"
)

func mustAST(t *testing.T, sql string) *migration.Schema {
	t.Helper()
	s := migration.NewSchema()
	if err := migration.BuildASTFromSQL(s, sql); err != nil {
		t.Fatalf("build: %v", err)
	}
	return s
}
