//ff:func feature=migration type=test control=sequence
//ff:what build_and_mig001_unit_test — buildAlterColumnNullable/Type + mig001CheckRename(Columns/Tables) + removeLegacyBaseline 단위 테스트
package migration

import (
	"testing"
)

func TestMig001CheckRenameColumns(t *testing.T) {
	prev := NewSchema()
	pt := ensureTable(prev, "users")
	pt.Columns = []*Column{{Name: "old"}}
	curr := NewSchema()
	ct := ensureTable(curr, "users")
	ct.Columns = []*Column{{Name: "new"}}

	// valid rename: from=old (in prev), to=new (in curr) -> no diags
	ok := mig001CheckRenameColumns(prev, curr, []RenameColumnHint{{Table: "users", From: "old", To: "new"}})
	if len(ok) != 0 {
		t.Errorf("valid rename should produce no diags, got %v", ok)
	}

	// from missing in prev + to missing in curr -> 2 diags
	bad := mig001CheckRenameColumns(prev, curr, []RenameColumnHint{{Table: "users", From: "ghost", To: "phantom"}})
	if len(bad) != 2 {
		t.Errorf("expected 2 MIG-001 diags, got %d: %v", len(bad), bad)
	}
}
