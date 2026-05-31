//ff:func feature=migration type=test control=sequence
//ff:what 마이그레이션 파이프라인 함수별 named 테스트 — tsma 함수명 매칭용 (parse/diff/emit/tokenizer 커버)
package migration

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func migPipelineSchemas(t *testing.T) (prev, curr *Schema, hints *Hints) {
	t.Helper()
	prev = mustAST(t, `
CREATE TABLE orgs (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL
);
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    age INTEGER,
    legacy_col TEXT,
    CHECK (age > 0)
);
CREATE INDEX idx_users_email ON users (email);
CREATE TABLE gone (id BIGINT PRIMARY KEY);`)

	curr = mustAST(t, `
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

	hints = ParseHints([]ddl.HintComment{
		{Tag: "cast", TableCtx: "users", ColumnCtx: "age", Args: map[string]string{"using": "age::bigint"}},
		{Tag: "backfill", TableCtx: "users", ColumnCtx: "age", Args: map[string]string{"default": "0"}},
		{Tag: "allow_destructive", TableCtx: "gone"},
		{Tag: "data_migration", TableCtx: "users", Args: map[string]string{"file": "u.sql"}},
	})
	return prev, curr, hints
}
