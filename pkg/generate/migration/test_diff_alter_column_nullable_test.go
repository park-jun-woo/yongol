//ff:func feature=migration type=test control=sequence
//ff:what TestDiff_AlterColumnNullable — NOT NULL 추가 시 AlterColumnNullable + SafetyError
package migration

import "testing"

func TestDiff_AlterColumnNullable(t *testing.T) {
	prev := mustAST(t, `CREATE TABLE t (id INTEGER);`)
	curr := mustAST(t, `CREATE TABLE t (id INTEGER NOT NULL);`)
	ops := Diff(prev, curr, nil)
	if len(ops) != 1 {
		t.Fatalf("expected 1 op, got %d", len(ops))
	}
	aop, ok := ops[0].(AlterColumnNullable)
	if !ok {
		t.Fatalf("expected AlterColumnNullable, got %T", ops[0])
	}
	if aop.To != false {
		t.Errorf("expected To=false (NOT NULL), got %v", aop.To)
	}
	if aop.SafetyLevel() != SafetyError {
		t.Errorf("NOT NULL add without backfill must ERROR, got %v", aop.SafetyLevel())
	}
}
