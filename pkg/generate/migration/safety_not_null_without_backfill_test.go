//ff:func feature=migration type=test control=sequence
//ff:what TestSafetyNotNullWithoutBackfill — SET NOT NULL + backfill 없으면 MIG-002 ERROR
package migration

import "testing"

func TestSafetyNotNullWithoutBackfill(t *testing.T) {
	if got := safetyNotNullWithoutBackfill(AlterColumnNullable{Table: "u", Column: "c", To: true}); got != nil {
		t.Errorf("SET NULL (To=true) want nil, got %v", got)
	}
	if got := safetyNotNullWithoutBackfill(AlterColumnNullable{Table: "u", Column: "c", To: false, Backfill: "0"}); got != nil {
		t.Errorf("with backfill want nil, got %v", got)
	}
	issues := safetyNotNullWithoutBackfill(AlterColumnNullable{Table: "u", Column: "c", To: false})
	if len(issues) != 1 || issues[0].RuleID != "MIG-002" || issues[0].Level != SafetyError {
		t.Errorf("got %+v, want one MIG-002 error", issues)
	}
}
