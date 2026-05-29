//ff:func feature=manifest type=test control=iteration dimension=1
//ff:what ParseTables — CREATE UNIQUE INDEX / CREATE INDEX 수집

package ddl

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseTables_CreateIndex(t *testing.T) {
	dir := t.TempDir()
	sql := `CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    email VARCHAR(255)
);
CREATE UNIQUE INDEX idx_users_email ON users (email);
CREATE INDEX idx_users_id ON users (id);
`
	if err := os.WriteFile(filepath.Join(dir, "schema.sql"), []byte(sql), 0o644); err != nil {
		t.Fatal(err)
	}
	tables, diags := ParseTables(dir)
	if len(diags) > 0 {
		t.Fatalf("unexpected diags: %v", diags)
	}
	if len(tables) != 1 {
		t.Fatalf("tables count = %d", len(tables))
	}
	tb := tables[0]
	// primary key inline + 2 CREATE INDEX (one unique, one regular)
	// plus may include any from PK. Filter the explicitly named ones.
	var uniq, reg bool
	for _, ix := range tb.Indexes {
		if ix.Name == "idx_users_email" && ix.IsUnique && len(ix.Columns) == 1 && ix.Columns[0] == "email" {
			uniq = true
		}
		if ix.Name == "idx_users_id" && !ix.IsUnique && len(ix.Columns) == 1 && ix.Columns[0] == "id" {
			reg = true
		}
	}
	if !uniq {
		t.Errorf("expected unique index idx_users_email: %v", tb.Indexes)
	}
	if !reg {
		t.Errorf("expected regular index idx_users_id: %v", tb.Indexes)
	}
}
