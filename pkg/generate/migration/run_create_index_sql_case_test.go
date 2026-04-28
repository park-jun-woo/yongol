//ff:func feature=migration type=test-helper control=sequence
//ff:what runCreateIndexSQLCase — 단일 CreateIndex.SQL() 서브테스트 실행

package migration

import (
	"strings"
	"testing"
)

// runCreateIndexSQLCase executes one CreateIndex.SQL emit case as a
// t.Run subtest. Keeps TestCreateIndex_SQL_EmitsUsing at depth 1 (range)
// instead of range→t.Run(closure)→if (depth 3).
func runCreateIndexSQLCase(t *testing.T, name string, idx *Index, substr, notHas string) {
	t.Helper()
	t.Run(name, func(t *testing.T) {
		op := CreateIndex{Table: "t", Index: idx}
		got := op.SQL()
		if substr != "" && !strings.Contains(got, substr) {
			t.Errorf("SQL = %q, want substring %q", got, substr)
		}
		if notHas != "" && strings.Contains(got, notHas) {
			t.Errorf("SQL = %q, unexpected substring %q", got, notHas)
		}
	})
}
