//ff:func feature=migration type=test control=sequence
//ff:what TestRenameTable_SafetyLevel — 항상 SafetySafe
package migration

import "testing"

func TestRenameTable_SafetyLevel(t *testing.T) {
	if got := (RenameTable{}).SafetyLevel(); got != SafetySafe {
		t.Errorf("RenameTable.SafetyLevel() = %v, want SafetySafe", got)
	}
}
