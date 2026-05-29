//ff:func feature=migration type=test control=sequence
//ff:what TestDiff_CreateTable — 빈 prev vs 새 테이블 curr 시 CreateTable 생성
package migration

import "testing"

func TestDiff_CreateTable(t *testing.T) {
	prev := NewSchema()
	curr := mustAST(t, `CREATE TABLE users (id BIGSERIAL PRIMARY KEY);`)
	ops := Diff(prev, curr, nil)
	if len(ops) == 0 {
		t.Fatalf("expected CreateTable op, got none")
	}
	if _, ok := ops[0].(CreateTable); !ok {
		t.Errorf("op[0] should be CreateTable, got %T", ops[0])
	}
}
