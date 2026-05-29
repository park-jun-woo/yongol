//ff:func feature=migration type=test control=sequence
//ff:what TestSafetyDropTable — @allow_destructive 없으면 MIG-004 WARNING
package migration

import "testing"

func TestSafetyDropTable(t *testing.T) {
	if got := safetyDropTable(DropTable{Name: "users", AllowDestructive: true}); got != nil {
		t.Errorf("with allow_destructive want nil, got %v", got)
	}
	issues := safetyDropTable(DropTable{Name: "users"})
	if len(issues) != 1 {
		t.Fatalf("want 1 issue, got %d", len(issues))
	}
	if issues[0].RuleID != "MIG-004" || issues[0].Level != SafetyWarning {
		t.Errorf("got %+v, want MIG-004 warning", issues[0])
	}
}
