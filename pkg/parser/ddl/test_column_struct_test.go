//ff:func feature=manifest type=test control=iteration dimension=1
//ff:what Column — RawType 보존 + 통합 메타 회귀 (Phase002)

package ddl

import (
	"os"
	"path/filepath"
	"testing"
)

func TestColumnStruct_RawTypePreservation(t *testing.T) {
	dir := t.TempDir()
	sql := `CREATE TABLE preserve (
    id BIGINT NOT NULL,
    email VARCHAR(255) NOT NULL,
    tags TEXT[],
    amount NUMERIC(10,2),
    payload JSONB NOT NULL DEFAULT '{}',
    org_id UUID NOT NULL,
    role VARCHAR(20) NOT NULL DEFAULT 'member' CHECK (role IN ('member','admin')),
    notes TEXT, -- @nullable
    archived_col BIGINT, -- @archived
    password VARCHAR(60) NOT NULL -- @sensitive
);`
	if err := os.WriteFile(filepath.Join(dir, "preserve.sql"), []byte(sql), 0o644); err != nil {
		t.Fatal(err)
	}
	tables, diags := ParseTables(dir)
	if len(diags) > 0 {
		t.Fatalf("unexpected diags: %v", diags)
	}
	if len(tables) != 1 {
		t.Fatalf("tables count = %d, want 1", len(tables))
	}
	tb := tables[0]

	cases := []struct {
		col           string
		rawType       string
		notNull       bool
		nullableAnnot bool
		hasDefault    bool
		defaultLit    string
		varcharLen    int
		archived      bool
		sensitive     bool
		checkLen      int
	}{
		{"id", "BIGINT", true, false, false, "", 0, false, false, 0},
		{"email", "VARCHAR(255)", true, false, false, "", 255, false, false, 0},
		{"tags", "TEXT[]", false, false, false, "", 0, false, false, 0},
		{"amount", "NUMERIC(10,2)", false, false, false, "", 0, false, false, 0},
		{"payload", "JSONB", true, false, true, "{}", 0, false, false, 0},
		{"org_id", "UUID", true, false, false, "", 0, false, false, 0},
		{"role", "VARCHAR(20)", true, false, true, "member", 20, false, false, 2},
		{"notes", "TEXT", false, true, false, "", 0, false, false, 0},
		{"archived_col", "BIGINT", false, false, false, "", 0, true, false, 0},
		{"password", "VARCHAR(60)", true, false, false, "", 60, false, true, 0},
	}
	for _, c := range cases {
		got, ok := tb.Columns[c.col]
		if !ok {
			t.Errorf("Columns[%s] missing", c.col)
			continue
		}
		if got.RawType != c.rawType {
			t.Errorf("Columns[%s].RawType = %q, want %q", c.col, got.RawType, c.rawType)
		}
		if got.NotNull != c.notNull {
			t.Errorf("Columns[%s].NotNull = %v, want %v", c.col, got.NotNull, c.notNull)
		}
		if got.NullableAnnot != c.nullableAnnot {
			t.Errorf("Columns[%s].NullableAnnot = %v, want %v", c.col, got.NullableAnnot, c.nullableAnnot)
		}
		if got.HasDefault != c.hasDefault {
			t.Errorf("Columns[%s].HasDefault = %v, want %v", c.col, got.HasDefault, c.hasDefault)
		}
		if got.DefaultLiteral != c.defaultLit {
			t.Errorf("Columns[%s].DefaultLiteral = %q, want %q", c.col, got.DefaultLiteral, c.defaultLit)
		}
		if got.VarcharLen != c.varcharLen {
			t.Errorf("Columns[%s].VarcharLen = %d, want %d", c.col, got.VarcharLen, c.varcharLen)
		}
		if got.Archived != c.archived {
			t.Errorf("Columns[%s].Archived = %v, want %v", c.col, got.Archived, c.archived)
		}
		if got.Sensitive != c.sensitive {
			t.Errorf("Columns[%s].Sensitive = %v, want %v", c.col, got.Sensitive, c.sensitive)
		}
		if len(got.CheckEnum) != c.checkLen {
			t.Errorf("Columns[%s].CheckEnum len = %d, want %d", c.col, len(got.CheckEnum), c.checkLen)
		}
	}

	// ColumnOrder must match definition order.
	wantOrder := []string{"id", "email", "tags", "amount", "payload", "org_id", "role", "notes", "archived_col", "password"}
	if len(tb.ColumnOrder) != len(wantOrder) {
		t.Fatalf("ColumnOrder len = %d, want %d", len(tb.ColumnOrder), len(wantOrder))
	}
	for i, c := range wantOrder {
		if tb.ColumnOrder[i] != c {
			t.Errorf("ColumnOrder[%d] = %q, want %q", i, tb.ColumnOrder[i], c)
		}
	}

	// Missing column lookup returns ok=false.
	if _, ok := tb.Columns["does_not_exist"]; ok {
		t.Errorf("Columns[does_not_exist] should not be present")
	}
}
