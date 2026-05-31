//ff:func feature=migration type=test control=sequence
//ff:what diff_helpers_unit_test — checkDiffForOne/findPrevViaRenameHint/checkDropOps/checkAlterOrAddOps/allCreateTableOps/createTableWithDeps/parseTableFKRef 단위 테스트
package migration

import (
	"testing"
)

func TestFindPrevViaRenameHint(t *testing.T) {
	old := &Column{Name: "old_name"}
	prevMap := map[string]*Column{"old_name": old}
	rules := []RenameColumnHint{{Table: "users", From: "old_name", To: "new_name"}}

	if got := findPrevViaRenameHint("new_name", prevMap, rules, "users"); got != old {
		t.Errorf("matching rule should return prev column, got %v", got)
	}
	if got := findPrevViaRenameHint("new_name", prevMap, rules, "other_table"); got != nil {
		t.Errorf("table mismatch should return nil")
	}
	if got := findPrevViaRenameHint("unknown", prevMap, rules, "users"); got != nil {
		t.Errorf("no matching To should return nil")
	}
	// rule matches but From column missing in prevMap
	if got := findPrevViaRenameHint("new_name", map[string]*Column{}, rules, "users"); got != nil {
		t.Errorf("missing From in prevMap should return nil")
	}
}
