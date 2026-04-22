//ff:func feature=manifest type=test control=sequence
//ff:what DDL helper (parseCheckEnum / parseRef / extractVarcharLen / pgTypeToGo / parseInlineFK / parseConstraintFK / extractDefaultString / extractParenColumns / isArchivedAnnotation / extractTableName / findTableKeyword) happy/edge 회귀

package ddl

import "testing"

func TestParseCheckEnum_Basic(t *testing.T) {
	col, vals := parseCheckEnum(`status VARCHAR(32) NOT NULL CHECK (status IN ('a','b','c'))`)
	if col != "status" {
		t.Errorf("col = %q, want status", col)
	}
	if len(vals) != 3 {
		t.Fatalf("vals = %v", vals)
	}
}

func TestParseCheckEnum_NoMatch(t *testing.T) {
	col, vals := parseCheckEnum("no check clause here")
	if col != "" || vals != nil {
		t.Errorf("got col=%q vals=%v, want empty", col, vals)
	}
}

func TestParseRef_WithParen(t *testing.T) {
	table, col := parseRef("users(id)")
	if table != "users" || col != "id" {
		t.Errorf("got (%q,%q), want (users,id)", table, col)
	}
}

func TestParseRef_NoParen(t *testing.T) {
	table, col := parseRef("users")
	if table != "users" || col != "" {
		t.Errorf("got (%q,%q), want (users,\"\")", table, col)
	}
}

func TestExtractVarcharLen_Match(t *testing.T) {
	if got := extractVarcharLen("VARCHAR(64)"); got != 64 {
		t.Errorf("got %d, want 64", got)
	}
}

func TestExtractVarcharLen_NoMatch(t *testing.T) {
	if got := extractVarcharLen("TEXT"); got != 0 {
		t.Errorf("got %d, want 0", got)
	}
}

func TestPgTypeToGo_Variants(t *testing.T) {
	cases := map[string]string{
		"BIGINT":      "int64",
		"SERIAL":      "int64",
		"TEXT":        "string",
		"BOOLEAN":     "bool",
		"TIMESTAMPTZ": "time.Time",
		"JSONB":       "json.RawMessage",
		"NUMERIC":     "float64",
		"VARCHAR(64)": "string",
		"UNKNOWN_XY":  "string",
	}
	for in, want := range cases {
		if got := pgTypeToGo(in); got != want {
			t.Errorf("pgTypeToGo(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestParseInlineFK_Match(t *testing.T) {
	fk, ok := parseInlineFK("user_id", []string{"user_id", "BIGINT", "NOT", "NULL", "REFERENCES", "users(id),"})
	if !ok {
		t.Fatalf("expected ok")
	}
	if fk.Column != "user_id" || fk.RefTable != "users" || fk.RefColumn != "id" {
		t.Errorf("fk = %+v", fk)
	}
}

func TestParseInlineFK_NoRef(t *testing.T) {
	_, ok := parseInlineFK("user_id", []string{"user_id", "BIGINT", "NOT", "NULL"})
	if ok {
		t.Errorf("expected no inline FK")
	}
}

func TestParseConstraintFK_Match(t *testing.T) {
	fk, ok := parseConstraintFK("CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id)")
	if !ok {
		t.Fatalf("expected ok")
	}
	if fk.Column != "user_id" || fk.RefTable != "users" || fk.RefColumn != "id" {
		t.Errorf("fk = %+v", fk)
	}
}

func TestParseConstraintFK_Invalid(t *testing.T) {
	if _, ok := parseConstraintFK("not a fk line"); ok {
		t.Errorf("expected false")
	}
}

func TestExtractDefaultString_Match(t *testing.T) {
	if got := extractDefaultString("status VARCHAR(32) NOT NULL DEFAULT 'draft'"); got != "draft" {
		t.Errorf("got %q, want draft", got)
	}
}

func TestExtractDefaultString_NoString(t *testing.T) {
	if got := extractDefaultString("count INTEGER NOT NULL DEFAULT 0"); got != "" {
		t.Errorf("got %q, want empty (numeric default)", got)
	}
}

func TestExtractParenColumns_Multi(t *testing.T) {
	got := extractParenColumns("UNIQUE (a, b , c)")
	if len(got) != 3 || got[0] != "a" || got[1] != "b" || got[2] != "c" {
		t.Errorf("got %v", got)
	}
}

func TestExtractParenColumns_NoParen(t *testing.T) {
	if got := extractParenColumns("PRIMARY KEY id"); got != nil {
		t.Errorf("got %v, want nil", got)
	}
}

func TestIsArchivedAnnotation(t *testing.T) {
	if !isArchivedAnnotation("-- @archived") {
		t.Errorf("missed -- @archived")
	}
	if isArchivedAnnotation("-- some note") {
		t.Errorf("false positive")
	}
	if isArchivedAnnotation("CREATE TABLE x (") {
		t.Errorf("false positive for CREATE TABLE")
	}
}

func TestFindTableKeyword(t *testing.T) {
	if got := findTableKeyword([]string{"CREATE", "TABLE", "users", "("}); got != 1 {
		t.Errorf("got %d, want 1", got)
	}
	if got := findTableKeyword([]string{"CREATE", "INDEX", "idx"}); got != -1 {
		t.Errorf("got %d, want -1", got)
	}
}

func TestExtractTableName_WritesTable(t *testing.T) {
	tables := map[string]*Table{}
	name := extractTableName("CREATE TABLE users (", tables, "x.sql", 5)
	if name != "users" {
		t.Errorf("name = %q", name)
	}
	tb, ok := tables["users"]
	if !ok {
		t.Fatalf("users not registered")
	}
	if tb.File != "x.sql" || tb.Line != 5 {
		t.Errorf("File/Line = %q/%d", tb.File, tb.Line)
	}
}
