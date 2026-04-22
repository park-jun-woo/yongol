//ff:func feature=migration type=test control=iteration dimension=1
//ff:what Diff 컬럼 매트릭스 — 추가/삭제/타입/Nullable/Default
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

func TestDiff_DropColumn(t *testing.T) {
	prev := mustAST(t, `CREATE TABLE users (id BIGSERIAL PRIMARY KEY, name TEXT);`)
	curr := mustAST(t, `CREATE TABLE users (id BIGSERIAL PRIMARY KEY);`)
	ops := Diff(prev, curr, nil)
	if len(ops) != 1 {
		t.Fatalf("expected 1 op, got %d", len(ops))
	}
	if _, ok := ops[0].(DropColumn); !ok {
		t.Errorf("expected DropColumn, got %T", ops[0])
	}
	if !ops[0].Destructive() {
		t.Errorf("DropColumn should be destructive")
	}
	if ops[0].SafetyLevel() != SafetyWarning {
		t.Errorf("DropColumn without allow_destructive should WARN")
	}
}

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

func TestDiff_EmptyDiff(t *testing.T) {
	sql := `CREATE TABLE t (id BIGSERIAL PRIMARY KEY, name TEXT NOT NULL);`
	ops := Diff(mustAST(t, sql), mustAST(t, sql), nil)
	if len(ops) != 0 {
		t.Errorf("expected no ops, got %d: %+v", len(ops), ops)
	}
}

func TestDiff_Determinism(t *testing.T) {
	prev := mustAST(t, `CREATE TABLE t (id INTEGER);`)
	curr := mustAST(t, `CREATE TABLE t (id INTEGER, name TEXT, age INTEGER);`)
	a := Diff(prev, curr, nil)
	b := Diff(prev, curr, nil)
	if len(a) != len(b) {
		t.Fatalf("different lengths: %d vs %d", len(a), len(b))
	}
	for i := range a {
		if a[i].Description() != b[i].Description() {
			t.Errorf("order differs at %d: %q vs %q", i, a[i].Description(), b[i].Description())
		}
	}
}
