//ff:func feature=validate type=test control=sequence topic=ddl-structural
//ff:what TestDDLHelpers — unit tests for the pure DDL validate helper functions
package ddl

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestToStringSlice(t *testing.T) {
	tests := []struct {
		name string
		in   interface{}
		want []string
	}{
		{"single string", "a", []string{"a"}},
		{"empty string", "", []string{""}},
		{"slice of strings", []interface{}{"a", "b"}, []string{"a", "b"}},
		{"slice with non-strings", []interface{}{"a", 1, "b"}, []string{"a", "b"}},
		{"slice all non-strings", []interface{}{1, 2}, nil},
		{"unrecognised type int", 42, nil},
		{"unrecognised type nil", nil, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := toStringSlice(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toStringSlice(%v) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}

func TestExtractStringsFromSlice(t *testing.T) {
	tests := []struct {
		name string
		in   []interface{}
		want []string
	}{
		{"all strings", []interface{}{"x", "y"}, []string{"x", "y"}},
		{"mixed", []interface{}{"x", 3, true, "y"}, []string{"x", "y"}},
		{"none", []interface{}{1, 2, 3}, nil},
		{"empty", []interface{}{}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := extractStringsFromSlice(tt.in)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSerialReplacement(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"bigserial", "BIGINT GENERATED ALWAYS AS IDENTITY"},
		{"serial", "INTEGER GENERATED ALWAYS AS IDENTITY"},
		{"smallserial", "SMALLINT GENERATED ALWAYS AS IDENTITY"},
		{"unknown", "BIGINT GENERATED ALWAYS AS IDENTITY"},
		{"", "BIGINT GENERATED ALWAYS AS IDENTITY"},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			if got := serialReplacement(tt.in); got != tt.want {
				t.Errorf("serialReplacement(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}

func TestMatchSensitivePattern(t *testing.T) {
	tests := []struct {
		col  string
		want string
	}{
		{"user_password", "password"},
		{"PassWord", "password"},
		{"api_token", "token"},
		{"ssn", "ssn"},
		{"credit_card_no", "credit_card"},
		{"private_key", "private_key"},
		{"username", ""},
		{"email", ""},
		{"", ""},
	}
	for _, tt := range tests {
		t.Run(tt.col, func(t *testing.T) {
			if got := matchSensitivePattern(tt.col); got != tt.want {
				t.Errorf("matchSensitivePattern(%q) = %q, want %q", tt.col, got, tt.want)
			}
		})
	}
}

func TestHasInlineSensitiveAnnotation(t *testing.T) {
	tests := []struct {
		line string
		want bool
	}{
		{"password TEXT -- @sensitive", true},
		{"password TEXT --@sensitive", true},
		{"col TEXT -- @SENSITIVE", true},
		{"col TEXT -- @nosensitive", true},
		{"col TEXT --@nosensitive", true},
		{"password TEXT", false},
		{"col TEXT -- some comment", false},
	}
	for _, tt := range tests {
		t.Run(tt.line, func(t *testing.T) {
			if got := hasInlineSensitiveAnnotation(tt.line); got != tt.want {
				t.Errorf("hasInlineSensitiveAnnotation(%q) = %v, want %v", tt.line, got, tt.want)
			}
		})
	}
}

func TestIsSkippableDDLLine(t *testing.T) {
	tests := []struct {
		line string
		want bool
	}{
		{"", true},
		{"-- comment", true},
		{");", true},
		{"CREATE TABLE foo (", true},
		{"INSERT INTO foo VALUES (1)", true},
		{"ON CONFLICT DO NOTHING", true},
		{"VALUES (1, 2)", true},
		{"PRIMARY KEY (id)", true},
		{"UNIQUE (email)", true},
		{"CHECK (x > 0)", true},
		{"FOREIGN KEY (a) REFERENCES b", true},
		{"CONSTRAINT fk_x FOREIGN KEY", true},
		{"id BIGINT", false},
		{"name TEXT NOT NULL", false},
	}
	for _, tt := range tests {
		t.Run(tt.line, func(t *testing.T) {
			if got := isSkippableDDLLine(tt.line); got != tt.want {
				t.Errorf("isSkippableDDLLine(%q) = %v, want %v", tt.line, got, tt.want)
			}
		})
	}
}

func TestIsSentinelAnnotationLine(t *testing.T) {
	tests := []struct {
		line string
		want bool
	}{
		{"-- @sentinel", true},
		{"--  @sentinel", true},
		{"--@sentinel", true},
		{"-- @other", false},
		{"id BIGINT", false},
		{"", false},
	}
	for _, tt := range tests {
		t.Run(tt.line, func(t *testing.T) {
			if got := isSentinelAnnotationLine(tt.line); got != tt.want {
				t.Errorf("isSentinelAnnotationLine(%q) = %v, want %v", tt.line, got, tt.want)
			}
		})
	}
}

func TestHasSentinelInsert(t *testing.T) {
	tests := []struct {
		name    string
		content string
		table   string
		want    bool
	}{
		{
			"sentinel present",
			"INSERT INTO users (id, name) VALUES (0, 'system');",
			"users", true,
		},
		{
			"insert but no zero sentinel",
			"INSERT INTO users (id, name) VALUES (1, 'alice');",
			"users", false,
		},
		{
			"no insert for table",
			"INSERT INTO posts (id) VALUES (0);",
			"users", false,
		},
		{
			"case insensitive keyword",
			"insert into users values (0, 'x');",
			"users", true,
		},
		{
			"empty content",
			"", "users", false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasSentinelInsert(tt.content, tt.table); got != tt.want {
				t.Errorf("hasSentinelInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScanValidateLineForTerminator(t *testing.T) {
	tests := []struct {
		name         string
		line         string
		inSingle     bool
		wantDone     bool
		wantInSingle bool
	}{
		{"unquoted semicolon", "VALUES (1);", false, true, false},
		{"no semicolon", "VALUES (1)", false, false, false},
		{"semicolon inside quote ignored", "'a;b'", false, false, false},
		{"open quote carries over", "'unterminated", false, false, true},
		{"close quote from carried state", "more'", true, false, false},
		{"escaped doubled quote stays in literal", "'it''s';", false, true, false},
		{"semicolon while inside single is ignored", ";", true, false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			done, inSingle := scanValidateLineForTerminator(tt.line, tt.inSingle)
			if done != tt.wantDone || inSingle != tt.wantInSingle {
				t.Errorf("scanValidateLineForTerminator(%q,%v) = (%v,%v), want (%v,%v)",
					tt.line, tt.inSingle, done, inSingle, tt.wantDone, tt.wantInSingle)
			}
		})
	}
}

func TestFindTableBlockEnd(t *testing.T) {
	lines := []string{
		"CREATE TABLE foo (",
		"  id BIGINT",
		");",
		"CREATE TABLE bar (",
		"  x TEXT",
	}
	if got := findTableBlockEnd(lines, 0); got != 2 {
		t.Errorf("findTableBlockEnd terminated block = %d, want 2", got)
	}
	// No terminator after start index → returns last index.
	if got := findTableBlockEnd(lines, 3); got != len(lines)-1 {
		t.Errorf("findTableBlockEnd unterminated = %d, want %d", got, len(lines)-1)
	}
}

func TestExtractTableBlocks(t *testing.T) {
	f := sqlFile{
		path:    "schema.sql",
		name:    "schema.sql",
		content: "CREATE TABLE users (\n  id BIGINT\n);\nCREATE TABLE IF NOT EXISTS posts (\n  id BIGINT\n);\n",
	}
	blocks := extractTableBlocks(f)
	if len(blocks) != 2 {
		t.Fatalf("got %d blocks, want 2", len(blocks))
	}
	if blocks[0].tableName != "users" {
		t.Errorf("block0 name = %q, want users", blocks[0].tableName)
	}
	if blocks[1].tableName != "posts" {
		t.Errorf("block1 name = %q, want posts", blocks[1].tableName)
	}
	if blocks[0].startLine != 1 {
		t.Errorf("block0 startLine = %d, want 1", blocks[0].startLine)
	}

	// Content with no CREATE TABLE → no blocks.
	if got := extractTableBlocks(sqlFile{content: "SELECT 1;"}); got != nil {
		t.Errorf("expected nil blocks for non-DDL, got %v", got)
	}
}

func TestAllTableContents(t *testing.T) {
	files := []sqlFile{
		{content: "CREATE TABLE users (\n  id BIGINT\n);\n"},
		{content: "CREATE TABLE posts (\n  id BIGINT\n);\n"},
	}
	m := allTableContents(files)
	if _, ok := m["users"]; !ok {
		t.Error("missing users key")
	}
	if _, ok := m["posts"]; !ok {
		t.Error("missing posts key")
	}
	if len(m) != 2 {
		t.Errorf("got %d keys, want 2", len(m))
	}
}

func TestScanInsertsWithAnnotations(t *testing.T) {
	content := "" +
		"CREATE TABLE users (id BIGINT);\n" +
		"-- @sentinel\n" +
		"INSERT INTO users VALUES (0, 'system');\n" +
		"INSERT INTO posts VALUES (1, 'hi');\n"
	got := scanInsertsWithAnnotations(content)
	if len(got) != 2 {
		t.Fatalf("got %d inserts, want 2", len(got))
	}
	if got[0].Table != "users" || !got[0].Annotated {
		t.Errorf("first insert = %+v, want users annotated", got[0])
	}
	if got[1].Table != "posts" || got[1].Annotated {
		t.Errorf("second insert = %+v, want posts not annotated", got[1])
	}
}

func TestD08CheckLine(t *testing.T) {
	f := sqlFile{path: "schema.sql"}
	blk := tableBlock{tableName: "users", startLine: 1}

	// SERIAL column → diagnostic produced.
	d := d08CheckLine(f, blk, "  id SERIAL,", 1)
	if d == nil {
		t.Fatal("expected diagnostic for SERIAL column")
	}
	if !strings.Contains(d.Message, "[D-8]") || !strings.Contains(d.Message, "users.id") {
		t.Errorf("unexpected message: %q", d.Message)
	}
	if !strings.Contains(d.Advice, "IDENTITY") {
		t.Errorf("advice should mention IDENTITY: %q", d.Advice)
	}
	if d.Line != blk.startLine+1 {
		t.Errorf("line = %d, want %d", d.Line, blk.startLine+1)
	}

	// BIGSERIAL also flagged.
	if got := d08CheckLine(f, blk, "n BIGSERIAL", 0); got == nil {
		t.Error("expected diagnostic for BIGSERIAL")
	}

	// Non-serial column → nil.
	if got := d08CheckLine(f, blk, "  id BIGINT,", 0); got != nil {
		t.Errorf("expected nil for BIGINT, got %+v", got)
	}
	// Skippable line → nil.
	if got := d08CheckLine(f, blk, ");", 0); got != nil {
		t.Error("expected nil for skippable line")
	}
	// Line without a column-name match → nil.
	if got := d08CheckLine(f, blk, "123 SERIAL", 0); got != nil {
		t.Error("expected nil when no leading identifier")
	}
	// Single token (no type) → nil.
	if got := d08CheckLine(f, blk, "id", 0); got != nil {
		t.Error("expected nil for single-token line")
	}
}

func TestScanPositionals(t *testing.T) {
	tmp := t.TempDir()
	path := filepath.Join(tmp, "q.sql")
	content := "" +
		"-- name: GetUser :one\n" +
		"SELECT * FROM users\n" +
		"WHERE id = $1\n" +
		"  AND org = $2;\n" +
		"-- name: ListPosts :many\n" +
		"SELECT * FROM posts WHERE owner = $1;\n"
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
	// Query starts at line 1 (the first -- name: header). FindString returns
	// the first $N per line, so $1 (line 3) and $2 (line 4) are collected;
	// scanning stops at the next -- name: header (line 5).
	hits := scanPositionals(path, 1)
	if len(hits) != 2 {
		t.Fatalf("got %d hits, want 2 (stops at next -- name:)", len(hits))
	}
	if hits[0].param != "$1" || hits[1].param != "$2" {
		t.Errorf("params = %v", hits)
	}
	if hits[0].line != 3 {
		t.Errorf("first hit line = %d, want 3", hits[0].line)
	}

	// Nonexistent file → nil.
	if got := scanPositionals(filepath.Join(tmp, "missing.sql"), 1); got != nil {
		t.Errorf("expected nil for missing file, got %v", got)
	}
}

func TestCollectValidateInsertScan(t *testing.T) {
	lines := []string{
		"INSERT INTO users VALUES (",
		"  0, 'sys'",
		");",
		"SELECT 1;",
	}
	r, next := collectValidateInsertScan(lines, 0, "users", true)
	if r.Table != "users" {
		t.Errorf("table = %q, want users", r.Table)
	}
	if !r.Annotated {
		t.Error("expected annotated true")
	}
	if r.StartLine != 1 {
		t.Errorf("startLine = %d, want 1", r.StartLine)
	}
	if next != 3 {
		t.Errorf("next index = %d, want 3", next)
	}
}
