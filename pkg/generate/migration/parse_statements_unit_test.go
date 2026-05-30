//ff:func feature=migration type=test control=iteration dimension=1
//ff:what parse_statements_unit_test — parseCreateTable/parseCreateIndex/parseAlterTable + 하위 parse 헬퍼 통합 단위 테스트
package migration

import (
	"testing"
)

func TestEnsureTable(t *testing.T) {
	s := NewSchema()
	a := ensureTable(s, "users")
	if a == nil || a.Name != "users" {
		t.Fatalf("ensureTable created wrong table: %+v", a)
	}
	b := ensureTable(s, "users")
	if a != b {
		t.Errorf("ensureTable should return the same instance for an existing name")
	}
	if len(s.Tables) != 1 {
		t.Errorf("schema should have exactly 1 table, got %d", len(s.Tables))
	}
}

func TestParseCreateTable_Columns(t *testing.T) {
	s := NewSchema()
	stmt := `CREATE TABLE users (
		id BIGINT NOT NULL,
		email VARCHAR(255) NOT NULL,
		status TEXT DEFAULT 'active',
		created_at TIMESTAMPTZ,
		PRIMARY KEY (id)
	)`
	if err := parseCreateTable(s, stmt); err != nil {
		t.Fatalf("parseCreateTable error: %v", err)
	}
	tbl := s.Tables["users"]
	if tbl == nil {
		t.Fatalf("users table not registered")
	}
	if len(tbl.Columns) != 4 {
		t.Fatalf("got %d columns, want 4: %+v", len(tbl.Columns), tbl.Columns)
	}
	id := tbl.Columns[0]
	if id.Name != "id" || id.Nullable {
		t.Errorf("id column wrong: %+v", id)
	}
	email := tbl.Columns[1]
	if email.Name != "email" || email.Type.Base != "VARCHAR" || email.Type.Length != 255 {
		t.Errorf("email column wrong: %+v (type %+v)", email, email.Type)
	}
	status := tbl.Columns[2]
	if status.Default != "'active'" {
		t.Errorf("status default = %q, want 'active'", status.Default)
	}
	createdAt := tbl.Columns[3]
	if !createdAt.Nullable {
		t.Errorf("created_at should be nullable")
	}
	if len(tbl.PrimaryKey) != 1 || tbl.PrimaryKey[0] != "id" {
		t.Errorf("PK = %v, want [id]", tbl.PrimaryKey)
	}
}

func TestParseCreateTable_IfNotExists(t *testing.T) {
	s := NewSchema()
	if err := parseCreateTable(s, "CREATE TABLE IF NOT EXISTS t (id INTEGER NOT NULL)"); err != nil {
		t.Fatalf("error: %v", err)
	}
	if s.Tables["t"] == nil {
		t.Errorf("IF NOT EXISTS table not parsed")
	}
}

func TestParseCreateTable_Unparseable(t *testing.T) {
	s := NewSchema()
	if err := parseCreateTable(s, "CREATE INDEX foo ON bar (x)"); err == nil {
		t.Errorf("expected error for non-CREATE-TABLE statement")
	}
}

func TestParseCreateTable_InlineConstraints(t *testing.T) {
	s := NewSchema()
	stmt := `CREATE TABLE orders (
		id INTEGER NOT NULL,
		user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
		qty INTEGER CHECK (qty > 0),
		code TEXT UNIQUE,
		PRIMARY KEY (id),
		CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id),
		CHECK (qty <= 100)
	)`
	if err := parseCreateTable(s, stmt); err != nil {
		t.Fatalf("error: %v", err)
	}
	tbl := s.Tables["orders"]
	// inline REFERENCES + named FK -> 2 FKs
	if len(tbl.ForeignKeys) != 2 {
		t.Errorf("got %d FKs, want 2: %+v", len(tbl.ForeignKeys), tbl.ForeignKeys)
	}
	// find the inline one with ON DELETE CASCADE
	var foundCascade bool
	for _, fk := range tbl.ForeignKeys {
		if fk.RefTable == "users" && fk.OnDelete == "CASCADE" {
			foundCascade = true
		}
	}
	if !foundCascade {
		t.Errorf("inline FK ON DELETE CASCADE not parsed: %+v", tbl.ForeignKeys)
	}
	// inline CHECK + table-level CHECK -> 2 checks
	if len(tbl.Checks) != 2 {
		t.Errorf("got %d checks, want 2: %+v", len(tbl.Checks), tbl.Checks)
	}
	// UNIQUE code -> 1 index
	if len(tbl.Indexes) != 1 || !tbl.Indexes[0].Unique {
		t.Errorf("expected 1 unique index for code: %+v", tbl.Indexes)
	}
}

func TestParseCreateTable_NamedConstraints(t *testing.T) {
	s := NewSchema()
	stmt := `CREATE TABLE t (
		a INTEGER NOT NULL,
		b INTEGER NOT NULL,
		CONSTRAINT pk_t PRIMARY KEY (a, b),
		CONSTRAINT uq_b UNIQUE (b),
		CONSTRAINT chk_a CHECK (a > 0)
	)`
	if err := parseCreateTable(s, stmt); err != nil {
		t.Fatalf("error: %v", err)
	}
	tbl := s.Tables["t"]
	if len(tbl.PrimaryKey) != 2 {
		t.Errorf("PK = %v, want [a b]", tbl.PrimaryKey)
	}
	if len(tbl.Indexes) != 1 || tbl.Indexes[0].Name != "uq_b" {
		t.Errorf("named unique index wrong: %+v", tbl.Indexes)
	}
	if len(tbl.Checks) != 1 || tbl.Checks[0].Name != "chk_a" {
		t.Errorf("named check wrong: %+v", tbl.Checks)
	}
}

