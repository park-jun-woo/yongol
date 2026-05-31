//ff:func feature=migration type=test control=sequence
//ff:what TestCheckSafetyForOp — Operation 타입별 안전 점검 디스패치
package migration

import (
	"testing"
)

func TestCheckSafetyForOp(t *testing.T) {
	// AddColumn NOT NULL without default → MIG-002
	if got := checkSafetyForOp(AddColumn{Table: "u", Column: &Column{Name: "c", Nullable: false}}); len(got) != 1 || got[0].RuleID != "MIG-002" {
		t.Errorf("AddColumn dispatch = %+v, want MIG-002", got)
	}
	// DropTable without allow_destructive → MIG-004
	if got := checkSafetyForOp(DropTable{Name: "u"}); len(got) != 1 || got[0].RuleID != "MIG-004" {
		t.Errorf("DropTable dispatch = %+v, want MIG-004", got)
	}
	// Unhandled op type → nil
	if got := checkSafetyForOp(CreateTable{Table: &Table{Name: "x"}}); got != nil {
		t.Errorf("CreateTable dispatch = %+v, want nil", got)
	}
}
