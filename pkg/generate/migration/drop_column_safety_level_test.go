//ff:func feature=migration type=test control=sequence
//ff:what TestDropColumn_SafetyLevel — allow_destructive 없으면 Warning
package migration

import (
	"testing"
)

func TestDropColumn_SafetyLevel(t *testing.T) {
	if got := (DropColumn{AllowDestructive: true}).SafetyLevel(); got != SafetySafe {
		t.Errorf("AllowDestructive=true: %v, want SafetySafe", got)
	}
	if got := (DropColumn{AllowDestructive: false}).SafetyLevel(); got != SafetyWarning {
		t.Errorf("AllowDestructive=false: %v, want SafetyWarning", got)
	}
}
