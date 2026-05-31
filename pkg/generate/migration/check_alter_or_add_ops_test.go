//ff:func feature=migration type=test control=sequence
//ff:what diff_helpers_unit_test — checkDiffForOne/findPrevViaRenameHint/checkDropOps/checkAlterOrAddOps/allCreateTableOps/createTableWithDeps/parseTableFKRef 단위 테스트
package migration

import (
	"testing"
)

func TestCheckAlterOrAddOps(t *testing.T) {
	prev := map[string]*CheckConstraint{"same": {Name: "same", Expression: "x > 0"}}
	curr := map[string]*CheckConstraint{
		"same": {Name: "same", Expression: "x > 0"},
		"new":  {Name: "new", Expression: "y > 0"},
	}
	ops := checkAlterOrAddOps("t", []string{"new", "same"}, prev, curr)
	// "new" -> 1 Add; "same" -> 0
	if len(ops) != 1 {
		t.Fatalf("got %d ops, want 1: %#v", len(ops), ops)
	}
	if _, ok := ops[0].(AddCheck); !ok {
		t.Errorf("expected AddCheck, got %T", ops[0])
	}
}
