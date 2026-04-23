//ff:func feature=migration type=test control=sequence
//ff:what TestDiff_AlterColumnDefault — DEFAULT 값 변경 시 AlterColumnDefault 생성
package migration

import "testing"

func TestDiff_AlterColumnDefault(t *testing.T) {
	prev := mustAST(t, `CREATE TABLE t (role VARCHAR(32) DEFAULT 'member');`)
	curr := mustAST(t, `CREATE TABLE t (role VARCHAR(32) DEFAULT 'guest');`)
	ops := Diff(prev, curr, nil)
	if len(ops) != 1 {
		t.Fatalf("expected 1 op, got %d", len(ops))
	}
	if _, ok := ops[0].(AlterColumnDefault); !ok {
		t.Errorf("expected AlterColumnDefault, got %T", ops[0])
	}
}
