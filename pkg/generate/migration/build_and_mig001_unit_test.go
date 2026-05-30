//ff:func feature=migration type=test control=iteration dimension=1
//ff:what build_and_mig001_unit_test — buildAlterColumnNullable/Type + mig001CheckRename(Columns/Tables) + removeLegacyBaseline 단위 테스트
package migration

import (
	"os"
	"path/filepath"
	"testing"
)

func TestBuildAlterColumnNullable(t *testing.T) {
	// nil hints -> no backfill
	op := buildAlterColumnNullable("users", "email", true, false, nil)
	if op.Table != "users" || op.Column != "email" || op.From != true || op.To != false || op.Backfill != "" {
		t.Errorf("nil hints op wrong: %+v", op)
	}
	// matching backfill hint applied
	hints := &Hints{Backfills: map[colKey]string{{Table: "users", Column: "email"}: "''"}}
	op = buildAlterColumnNullable("users", "email", true, false, hints)
	if op.Backfill != "''" {
		t.Errorf("backfill not applied: %+v", op)
	}
	// non-matching column -> no backfill
	op = buildAlterColumnNullable("users", "other", true, false, hints)
	if op.Backfill != "" {
		t.Errorf("backfill should not apply to non-matching column: %+v", op)
	}
}

func TestBuildAlterColumnType(t *testing.T) {
	from := CanonicalType{Base: "INTEGER"}
	to := CanonicalType{Base: "BIGINT"}
	op := buildAlterColumnType("t", "id", from, to, nil)
	if op.Using != "" {
		t.Errorf("nil hints should leave Using empty: %+v", op)
	}
	hints := &Hints{Casts: map[colKey]string{{Table: "t", Column: "id"}: "id::bigint"}}
	op = buildAlterColumnType("t", "id", from, to, hints)
	if op.Using != "id::bigint" {
		t.Errorf("cast USING not applied: %+v", op)
	}
}

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

func TestRemoveLegacyBaseline(t *testing.T) {
	dir := t.TempDir()
	// no file -> no-op, no panic
	removeLegacyBaseline(dir)

	legacy := filepath.Join(dir, LegacySnapshotFileName)
	if err := os.WriteFile(legacy, []byte("stale"), 0644); err != nil {
		t.Fatalf("write: %v", err)
	}
	removeLegacyBaseline(dir)
	if _, err := os.Stat(legacy); !os.IsNotExist(err) {
		t.Errorf("legacy baseline should have been removed, stat err = %v", err)
	}
}
