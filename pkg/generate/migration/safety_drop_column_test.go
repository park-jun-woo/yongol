//ff:func feature=migration type=test control=sequence
//ff:what TestSafetyDropColumn — @allow_destructive 없으면 MIG-004 WARNING
package migration

import "testing"

func TestSafetyDropColumn(t *testing.T) {
	if got := safetyDropColumn(DropColumn{Table: "u", Column: "c", AllowDestructive: true}); got != nil {
		t.Errorf("with allow_destructive want nil, got %v", got)
	}
	issues := safetyDropColumn(DropColumn{Table: "u", Column: "c"})
	if len(issues) != 1 || issues[0].RuleID != "MIG-004" || issues[0].Level != SafetyWarning {
		t.Errorf("got %+v, want one MIG-004 warning", issues)
	}
}
