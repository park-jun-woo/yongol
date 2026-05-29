//ff:func feature=migration type=test control=sequence
//ff:what TestDropCheck_SQL — ALTER TABLE ... DROP CONSTRAINT ...;
package migration

import "testing"

func TestDropCheck_SQL(t *testing.T) {
	op := DropCheck{Table: "users", Name: "users_age_check"}
	want := "ALTER TABLE users DROP CONSTRAINT users_age_check;"
	if got := op.SQL(); got != want {
		t.Errorf("SQL() = %q, want %q", got, want)
	}
}
