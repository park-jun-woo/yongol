//ff:func feature=manifest type=test control=sequence
//ff:what ParseTables — `-- @archived` 주석이 CREATE TABLE 위에 오면 Archived=true

package ddl

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseTables_ArchivedTable(t *testing.T) {
	dir := t.TempDir()
	sql := `-- @archived
CREATE TABLE legacy (
    id BIGSERIAL PRIMARY KEY
);`
	if err := os.WriteFile(filepath.Join(dir, "legacy.sql"), []byte(sql), 0o644); err != nil {
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
}
