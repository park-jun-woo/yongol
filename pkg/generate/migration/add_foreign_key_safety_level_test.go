//ff:func feature=migration type=test control=sequence
//ff:what TestAddForeignKey_SafetyLevel — 항상 SafetySafe
package migration

import "testing"

func TestAddForeignKey_SafetyLevel(t *testing.T) {
	if got := (AddForeignKey{}).SafetyLevel(); got != SafetySafe {
		t.Errorf("AddForeignKey.SafetyLevel() = %v, want SafetySafe", got)
	}
}
