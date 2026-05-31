//ff:func feature=migration type=test control=sequence
//ff:what TestMigrationE2EZeroCov — ParseHints / BuildASTFromSQL / Diff / ApplyHintsToOps / EmitSQL 풀 파이프라인 커버

package migration

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestParseHints_AllTags_ZeroCov(t *testing.T) {
	comments := []ddl.HintComment{
		{Tag: "cast", TableCtx: "users", ColumnCtx: "age", Args: map[string]string{"using": "age::int"}},
		{Tag: "backfill", TableCtx: "users", ColumnCtx: "age", Args: map[string]string{"default": "0"}},
		{Tag: "rename", TableCtx: "users", ColumnCtx: "fullname", Args: map[string]string{"from": "name"}},
		{Tag: "rename", BlockAbove: true, Args: map[string]string{"from": "old_users", "to": "users"}},
		{Tag: "data_migration", TableCtx: "users", Args: map[string]string{"file": "users.sql"}},
		{Tag: "allow_destructive", TableCtx: "users"},
	}
	h := ParseHints(comments)
	if len(h.RenameColumns) == 0 {
		t.Errorf("expected a column rename hint")
	}
	if len(h.RenameTables) == 0 {
		t.Errorf("expected a table rename hint")
	}
	if !h.AllowDestructive["users"] {
		t.Errorf("expected allow_destructive for users")
	}
	if len(h.DataMigrations) == 0 {
		t.Errorf("expected a data migration hint")
	}
}

func TestBuildASTRich_ZeroCov(t *testing.T) {
	sql := `
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    status VARCHAR(20) NOT NULL DEFAULT 'active' CHECK (status IN ('active','banned')),
    org_id BIGINT NOT NULL REFERENCES orgs(id)
);
CREATE INDEX idx_users_email ON users (email);
CREATE TABLE orgs (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL
);`
	s := NewSchema()
	if err := BuildASTFromSQL(s, sql); err != nil {
		t.Fatalf("BuildASTFromSQL: %v", err)
	}
	if _, ok := s.Tables["users"]; !ok {
		t.Fatalf("users table not parsed")
	}
	if len(s.Tables["users"].Columns) != 4 {
		t.Errorf("users columns = %d, want 4", len(s.Tables["users"].Columns))
	}
}

func TestDiffAndEmit_FullPipeline_ZeroCov(t *testing.T) {
	prev := mustAST(t, `
CREATE TABLE orgs (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL
);
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    age INTEGER,
    legacy_col TEXT
);
CREATE INDEX idx_users_email ON users (email);`)

	curr := mustAST(t, `
CREATE TABLE orgs (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL
);
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    email VARCHAR(320) NOT NULL,
    age BIGINT NOT NULL,
    org_id BIGINT NOT NULL REFERENCES orgs(id),
    CHECK (age >= 0)
);
CREATE INDEX idx_users_email ON users (email) WHERE email IS NOT NULL;`)

	hints := ParseHints([]ddl.HintComment{
		{Tag: "cast", TableCtx: "users", ColumnCtx: "age", Args: map[string]string{"using": "age::bigint"}},
		{Tag: "backfill", TableCtx: "users", ColumnCtx: "age", Args: map[string]string{"value": "0"}},
		{Tag: "allow_destructive", TableCtx: "users"},
	})

	ops := Diff(prev, curr, hints)
	if len(ops) == 0 {
		t.Fatalf("expected diff operations, got none")
	}

	withHints := ApplyHintsToOps(ops, hints)
	sql := EmitSQL(withHints, EmitOptions{YongolVersion: "v0.0.0"})
	if !strings.Contains(sql, "ALTER TABLE users") {
		t.Errorf("emitted SQL missing ALTER TABLE users:\n%s", sql)
	}

	desc := InferDescription(ops)
	if desc == "" {
		t.Errorf("InferDescription returned empty")
	}
}

func TestDiffCreateAndDrop_ZeroCov(t *testing.T) {
	prev := mustAST(t, `CREATE TABLE gone (id BIGINT PRIMARY KEY);`)
	curr := mustAST(t, `CREATE TABLE fresh (id BIGINT PRIMARY KEY, label TEXT NOT NULL);`)
	hints := ParseHints([]ddl.HintComment{{Tag: "allow_destructive", TableCtx: "gone"}})
	ops := Diff(prev, curr, hints)
	withHints := ApplyHintsToOps(ops, hints)
	sql := EmitSQL(withHints, EmitOptions{})
	if !strings.Contains(sql, "CREATE TABLE fresh") {
		t.Errorf("missing CREATE TABLE fresh:\n%s", sql)
	}
}
