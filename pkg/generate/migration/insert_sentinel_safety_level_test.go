//ff:func feature=migration type=test control=sequence
//ff:what TestInsertSentinel_SafetyLevel — 항상 SafetySafe
package migration

import "testing"

func TestInsertSentinel_SafetyLevel(t *testing.T) {
	if got := (InsertSentinel{Table: "roles"}).SafetyLevel(); got != SafetySafe {
		t.Errorf("InsertSentinel.SafetyLevel() = %v, want SafetySafe", got)
	}
}
