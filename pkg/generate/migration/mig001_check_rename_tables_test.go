//ff:func feature=migration type=test control=sequence
//ff:what build_and_mig001_unit_test — buildAlterColumnNullable/Type + mig001CheckRename(Columns/Tables) + removeLegacyBaseline 단위 테스트
package migration

import (
	"testing"
)

func TestMig001CheckRenameTables(t *testing.T) {
	prev := NewSchema()
	ensureTable(prev, "old_t")
	curr := NewSchema()
	ensureTable(curr, "new_t")

	ok := mig001CheckRenameTables(prev, curr, []RenameTableHint{{From: "old_t", To: "new_t"}})
	if len(ok) != 0 {
		t.Errorf("valid table rename should produce no diags, got %v", ok)
	}

	bad := mig001CheckRenameTables(prev, curr, []RenameTableHint{{From: "ghost", To: "phantom"}})
	if len(bad) != 2 {
		t.Errorf("expected 2 MIG-001 diags, got %d: %v", len(bad), bad)
	}
}
