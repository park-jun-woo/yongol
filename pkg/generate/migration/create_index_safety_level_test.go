//ff:func feature=migration type=test control=sequence
//ff:what TestCreateIndex_SafetyLevel — 항상 SafetySafe
package migration

import "testing"

func TestCreateIndex_SafetyLevel(t *testing.T) {
	if got := (CreateIndex{}).SafetyLevel(); got != SafetySafe {
		t.Errorf("CreateIndex.SafetyLevel() = %v, want SafetySafe", got)
	}
}
