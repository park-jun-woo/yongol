//ff:func feature=migration type=test control=sequence
//ff:what TestDropColumn_SQL — DROP COLUMN 문장 렌더 확인
package migration

import "testing"

func TestDropColumn_SQL(t *testing.T) {
	op := DropColumn{Table: "users", Column: "age"}
	if got, want := op.SQL(), "ALTER TABLE users DROP COLUMN age;"; got != want {
		t.Errorf("SQL() = %q, want %q", got, want)
	}
}
