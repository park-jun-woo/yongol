//ff:func feature=migration type=test control=sequence
//ff:what diff_helpers_unit_test — checkDiffForOne/findPrevViaRenameHint/checkDropOps/checkAlterOrAddOps/allCreateTableOps/createTableWithDeps/parseTableFKRef 단위 테스트
package migration

import (
	"testing"
)

func TestCheckDiffForOne(t *testing.T) {
	cur := map[string]*CheckConstraint{"c": {Name: "c", Expression: "x > 0"}}

	// added (not in prev)
	ops := checkDiffForOne("t", "c", map[string]*CheckConstraint{}, cur)
	if len(ops) != 1 {
		t.Fatalf("added: got %d ops, want 1", len(ops))
	}
	if _, ok := ops[0].(AddCheck); !ok {
		t.Errorf("added: op0 should be AddCheck, got %T", ops[0])
	}

	// unchanged -> no ops
	prevSame := map[string]*CheckConstraint{"c": {Name: "c", Expression: "x > 0"}}
	if ops := checkDiffForOne("t", "c", prevSame, cur); ops != nil {
		t.Errorf("unchanged: expected nil ops, got %#v", ops)
	}

	// changed expression -> Drop + Add
	prevDiff := map[string]*CheckConstraint{"c": {Name: "c", Expression: "x > 5"}}
	ops = checkDiffForOne("t", "c", prevDiff, cur)
	if len(ops) != 2 {
		t.Fatalf("changed: got %d ops, want 2", len(ops))
	}
	if _, ok := ops[0].(DropCheck); !ok {
		t.Errorf("changed: op0 should be DropCheck, got %T", ops[0])
	}
	if _, ok := ops[1].(AddCheck); !ok {
		t.Errorf("changed: op1 should be AddCheck, got %T", ops[1])
	}
}
