//ff:func feature=migration type=test control=sequence
//ff:what TestCreateTable_Destructive — 신규 생성은 항상 false
package migration

import "testing"

func TestCreateTable_Destructive(t *testing.T) {
	if (CreateTable{Table: &Table{Name: "x"}}).Destructive() {
		t.Error("CreateTable.Destructive() = true, want false")
	}
}
