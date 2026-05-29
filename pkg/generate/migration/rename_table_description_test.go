//ff:func feature=migration type=test control=sequence
//ff:what TestRenameTable_Description — from → to 표기 확인
package migration

import "testing"

func TestRenameTable_Description(t *testing.T) {
	if got, want := (RenameTable{From: "old", To: "new"}).Description(), "rename table old → new"; got != want {
		t.Errorf("Description() = %q, want %q", got, want)
	}
}
