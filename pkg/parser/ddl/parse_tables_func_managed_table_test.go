//ff:func feature=manifest type=test control=sequence
//ff:what ParseTables — `-- @func-managed` 단독/@archived 와 stacking 시 플래그 반영
package ddl

import (
	"os"
	"path/filepath"
	"testing"
)

// TestParseTables_FuncManagedTable verifies a standalone `-- @func-managed`
// comment above CREATE TABLE sets FuncManaged (and leaves Archived false).
func TestParseTables_FuncManagedTable(t *testing.T) {
	dir := t.TempDir()
	sql := `-- @func-managed
CREATE TABLE bids (
    id BIGSERIAL PRIMARY KEY
);`
	if err := os.WriteFile(filepath.Join(dir, "bids.sql"), []byte(sql), 0o644); err != nil {
		t.Fatal(err)
	}
	tables, diags := ParseTables(dir)
	if len(diags) > 0 {
		t.Fatalf("unexpected diags: %v", diags)
	}
	if len(tables) != 1 {
		t.Fatal("no tables")
	}
	if !tables[0].FuncManaged {
		t.Errorf("FuncManaged = false, want true")
	}
	if tables[0].Archived {
		t.Errorf("Archived = true, want false (func-managed must not imply archived)")
	}
}

// TestParseTables_StackedArchivedAndFuncManaged verifies that stacking
// `-- @archived` and `-- @func-managed` on consecutive lines above a
// CREATE TABLE sets BOTH flags independently.
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
