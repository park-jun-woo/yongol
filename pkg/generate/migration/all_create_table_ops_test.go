//ff:func feature=migration type=test control=sequence
//ff:what diff_helpers_unit_test — checkDiffForOne/findPrevViaRenameHint/checkDropOps/checkAlterOrAddOps/allCreateTableOps/createTableWithDeps/parseTableFKRef 단위 테스트
package migration

import (
	"testing"
)

func TestAllCreateTableOps(t *testing.T) {
	allCreate := []Operation{
		CreateTable{Table: &Table{Name: "a"}},
		CreateTable{Table: &Table{Name: "b"}},
	}
	if !allCreateTableOps(allCreate) {
		t.Errorf("all-CreateTable slice should return true")
	}
	mixed := []Operation{
		CreateTable{Table: &Table{Name: "a"}},
		DropTable{Name: "b"},
	}
	if allCreateTableOps(mixed) {
		t.Errorf("mixed slice should return false")
	}
	if !allCreateTableOps(nil) {
		t.Errorf("empty slice should vacuously return true")
	}
}
