//ff:func feature=migration type=test control=iteration dimension=1
//ff:what CheckSafety / ApplyHintsToOps — 파괴성 분류 + hint 소비
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

func TestCheckSafety_DropTable_Warning(t *testing.T) {
	ops := []Operation{DropTable{Name: "old_t"}}
	issues := CheckSafety(ops)
	if len(issues) != 1 || issues[0].RuleID != "MIG-004" {
		t.Errorf("expected MIG-004, got %+v", issues)
	}
}

func TestCheckSafety_DropTable_AllowDestructive_Silences(t *testing.T) {
	ops := []Operation{DropTable{Name: "old_t"}}
	hints := &Hints{AllowDestructive: map[string]bool{"old_t": true}}
	ops = ApplyHintsToOps(ops, hints)
	issues := CheckSafety(ops)
	if len(issues) != 0 {
		t.Errorf("expected no issues with @allow_destructive, got %+v", issues)
	}
}

func TestCheckSafety_RiskyCast(t *testing.T) {
	ops := []Operation{
		AlterColumnType{
			Table: "t", Column: "c",
			From: CanonicalType{Base: "INTEGER"},
			To:   CanonicalType{Base: "TEXT"},
		},
	}
	issues := CheckSafety(ops)
	if len(issues) != 1 || issues[0].RuleID != "MIG-005" {
		t.Errorf("expected MIG-005, got %+v", issues)
	}
}

func TestCheckSafety_CastHint_Silences(t *testing.T) {
	ops := []Operation{
		AlterColumnType{
			Table: "t", Column: "c",
			From: CanonicalType{Base: "INTEGER"},
			To:   CanonicalType{Base: "TEXT"},
		},
	}
	hints := &Hints{
		Casts: map[colKey]string{
			{Table: "t", Column: "c"}: "c::text",
		},
	}
	ops = ApplyHintsToOps(ops, hints)
	issues := CheckSafety(ops)
	if len(issues) != 0 {
		t.Errorf("expected no issues with @cast, got %+v", issues)
	}
}