func TestParseCreateIndex(t *testing.T) {
	s := NewSchema()
	cases := []struct {
		name    string
		stmt    string
		idxName string
		unique  bool
		method  string
		where   string
		cols    []string
	}{
		{"plain", "CREATE INDEX idx_a ON users (email)", "idx_a", false, "", "", []string{"email"}},
		{"unique", "CREATE UNIQUE INDEX idx_u ON users (email)", "idx_u", true, "", "", []string{"email"}},
		{"using gin", "CREATE INDEX idx_g ON users USING gin (doc)", "idx_g", false, "gin", "", []string{"doc"}},
		{"partial", "CREATE INDEX idx_w ON users (email) WHERE email IS NOT NULL", "idx_w", false, "", "email IS NOT NULL", []string{"email"}},
		{"multi col", "CREATE INDEX idx_m ON users (a, b)", "idx_m", false, "", "", []string{"a", "b"}},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			if err := parseCreateIndex(s, c.stmt); err != nil {
				t.Fatalf("error: %v", err)
			}
			tbl := s.Tables["users"]
			var idx *Index
			for _, ix := range tbl.Indexes {
				if ix.Name == c.idxName {
					idx = ix
				}
			}
			if idx == nil {
				t.Fatalf("index %s not found", c.idxName)
			}
			if idx.Unique != c.unique || idx.Method != c.method || idx.Where != c.where {
				t.Errorf("idx %+v mismatch (want unique=%v method=%q where=%q)", idx, c.unique, c.method, c.where)
			}
			if len(idx.Columns) != len(c.cols) {
				t.Errorf("cols %v, want %v", idx.Columns, c.cols)
			}
		})
	}
}

func TestParseCreateIndex_Unparseable(t *testing.T) {
	s := NewSchema()
	if err := parseCreateIndex(s, "SELECT 1"); err != nil {
		t.Errorf("non-index statement should be skipped silently, got %v", err)
	}
}

func TestParseAlterTable(t *testing.T) {
	s := NewSchema()
	stmt := "ALTER TABLE orders ADD CONSTRAINT fk_u FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE SET NULL ON UPDATE CASCADE"
	if err := parseAlterTable(s, stmt); err != nil {
		t.Fatalf("error: %v", err)
	}
	tbl := s.Tables["orders"]
	if len(tbl.ForeignKeys) != 1 {
		t.Fatalf("expected 1 FK, got %d", len(tbl.ForeignKeys))
	}
	fk := tbl.ForeignKeys[0]
	if fk.Name != "fk_u" || fk.RefTable != "users" || fk.OnDelete != "SET NULL" || fk.OnUpdate != "CASCADE" {
		t.Errorf("FK parsed wrong: %+v", fk)
	}
	if len(fk.Columns) != 1 || fk.Columns[0] != "user_id" {
		t.Errorf("FK columns wrong: %v", fk.Columns)
	}
}

func TestParseAlterTable_NoConstraintName(t *testing.T) {
	s := NewSchema()
	if err := parseAlterTable(s, "ALTER TABLE orders ADD FOREIGN KEY (user_id) REFERENCES users (id)"); err != nil {
		t.Fatalf("error: %v", err)
	}
	fk := s.Tables["orders"].ForeignKeys[0]
	if fk.Name == "" {
		t.Errorf("FK name should be auto-generated when omitted")
	}
}

func TestParseAlterTable_NonMatching(t *testing.T) {
	s := NewSchema()
	if err := parseAlterTable(s, "ALTER TABLE t ADD COLUMN x INTEGER"); err != nil {
		t.Errorf("non-FK ALTER should be skipped silently, got %v", err)
	}
	if len(s.Tables) != 0 {
		t.Errorf("non-matching ALTER should not create tables")
	}
}

func TestApplyRefOnActions(t *testing.T) {
	fk := &ForeignKey{}
	toks := []string{"users(id)", "ON", "DELETE", "SET", "NULL", "ON", "UPDATE", "CASCADE"}
	end := applyRefOnActions(fk, toks, 1)
	if fk.OnDelete != "SET NULL" || fk.OnUpdate != "CASCADE" {
		t.Errorf("OnDelete=%q OnUpdate=%q, want SET NULL / CASCADE", fk.OnDelete, fk.OnUpdate)
	}
	if end != len(toks) {
		t.Errorf("consumed index = %d, want %d", end, len(toks))
	}
}

func TestCollectOnAction(t *testing.T) {
	cases := []struct {
		toks []string
		want string
		n    int
	}{
		{[]string{"CASCADE"}, "CASCADE", 1},
		{[]string{"set", "null"}, "SET NULL", 2},
		{[]string{"no", "action"}, "NO ACTION", 2},
		{[]string{"restrict", "X"}, "RESTRICT", 1},
		{nil, "", 0},
	}
	for _, c := range cases {
		got, n := collectOnAction(c.toks)
		if got != c.want || n != c.n {
			t.Errorf("collectOnAction(%v) = (%q,%d), want (%q,%d)", c.toks, got, n, c.want, c.n)
		}
	}
}

func TestApplyAlterFKOnActions(t *testing.T) {
	fk := &ForeignKey{}
	applyAlterFKOnActions(fk, " ON DELETE CASCADE ON UPDATE RESTRICT")
	if fk.OnDelete != "CASCADE" || fk.OnUpdate != "RESTRICT" {
		t.Errorf("got OnDelete=%q OnUpdate=%q", fk.OnDelete, fk.OnUpdate)
	}
	fk2 := &ForeignKey{}
	applyAlterFKOnActions(fk2, "")
	if fk2.OnDelete != "" || fk2.OnUpdate != "" {
		t.Errorf("empty tail should leave actions empty: %+v", fk2)
	}
}
