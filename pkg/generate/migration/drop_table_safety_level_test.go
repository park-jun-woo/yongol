//ff:func feature=migration type=test control=selection dimension=2
//ff:what TestDropTable_SafetyLevel — allow_destructive 없으면 Warning
package migration

import "testing"

func TestDropTable_SafetyLevel(t *testing.T) {
	if got := (DropTable{AllowDestructive: true}).SafetyLevel(); got != SafetySafe {
		t.Errorf("AllowDestructive=true: %v, want SafetySafe", got)
	}
	if got := (DropTable{AllowDestructive: false}).SafetyLevel(); got != SafetyWarning {
		t.Errorf("AllowDestructive=false: %v, want SafetyWarning", got)
	}
}
