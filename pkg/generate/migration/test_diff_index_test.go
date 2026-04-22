//ff:func feature=migration type=test control=iteration dimension=1
//ff:what Diff index — 추가/삭제/재생성
package migration

import "testing"

func TestDiff_Index_Add(t *testing.T) {
	prev := mustAST(t, `CREATE TABLE t (id BIGSERIAL PRIMARY KEY, name TEXT);`)
	curr := mustAST(t, `CREATE TABLE t (id BIGSERIAL PRIMARY KEY, name TEXT);
CREATE INDEX idx_t_name ON t (name);`)
	ops := Diff(prev, curr, nil)
	if len(ops) != 1 {
		t.Fatalf("expected 1 op, got %d: %+v", len(ops), ops)
	}
	if _, ok := ops[0].(CreateIndex); !ok {
		t.Errorf("expected CreateIndex, got %T", ops[0])
	}
}

func TestDiff_Index_Drop(t *testing.T) {
	prev := mustAST(t, `CREATE TABLE t (id BIGSERIAL PRIMARY KEY, name TEXT);
CREATE INDEX idx_t_name ON t (name);`)
	curr := mustAST(t, `CREATE TABLE t (id BIGSERIAL PRIMARY KEY, name TEXT);`)
	ops := Diff(prev, curr, nil)
	if len(ops) != 1 {
		t.Fatalf("expected 1 op, got %d", len(ops))
	}
	if _, ok := ops[0].(DropIndex); !ok {
		t.Errorf("expected DropIndex, got %T", ops[0])
	}
}
