//ff:func feature=migration type=test control=sequence
//ff:what TestDiff_AddColumn — curr 에만 있는 컬럼이면 AddColumn 생성
package migration

import "testing"

func TestDiff_AddColumn(t *testing.T) {
	prev := mustAST(t, `CREATE TABLE users (id BIGSERIAL PRIMARY KEY);`)
	curr := mustAST(t, `CREATE TABLE users (id BIGSERIAL PRIMARY KEY, name TEXT);`)
	ops := Diff(prev, curr, nil)
	if len(ops) != 1 {
		t.Fatalf("expected 1 op, got %d: %+v", len(ops), ops)
	}
	if _, ok := ops[0].(AddColumn); !ok {
		t.Errorf("expected AddColumn, got %T", ops[0])
	}
}
