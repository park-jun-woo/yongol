//ff:func feature=migration type=test control=sequence
//ff:what TestDropColumn_Destructive — DROP COLUMN 은 항상 파괴적
package migration

import "testing"

func TestDropColumn_Destructive(t *testing.T) {
	if !(DropColumn{}).Destructive() {
		t.Error("DropColumn.Destructive() = false, want true")
	}
}
