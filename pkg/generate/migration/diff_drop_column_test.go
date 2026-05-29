//ff:func feature=migration type=test control=sequence
//ff:what TestDiff_DropColumn — prev 에만 있는 컬럼이면 DropColumn 생성 + WARN 분류
package migration

import "testing"

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
