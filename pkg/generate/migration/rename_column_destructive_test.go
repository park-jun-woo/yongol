//ff:func feature=migration type=test control=sequence
//ff:what TestRenameColumn_Destructive — rename 은 비파괴적(false)
package migration

import "testing"

func TestRenameColumn_Destructive(t *testing.T) {
	if (RenameColumn{}).Destructive() {
		t.Error("RenameColumn.Destructive() = true, want false")
	}
}
