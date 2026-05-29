//ff:func feature=migration type=test control=sequence
//ff:what TestRenameColumn_SQL — ALTER TABLE RENAME COLUMN 렌더
package migration

import "testing"

func TestRenameColumn_SQL(t *testing.T) {
	op := RenameColumn{Table: "users", From: "old", To: "new"}
	if got, want := op.SQL(), "ALTER TABLE users RENAME COLUMN old TO new;"; got != want {
		t.Errorf("SQL() = %q, want %q", got, want)
	}
}
