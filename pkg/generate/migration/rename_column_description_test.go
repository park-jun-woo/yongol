//ff:func feature=migration type=test control=sequence
//ff:what TestRenameColumn_Description — table.from → to 표기 확인
package migration

import "testing"

func TestRenameColumn_Description(t *testing.T) {
	op := RenameColumn{Table: "users", From: "old", To: "new"}
	if got, want := op.Description(), "rename column users.old → new"; got != want {
		t.Errorf("Description() = %q, want %q", got, want)
	}
}
