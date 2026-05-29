//ff:func feature=migration type=test control=sequence
//ff:what TestCheckSafety_Backfill_ApplyHints_Clears — backfill 힌트 적용 시 MIG-002 진단이 사라지는지 검증
package migration

import "testing"

func TestCheckSafety_Backfill_ApplyHints_Clears(t *testing.T) {
	ops := []Operation{
		AlterColumnNullable{Table: "users", Column: "email_verified", From: true, To: false},
	}
	hints := &Hints{
		Backfills: map[colKey]string{
			{Table: "users", Column: "email_verified"}: "false",
		},
	}
	ops = ApplyHintsToOps(ops, hints)
	issues := CheckSafety(ops)
	if len(issues) != 0 {
		t.Errorf("expected no issues after backfill applied, got %+v", issues)
	}
}
