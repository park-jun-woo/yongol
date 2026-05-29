//ff:func feature=migration type=test control=sequence
//ff:what TestRenameTable_Destructive — rename 은 비파괴적(false)
package migration

import "testing"

func TestRenameTable_Destructive(t *testing.T) {
	if (RenameTable{}).Destructive() {
		t.Error("RenameTable.Destructive() = true, want false")
	}
}
