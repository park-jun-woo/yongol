//ff:func feature=migration type=test control=sequence
//ff:what TestCheckSafety_NotNullWithoutBackfill — MIG-002 SafetyError 생성 검증
package migration

import "testing"

func TestCheckSafety_NotNullWithoutBackfill(t *testing.T) {
	ops := []Operation{
		AlterColumnNullable{Table: "users", Column: "email_verified", From: true, To: false},
	}
	issues := CheckSafety(ops)
	if len(issues) != 1 || issues[0].RuleID != "MIG-002" {
		t.Errorf("expected MIG-002 issue, got %+v", issues)
	}
	if issues[0].Level != SafetyError {
		t.Errorf("expected SafetyError, got %v", issues[0].Level)
	}
}
