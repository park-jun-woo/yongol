//ff:func feature=migration type=test control=iteration dimension=1
//ff:what render_helpers_unit_test — appendColumnLines/PK/FK/Check/Index/Sentinel + renderTable + CanonicalSQL 단위 테스트
package migration

import (
	"strings"
	"testing"
)

func TestAppendColumnLines(t *testing.T) {
	var b strings.Builder
	appendColumnLines(&b, []*Column{
		{Name: "id", Type: CanonicalType{Base: "INTEGER"}, Nullable: false},
		{Name: "name", Type: CanonicalType{Base: "TEXT"}, Nullable: true},
	})
	got := b.String()
	want := "\n    id INTEGER NOT NULL,\n    name TEXT"
	if got != want {
		t.Errorf("appendColumnLines = %q, want %q", got, want)
	}

	var empty strings.Builder
	appendColumnLines(&empty, nil)
	if empty.String() != "" {
		t.Errorf("empty cols should produce no output, got %q", empty.String())
	}
}

func TestAppendPrimaryKeyClause(t *testing.T) {
	var b strings.Builder
	appendPrimaryKeyClause(&b, []string{"id", "tenant"})
	if got := b.String(); got != ",\n    PRIMARY KEY (id, tenant)" {
		t.Errorf("got %q", got)
	}

	var empty strings.Builder
	appendPrimaryKeyClause(&empty, nil)
	if empty.String() != "" {
		t.Errorf("empty pk should be no-op, got %q", empty.String())
	}
}

func TestAppendForeignKeyClauses(t *testing.T) {
	var b strings.Builder
	appendForeignKeyClauses(&b, []*ForeignKey{
		{Name: "fk_x", Columns: []string{"a"}, RefTable: "other", RefColumns: []string{"id"}, OnDelete: "CASCADE", OnUpdate: "SET NULL"},
		{Name: "fk_y", Columns: []string{"b", "c"}, RefTable: "t2", RefColumns: []string{"d", "e"}},
	})
	got := b.String()
	if !strings.Contains(got, ",\n    CONSTRAINT fk_x FOREIGN KEY (a) REFERENCES other (id) ON DELETE CASCADE ON UPDATE SET NULL") {
		t.Errorf("fk_x clause missing/wrong: %q", got)
	}
	if !strings.Contains(got, ",\n    CONSTRAINT fk_y FOREIGN KEY (b, c) REFERENCES t2 (d, e)") {
		t.Errorf("fk_y clause missing/wrong: %q", got)
	}
	if strings.Contains(got, "fk_y") && strings.Contains(strings.Split(got, "fk_y")[1], "ON DELETE") {
		t.Errorf("fk_y should not emit ON DELETE when empty: %q", got)
	}
}

func TestAppendCheckClauses(t *testing.T) {
	var b strings.Builder
	appendCheckClauses(&b, []*CheckConstraint{{Name: "chk_pos", Expression: "x > 0"}})
	if got := b.String(); got != ",\n    CONSTRAINT chk_pos CHECK (x > 0)" {
		t.Errorf("got %q", got)
	}
}

func TestAppendIndexLines(t *testing.T) {
	tbl := &Table{
		Name: "users",
		Indexes: []*Index{
			{Name: "idx_b", Columns: []string{"b"}},
			{Name: "idx_a", Columns: []string{"a"}, Unique: true},
			{Name: "idx_gin", Columns: []string{"doc"}, Method: "gin"},
			{Name: "idx_w", Columns: []string{"a"}, Where: "a IS NOT NULL"},
		},
	}
	var b strings.Builder
	appendIndexLines(&b, tbl)
	got := b.String()
	lines := strings.Split(strings.TrimRight(got, "\n"), "\n")
	// name-sorted: idx_a, idx_b, idx_gin, idx_w
	want := []string{
		"CREATE UNIQUE INDEX idx_a ON users (a);",
		"CREATE INDEX idx_b ON users (b);",
		"CREATE INDEX idx_gin ON users USING gin (doc);",
		"CREATE INDEX idx_w ON users (a) WHERE a IS NOT NULL;",
	}
	if len(lines) != len(want) {
		t.Fatalf("got %d lines, want %d: %q", len(lines), len(want), got)
	}
	for i := range want {
		if lines[i] != want[i] {
			t.Errorf("line %d = %q, want %q", i, lines[i], want[i])
		}
	}
}

func TestAppendSentinelLines(t *testing.T) {
	var b strings.Builder
	appendSentinelLines(&b, []SentinelInsert{
		{SQL: "INSERT INTO t VALUES (1);\n\n"},
		{SQL: "INSERT INTO t VALUES (2);"},
	})
	want := "\nINSERT INTO t VALUES (1);\n\nINSERT INTO t VALUES (2);\n"
	if got := b.String(); got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestRenderTable(t *testing.T) {
	tbl := &Table{
		Name: "users",
		Columns: []*Column{
			{Name: "id", Type: CanonicalType{Base: "INTEGER"}, Nullable: false},
			{Name: "email", Type: CanonicalType{Base: "TEXT"}, Nullable: true},
		},
		PrimaryKey: []string{"id"},
		Indexes:    []*Index{{Name: "idx_email", Columns: []string{"email"}, Unique: true}},
	}
	got := renderTable(tbl)
	if !strings.HasPrefix(got, "CREATE TABLE users (") {
		t.Errorf("missing CREATE TABLE header: %q", got)
	}
	if !strings.Contains(got, "\n    id INTEGER NOT NULL,") {
		t.Errorf("missing id column: %q", got)
	}
	if !strings.Contains(got, ",\n    PRIMARY KEY (id)") {
		t.Errorf("missing PK clause: %q", got)
	}
	if !strings.Contains(got, "\n);\n") {
		t.Errorf("missing close paren: %q", got)
	}
	if !strings.Contains(got, "CREATE UNIQUE INDEX idx_email ON users (email);") {
		t.Errorf("missing index line: %q", got)
	}
}

func TestCanonicalSQL(t *testing.T) {
	if CanonicalSQL(nil) != "" {
		t.Errorf("nil schema should render empty")
	}
	s := &Schema{Tables: map[string]*Table{
		"zebra": {Name: "zebra", Columns: []*Column{{Name: "id", Type: CanonicalType{Base: "INTEGER"}, Nullable: false}}},
		"apple": {Name: "apple", Columns: []*Column{{Name: "id", Type: CanonicalType{Base: "INTEGER"}, Nullable: false}}},
	}}
	got := CanonicalSQL(s)
	ai := strings.Index(got, "CREATE TABLE apple")
	zi := strings.Index(got, "CREATE TABLE zebra")
	if ai < 0 || zi < 0 {
		t.Fatalf("both tables should render: %q", got)
	}
	if ai > zi {
		t.Errorf("tables should be alphabetically sorted (apple before zebra): %q", got)
	}
}
