//ff:func feature=migration type=test control=sequence
//ff:what TestAlterColumnDefault_SafetyLevel — 항상 SafetySafe
package migration

import "testing"

func TestAlterColumnDefault_SafetyLevel(t *testing.T) {
	if got := (AlterColumnDefault{}).SafetyLevel(); got != SafetySafe {
		t.Errorf("SafetyLevel() = %v, want SafetySafe", got)
	}
}
