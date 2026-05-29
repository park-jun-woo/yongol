//ff:func feature=migration type=test control=sequence
//ff:what TestDiff_AlterColumnType — INTEGER → BIGINT 변경 시 AlterColumnType 생성
package migration

import "testing"

func TestDiff_AlterColumnType(t *testing.T) {
	prev := mustAST(t, `CREATE TABLE t (id INTEGER);`)
	curr := mustAST(t, `CREATE TABLE t (id BIGINT);`)
	ops := Diff(prev, curr, nil)
	if len(ops) != 1 {
		t.Fatalf("expected 1 op, got %d", len(ops))
	}
	if _, ok := ops[0].(AlterColumnType); !ok {
		t.Errorf("expected AlterColumnType, got %T", ops[0])
	}
}
