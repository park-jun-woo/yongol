//ff:func feature=migration type=test control=sequence
//ff:what TestMigrationE2EZeroCov — ParseHints / BuildASTFromSQL / Diff / ApplyHintsToOps / EmitSQL 풀 파이프라인 커버
package migration

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

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
