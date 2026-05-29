//ff:func feature=migration type=test control=sequence
//ff:what TestRenameTable_SQL — ALTER TABLE RENAME TO 렌더
package migration

import "testing"

func TestRenameTable_SQL(t *testing.T) {
	if got, want := (RenameTable{From: "old", To: "new"}).SQL(), "ALTER TABLE old RENAME TO new;"; got != want {
		t.Errorf("SQL() = %q, want %q", got, want)
	}
}
