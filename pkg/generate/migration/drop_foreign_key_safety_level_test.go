//ff:func feature=migration type=test control=sequence
//ff:what TestDropForeignKey_SafetyLevel — 항상 SafetySafe
package migration

import "testing"

func TestDropForeignKey_SafetyLevel(t *testing.T) {
	if got := (DropForeignKey{}).SafetyLevel(); got != SafetySafe {
		t.Errorf("DropForeignKey.SafetyLevel() = %v, want SafetySafe", got)
	}
}
