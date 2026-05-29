//ff:func feature=migration type=test control=sequence
//ff:what TestRenameColumn_SafetyLevel — 항상 SafetySafe
package migration

import "testing"

func TestRenameColumn_SafetyLevel(t *testing.T) {
	if got := (RenameColumn{}).SafetyLevel(); got != SafetySafe {
		t.Errorf("RenameColumn.SafetyLevel() = %v, want SafetySafe", got)
	}
}
