//ff:func feature=migration type=test control=sequence
//ff:what Diff e2e — zenflow users 테이블에 컬럼 1개 추가 시나리오
package migration

import (
	"strings"
	"testing"
)

func TestDiffE2E_AddUsersEmailVerified(t *testing.T) {
	prev := mustAST(t, `
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE
);`)
	curr := mustAST(t, `
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    email_verified BOOLEAN NOT NULL DEFAULT false
);`)
	ops := Diff(prev, curr, nil)
	if len(ops) != 1 {
		t.Fatalf("expected 1 op, got %d: %+v", len(ops), ops)
	}
	desc := InferDescription(ops)
	if !strings.Contains(desc, "add_users_email_verified") {
		t.Errorf("desc = %q, want contains 'add_users_email_verified'", desc)
	}
	sql := EmitSQL(ops, EmitOptions{})
	if !strings.Contains(sql, "ALTER TABLE users ADD COLUMN email_verified") {
		t.Errorf("SQL missing ADD COLUMN: %s", sql)
	}
}
