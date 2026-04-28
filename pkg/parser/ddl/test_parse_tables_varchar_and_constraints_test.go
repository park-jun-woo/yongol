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

	if got := tb.Columns["email"].VarcharLen; got != 255 {
		t.Errorf("Columns[email].VarcharLen = %d, want 255", got)
	}
	if got := tb.Columns["name"].VarcharLen; got != 100 {
		t.Errorf("Columns[name].VarcharLen = %d, want 100", got)
	}
	if !tb.Columns["email"].NotNull {
		t.Errorf("Columns[email].NotNull = false, want true")
	}
	if tb.Columns["name"].NotNull {
		t.Errorf("Columns[name].NotNull = true, want false")
	}
	if got := tb.Columns["status"].DefaultLiteral; got != "draft" {
		t.Errorf("Columns[status].DefaultLiteral = %q, want 'draft'", got)
	}
	if !tb.Columns["status"].HasDefault {
		t.Errorf("Columns[status].HasDefault = false, want true")
	}
	if !tb.Columns["password_hash"].Sensitive {
		t.Errorf("Columns[password_hash].Sensitive = false, want true")
	}
	if !tb.Columns["old_ref"].Archived {
		t.Errorf("Columns[old_ref].Archived = false, want true")
	}
	if !tb.Columns["bio"].NullableAnnot {
		t.Errorf("Columns[bio].NullableAnnot = false, want true")
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
