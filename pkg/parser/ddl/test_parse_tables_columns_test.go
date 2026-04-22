//ff:func feature=manifest type=test control=sequence
//ff:what ParseTables — VARCHAR(N) / NOT NULL / DEFAULT / @archived / @sensitive 수집 검증

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

func TestParseTables_CheckEnum(t *testing.T) {
	dir := t.TempDir()
	sql := `CREATE TABLE orders (
    id BIGSERIAL PRIMARY KEY,
    status VARCHAR(32) NOT NULL CHECK (status IN ('pending','paid','cancelled'))
);`
	if err := os.WriteFile(filepath.Join(dir, "orders.sql"), []byte(sql), 0o644); err != nil {
		t.Fatal(err)
	}
	tables, diags := ParseTables(dir)
	if len(diags) > 0 {
		t.Fatalf("unexpected diags: %v", diags)
	}
	if len(tables) != 1 {
		t.Fatalf("tables count = %d", len(tables))
	}
	vals := tables[0].CheckEnums["status"]
	if len(vals) != 3 {
		t.Fatalf("CheckEnums[status] = %v, want 3 values", vals)
	}
	want := map[string]bool{"pending": true, "paid": true, "cancelled": true}
	for _, v := range vals {
		if !want[v] {
			t.Errorf("unexpected enum val %q", v)
		}
	}
}

func TestParseTables_ForeignKeys(t *testing.T) {
	dir := t.TempDir()
	sql := `CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY
);

CREATE TABLE posts (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id),
    editor_id BIGINT,
    CONSTRAINT fk_editor FOREIGN KEY (editor_id) REFERENCES users(id)
);`
	if err := os.WriteFile(filepath.Join(dir, "schema.sql"), []byte(sql), 0o644); err != nil {
		t.Fatal(err)
	}
	tables, diags := ParseTables(dir)
	if len(diags) > 0 {
		t.Fatalf("unexpected diags: %v", diags)
	}
	byName := map[string]Table{}
	for _, tb := range tables {
		byName[tb.Name] = tb
	}
	posts, ok := byName["posts"]
	if !ok {
		t.Fatal("posts missing")
	}
	if len(posts.ForeignKeys) != 2 {
		t.Fatalf("ForeignKeys count = %d, want 2: %v", len(posts.ForeignKeys), posts.ForeignKeys)
	}
	haveInline, haveConstraint := false, false
	for _, fk := range posts.ForeignKeys {
		if fk.RefTable != "users" || fk.RefColumn != "id" {
			t.Errorf("unexpected fk: %+v", fk)
		}
		if fk.Column == "user_id" {
			haveInline = true
		}
		if fk.Column == "editor_id" {
			haveConstraint = true
		}
	}
	if !haveInline {
		t.Errorf("inline FK on user_id missing")
	}
	if !haveConstraint {
		t.Errorf("named CONSTRAINT FK missing")
	}
}

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

func TestParseTables_ColumnOrderAndGoType(t *testing.T) {
	dir := t.TempDir()
	sql := `CREATE TABLE mixed (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    amount NUMERIC NOT NULL,
    active BOOLEAN NOT NULL DEFAULT 'false',
    payload JSONB,
    created_at TIMESTAMPTZ NOT NULL
);`
	if err := os.WriteFile(filepath.Join(dir, "mixed.sql"), []byte(sql), 0o644); err != nil {
		t.Fatal(err)
	}
	tables, diags := ParseTables(dir)
	if len(diags) > 0 {
		t.Fatalf("unexpected diags: %v", diags)
	}
	tb := tables[0]
	wantOrder := []string{"id", "name", "amount", "active", "payload", "created_at"}
	if len(tb.ColumnOrder) != len(wantOrder) {
		t.Fatalf("ColumnOrder len = %d, want %d: %v", len(tb.ColumnOrder), len(wantOrder), tb.ColumnOrder)
	}
	for i, c := range wantOrder {
		if tb.ColumnOrder[i] != c {
			t.Errorf("ColumnOrder[%d] = %q, want %q", i, tb.ColumnOrder[i], c)
		}
	}
	wantTypes := map[string]string{
		"id":         "int64",
		"name":       "string",
		"amount":     "float64",
		"active":     "bool",
		"payload":    "json.RawMessage",
		"created_at": "time.Time",
	}
	for col, wt := range wantTypes {
		if got := tb.Columns[col]; got != wt {
			t.Errorf("Columns[%s] = %q, want %q", col, got, wt)
		}
	}
}
