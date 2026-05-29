//ff:func feature=migration type=test control=selection dimension=2
//ff:what TestAlterColumnType_SafetyLevel — USING 있으면 Safe, 없으면 Warning
package migration

import "testing"

func TestAlterColumnType_SafetyLevel(t *testing.T) {
	if got := (AlterColumnType{Using: "c::int"}).SafetyLevel(); got != SafetySafe {
		t.Errorf("with USING: %v, want SafetySafe", got)
	}
	if got := (AlterColumnType{}).SafetyLevel(); got != SafetyWarning {
		t.Errorf("without USING: %v, want SafetyWarning", got)
	}
}
