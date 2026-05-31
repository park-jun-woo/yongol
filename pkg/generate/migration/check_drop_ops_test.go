//ff:func feature=migration type=test control=sequence
//ff:what diff_helpers_unit_test — checkDiffForOne/findPrevViaRenameHint/checkDropOps/checkAlterOrAddOps/allCreateTableOps/createTableWithDeps/parseTableFKRef 단위 테스트
package migration

import (
	"testing"
)

func TestCheckDropOps(t *testing.T) {
	prevNames := []string{"keep", "drop_me"}
	curr := map[string]*CheckConstraint{"keep": {Name: "keep"}}
	ops := checkDropOps("t", prevNames, curr)
	if len(ops) != 1 {
		t.Fatalf("got %d ops, want 1", len(ops))
	}
	dc, ok := ops[0].(DropCheck)
	if !ok || dc.Name != "drop_me" || dc.Table != "t" {
		t.Errorf("expected DropCheck drop_me on t, got %#v", ops[0])
	}
}
