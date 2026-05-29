//ff:func feature=migration type=test control=sequence
//ff:what TestDropIndex_SafetyLevel — 항상 SafetySafe
package migration

import "testing"

func TestDropIndex_SafetyLevel(t *testing.T) {
	if got := (DropIndex{Name: "ix"}).SafetyLevel(); got != SafetySafe {
		t.Errorf("DropIndex.SafetyLevel() = %v, want SafetySafe", got)
	}
}
