//ff:func feature=manifest type=test control=iteration dimension=1
//ff:what ParseTables — VARCHAR(N) / NOT NULL / DEFAULT / @archived / @sensitive / @nullable / PK / inline UNIQUE 종합

package ddl

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseTables_VarcharAndConstraints(t *testing.T) {
	dir := t.TempDir()
	sql := `CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    name VARCHAR(100),
    status VARCHAR(32) NOT NULL DEFAULT 'draft',
    password_hash TEXT NOT NULL, -- @sensitive
    old_ref BIGINT, -- @archived
    bio TEXT -- @nullable
);`
	path := filepath.Join(dir, "users.sql")
	if err := os.WriteFile(path, []byte(sql), 0o644); err != nil {
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

	if got := tb.VarcharLen["email"]; got != 255 {
		t.Errorf("VarcharLen[email] = %d, want 255", got)
	}
	if got := tb.VarcharLen["name"]; got != 100 {
		t.Errorf("VarcharLen[name] = %d, want 100", got)
	}
	if !tb.NotNullCols["email"] {
		t.Errorf("NotNullCols[email] = false, want true")
	}
	if tb.NotNullCols["name"] {
		t.Errorf("NotNullCols[name] = true, want false")
	}
	if got := tb.Defaults["status"]; got != "draft" {
		t.Errorf("Defaults[status] = %q, want 'draft'", got)
	}
	if !tb.SensitiveColumns["password_hash"] {
		t.Errorf("SensitiveColumns[password_hash] = false, want true")
	}
	if !tb.ArchivedColumns["old_ref"] {
		t.Errorf("ArchivedColumns[old_ref] = false, want true")
	}
	if !tb.NullableAnnot["bio"] {
		t.Errorf("NullableAnnot[bio] = false, want true")
	}
	if len(tb.PrimaryKey) != 1 || tb.PrimaryKey[0] != "id" {
		t.Errorf("PrimaryKey = %v, want [id]", tb.PrimaryKey)
	}
	// email UNIQUE → inline unique index
	var gotEmailUnique bool
	for _, ix := range tb.Indexes {
		if ix.IsUnique && len(ix.Columns) == 1 && ix.Columns[0] == "email" {
			gotEmailUnique = true
			break
		}
	}
	if !gotEmailUnique {
		t.Errorf("expected UNIQUE index on email, got %v", tb.Indexes)
	}
}
