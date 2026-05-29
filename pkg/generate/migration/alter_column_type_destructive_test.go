//ff:func feature=migration type=test control=sequence
//ff:what TestAlterColumnType_Destructive — 타입 변경은 항상 파괴적
package migration

import "testing"

func TestAlterColumnType_Destructive(t *testing.T) {
	if !(AlterColumnType{}).Destructive() {
		t.Error("AlterColumnType.Destructive() = false, want true")
	}
}
