//ff:func feature=manifest type=test control=sequence
//ff:what ParseTables — `-- @func-managed` 단독/@archived 와 stacking 시 플래그 반영
package ddl

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseTables_StackedArchivedAndFuncManaged(t *testing.T) {
	dir := t.TempDir()
	sql := `-- @archived
-- @func-managed
CREATE TABLE legacy_rpc (
    id BIGSERIAL PRIMARY KEY
);`
	if err := os.WriteFile(filepath.Join(dir, "legacy_rpc.sql"), []byte(sql), 0o644); err != nil {
		t.Fatal(err)
	}
	tables, diags := ParseTables(dir)
	if len(diags) > 0 {
		t.Fatalf("unexpected diags: %v", diags)
	}
	if len(tables) != 1 {
		t.Fatal("no tables")
	}
	if !tables[0].Archived {
		t.Errorf("Archived = false, want true")
	}
	if !tables[0].FuncManaged {
		t.Errorf("FuncManaged = false, want true")
	}
}
