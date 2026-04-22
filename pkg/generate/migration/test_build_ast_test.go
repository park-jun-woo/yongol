//ff:func feature=migration type=test control=iteration dimension=1
//ff:what BuildAST — zenflow DDL + 합성 DDL 로 AST 스키마 빌드 검증
package migration

import (
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestBuildASTFromSQL_Basic(t *testing.T) {
	sql := `
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL DEFAULT '',
    active BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
`
	s := NewSchema()
	if err := BuildASTFromSQL(s, sql); err != nil {
		t.Fatalf("parse: %v", err)
	}
	tbl, ok := s.Tables["users"]
	if !ok {
		t.Fatalf("users table not found, got %v", s.Tables)
	}
	if len(tbl.Columns) != 5 {
		t.Fatalf("expected 5 columns, got %d", len(tbl.Columns))
	}
	if tbl.Columns[0].Name != "id" || tbl.Columns[0].Type.Base != "BIGINT" {
		t.Errorf("id column: %+v", tbl.Columns[0])
	}
	if len(tbl.PrimaryKey) != 1 || tbl.PrimaryKey[0] != "id" {
		t.Errorf("PK: %+v", tbl.PrimaryKey)
	}
	email := tbl.Columns[1]
	if email.Type.Base != "VARCHAR" || email.Type.Length != 255 {
		t.Errorf("email type: %+v", email.Type)
	}
	if email.Nullable {
		t.Errorf("email should be NOT NULL")
	}
	// UNIQUE inline produced an index
	foundUnique := false
	for _, idx := range tbl.Indexes {
		if idx.Unique && len(idx.Columns) == 1 && idx.Columns[0] == "email" {
			foundUnique = true
		}
	}
	if !foundUnique {
		t.Errorf("unique index on email not generated: %+v", tbl.Indexes)
	}
	// Default normalisation — CURRENT_TIMESTAMP should already be canonical.
	if tbl.Columns[4].Default != "CURRENT_TIMESTAMP" {
		t.Errorf("created_at default: %q", tbl.Columns[4].Default)
	}
	// TRUE normalization
	if tbl.Columns[3].Default != "TRUE" {
		t.Errorf("active default: %q", tbl.Columns[3].Default)
	}
}

func TestBuildASTFromSQL_ForeignKey(t *testing.T) {
	sql := `
CREATE TABLE organizations (id BIGSERIAL PRIMARY KEY);
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    org_id BIGINT NOT NULL REFERENCES organizations(id) ON DELETE CASCADE
);
`
	s := NewSchema()
	if err := BuildASTFromSQL(s, sql); err != nil {
		t.Fatal(err)
	}
	u := s.Tables["users"]
	if u == nil || len(u.ForeignKeys) != 1 {
		t.Fatalf("expected 1 FK on users, got %+v", u)
	}
	fk := u.ForeignKeys[0]
	if fk.RefTable != "organizations" || fk.RefColumns[0] != "id" ||
		fk.OnDelete != "CASCADE" {
		t.Errorf("FK: %+v", fk)
	}
	if fk.Name != "users_org_id_fkey" {
		t.Errorf("FK name: %q", fk.Name)
	}
}

func TestBuildASTFromSQL_CreateIndex(t *testing.T) {
	sql := `
CREATE TABLE t (id BIGSERIAL PRIMARY KEY, name TEXT NOT NULL);
CREATE INDEX idx_t_name ON t (name);
CREATE UNIQUE INDEX uq_t_name ON t (name) WHERE name <> '';
`
	s := NewSchema()
	if err := BuildASTFromSQL(s, sql); err != nil {
		t.Fatal(err)
	}
	tbl := s.Tables["t"]
	if len(tbl.Indexes) != 2 {
		t.Fatalf("expected 2 indexes, got %d: %+v", len(tbl.Indexes), tbl.Indexes)
	}
	var plain, uq *Index
	for _, idx := range tbl.Indexes {
		if idx.Name == "idx_t_name" {
			plain = idx
		}
		if idx.Name == "uq_t_name" {
			uq = idx
		}
	}
	if plain == nil || plain.Unique {
		t.Errorf("plain idx: %+v", plain)
	}
	if uq == nil || !uq.Unique || !strings.Contains(uq.Where, "<>") {
		t.Errorf("unique idx: %+v", uq)
	}
}

func TestBuildASTFromDir_Zenflow(t *testing.T) {
	_, thisFile, _, _ := runtime.Caller(0)
	// pkg/generate/migration/test_build_ast_test.go → repo root four levels up.
	root := filepath.Join(filepath.Dir(thisFile), "..", "..", "..")
	ddlDir := filepath.Join(root, "dummys", "zenflow", "try-02", "specs", "db")
	s, err := BuildASTFromDir(ddlDir, []string{SnapshotFileName})
	if err != nil {
		t.Fatalf("BuildASTFromDir: %v", err)
	}
	// Expect at least the core zenflow tables
	expect := []string{"users", "organizations", "workflows", "actions"}
	for _, name := range expect {
		if _, ok := s.Tables[name]; !ok {
			t.Errorf("table %q missing from AST (have %v)", name, mapKeys(s.Tables))
		}
	}
}

func mapKeys[V any](m map[string]V) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	return out
}
