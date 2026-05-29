//ff:func feature=migration type=test control=sequence
//ff:what TestAddCheck_SafetyLevel — 항상 SafetySafe
package migration

import "testing"

func TestAddCheck_SafetyLevel(t *testing.T) {
	if got := (AddCheck{}).SafetyLevel(); got != SafetySafe {
		t.Errorf("SafetyLevel() = %v, want SafetySafe", got)
	}
}
