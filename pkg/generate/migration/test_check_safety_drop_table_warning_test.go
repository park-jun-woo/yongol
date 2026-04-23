//ff:func feature=migration type=test control=sequence
//ff:what TestCheckSafety_DropTable_Warning — MIG-004 WARNING 생성 검증
package migration

import "testing"

func TestCheckSafety_DropTable_Warning(t *testing.T) {
	ops := []Operation{DropTable{Name: "old_t"}}
	issues := CheckSafety(ops)
	if len(issues) != 1 || issues[0].RuleID != "MIG-004" {
		t.Errorf("expected MIG-004, got %+v", issues)
	}
}
