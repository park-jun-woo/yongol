//ff:func feature=migration type=test control=sequence
//ff:what TestDropTable_Destructive — DROP TABLE 은 항상 파괴적
package migration

import "testing"

func TestDropTable_Destructive(t *testing.T) {
	if !(DropTable{}).Destructive() {
		t.Error("DropTable.Destructive() = false, want true")
	}
}
