//ff:func feature=migration type=test control=sequence
//ff:what TestDropCheck_SafetyLevel — 항상 SafetySafe
package migration

import "testing"

func TestDropCheck_SafetyLevel(t *testing.T) {
	if got := (DropCheck{}).SafetyLevel(); got != SafetySafe {
		t.Errorf("DropCheck.SafetyLevel() = %v, want SafetySafe", got)
	}
}
