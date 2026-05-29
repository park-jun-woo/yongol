//ff:func feature=migration type=test control=sequence
//ff:what TestCreateTable_SafetyLevel — 항상 SafetySafe
package migration

import "testing"

func TestCreateTable_SafetyLevel(t *testing.T) {
	if got := (CreateTable{Table: &Table{Name: "x"}}).SafetyLevel(); got != SafetySafe {
		t.Errorf("CreateTable.SafetyLevel() = %v, want SafetySafe", got)
	}
}
