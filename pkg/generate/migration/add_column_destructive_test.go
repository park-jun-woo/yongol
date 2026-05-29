//ff:func feature=migration type=test control=sequence
//ff:what TestAddColumn_Destructive — 신규 컬럼 추가는 항상 false
package migration

import "testing"

func TestAddColumn_Destructive(t *testing.T) {
	if (AddColumn{Column: &Column{}}).Destructive() {
		t.Error("AddColumn.Destructive() = true, want false")
	}
}
