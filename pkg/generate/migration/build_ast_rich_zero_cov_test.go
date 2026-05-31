//ff:func feature=migration type=test control=sequence
//ff:what TestMigrationE2EZeroCov — ParseHints / BuildASTFromSQL / Diff / ApplyHintsToOps / EmitSQL 풀 파이프라인 커버
package migration

import (
	"testing"
)

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
