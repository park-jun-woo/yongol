//ff:func feature=migration type=test control=sequence
//ff:what TestBuildASTFromSQL_Basic — BIGSERIAL/VARCHAR/UNIQUE/DEFAULT 기본 파싱 확인
package migration

import "testing"

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
	assertBasicUsersTable(t, tbl)
}
