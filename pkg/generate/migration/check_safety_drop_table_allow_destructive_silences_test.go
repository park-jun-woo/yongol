//ff:func feature=migration type=test control=sequence
//ff:what TestCheckSafety_DropTable_AllowDestructive_Silences — @allow_destructive 힌트 적용 시 MIG-004 미발생
package migration

import "testing"

func TestCheckSafety_DropTable_AllowDestructive_Silences(t *testing.T) {
	ops := []Operation{DropTable{Name: "old_t"}}
	hints := &Hints{AllowDestructive: map[string]bool{"old_t": true}}
	ops = ApplyHintsToOps(ops, hints)
	issues := CheckSafety(ops)
	if len(issues) != 0 {
		t.Errorf("expected no issues with @allow_destructive, got %+v", issues)
	}
}
