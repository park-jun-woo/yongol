//ff:func feature=migration type=test control=sequence
//ff:what TestAddCheck_SQL — ADD CONSTRAINT CHECK 문장 렌더 확인
package migration

import "testing"

func TestAddCheck_SQL(t *testing.T) {
	op := AddCheck{Table: "users", Check: &CheckConstraint{Name: "users_age_check", Expression: "age >= 0"}}
	want := "ALTER TABLE users ADD CONSTRAINT users_age_check CHECK (age >= 0);"
	if got := op.SQL(); got != want {
		t.Errorf("SQL() = %q, want %q", got, want)
	}
}
