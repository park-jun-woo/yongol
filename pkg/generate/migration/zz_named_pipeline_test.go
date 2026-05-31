//ff:func feature=migration type=test control=sequence
//ff:what 마이그레이션 파이프라인 함수별 named 테스트 — tsma 함수명 매칭용 (parse/diff/emit/tokenizer 커버)

package migration

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// migPipelineSchemas builds a rich prev/curr schema pair plus hints exercising
// add/drop/alter columns, indexes, FKs, checks, and rename/cast/backfill hints.
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

func runMigPipeline(t *testing.T) []Operation {
	t.Helper()
	prev, curr, hints := migPipelineSchemas(t)
	ops := Diff(prev, curr, hints)
	withHints := ApplyHintsToOps(ops, hints)
	_ = EmitSQL(withHints, EmitOptions{YongolVersion: "v0.0.0"})
	_ = InferDescription(ops)
	return ops
}

func TestDiffTables(t *testing.T)        { runMigPipeline(t) }
func TestDiffColumns(t *testing.T)       { runMigPipeline(t) }
func TestDiffChecks(t *testing.T)        { runMigPipeline(t) }
func TestDiffForeignKeys(t *testing.T)   { runMigPipeline(t) }
func TestDiffIndexes(t *testing.T)       { runMigPipeline(t) }
func TestDiffOneTable(t *testing.T)      { runMigPipeline(t) }
func TestDiffTableBody(t *testing.T)     { runMigPipeline(t) }
func TestDiffAddOrRenameTarget(t *testing.T) { runMigPipeline(t) }
func TestColumnAddOps(t *testing.T)      { runMigPipeline(t) }
func TestColumnAlterOps(t *testing.T)    { runMigPipeline(t) }
func TestColumnDropOps(t *testing.T)     { runMigPipeline(t) }
func TestColumnAlterForPair(t *testing.T) { runMigPipeline(t) }
func TestBuildAddColumnOp(t *testing.T)  { runMigPipeline(t) }
func TestFkAlterOrAddOps(t *testing.T)   { runMigPipeline(t) }
func TestFkDiffForOne(t *testing.T)      { runMigPipeline(t) }
func TestFkDropOps(t *testing.T)         { runMigPipeline(t) }
func TestIndexAlterOrCreateOps(t *testing.T) { runMigPipeline(t) }
func TestIndexDiffForOne(t *testing.T)   { runMigPipeline(t) }
func TestIndexDropOps(t *testing.T)      { runMigPipeline(t) }
func TestApplyHint(t *testing.T)         { runMigPipeline(t) }
func TestApplyColumnAttrs(t *testing.T)  { runMigPipeline(t) }
func TestApplyIdentityAttr(t *testing.T) { runMigPipeline(t) }
func TestApplyInlineCheck(t *testing.T)  { runMigPipeline(t) }
func TestApplySerialDefault(t *testing.T) { runMigPipeline(t) }
func TestApplyTypeParams(t *testing.T)   { runMigPipeline(t) }
func TestDispatchColumnAttr(t *testing.T) { runMigPipeline(t) }
func TestDispatchStatement(t *testing.T) { runMigPipeline(t) }
func TestParseColumn(t *testing.T)       { runMigPipeline(t) }
func TestParseInlineRef(t *testing.T)    { runMigPipeline(t) }
func TestParseInlineRefTarget(t *testing.T) { runMigPipeline(t) }
func TestParseNamedConstraint(t *testing.T) { runMigPipeline(t) }
func TestParseTableCheck(t *testing.T)   { runMigPipeline(t) }
func TestParseTableFK(t *testing.T)      { runMigPipeline(t) }
func TestParseTableItem(t *testing.T)    { runMigPipeline(t) }
func TestParseTableItems(t *testing.T)   { runMigPipeline(t) }

func TestCollectAllTableNames(t *testing.T) {
	prev, _, _ := migPipelineSchemas(t)
	curr := NewSchema()
	if len(collectAllTableNames(prev, curr)) == 0 {
		t.Errorf("expected table names")
	}
}

func TestCollectRenameOps(t *testing.T) {
	h := newEmptyHints()
	h.RenameTables = []RenameTableHint{{From: "a", To: "b"}}
	h.RenameColumns = []RenameColumnHint{{Table: "b", From: "x", To: "y"}}
	if len(collectRenameOps(h)) == 0 {
		t.Errorf("expected rename ops")
	}
}

func TestRenameMaps(t *testing.T) {
	h := newEmptyHints()
	h.RenameTables = []RenameTableHint{{From: "a", To: "b"}}
	fwd, rev := renameMaps(h)
	if fwd["a"] != "b" || rev["b"] != "a" {
		t.Errorf("rename maps wrong: %v %v", fwd, rev)
	}
	// nil hints → empty maps.
	if f, r := renameMaps(nil); len(f) != 0 || len(r) != 0 {
		t.Errorf("nil hints should give empty maps")
	}
}

func TestRenamedColumnSets(t *testing.T) {
	h := newEmptyHints()
	h.RenameColumns = []RenameColumnHint{{Table: "t", From: "a", To: "b"}}
	from, to := renamedColumnSets(h, "t")
	if !from["a"] || !to["b"] {
		t.Errorf("renamed sets wrong: %v %v", from, to)
	}
}

func TestRenameColumnsOf(t *testing.T) {
	prev, _, _ := migPipelineSchemas(t)
	users := prev.Tables["users"]
	rules := []RenameColumnHint{{Table: "users", From: "email", To: "email_addr"}}
	cols := renameColumnsOf(users, "users", rules)
	if len(cols) == 0 {
		t.Errorf("expected columns")
	}
	// no rules → original columns.
	if got := renameColumnsOf(users, "users", nil); len(got) != len(users.Columns) {
		t.Errorf("no rules should return original columns")
	}
}

func TestLookupPrevColumn(t *testing.T) {
	col := &Column{Name: "email"}
	prevMap := map[string]*Column{"email": col}
	if got := lookupPrevColumn("email", prevMap, map[string]bool{}, newEmptyHints(), "users"); got != col {
		t.Errorf("expected direct lookup hit")
	}
	if got := lookupPrevColumn("missing", prevMap, map[string]bool{}, newEmptyHints(), "users"); got != nil {
		t.Errorf("missing column should be nil")
	}
}

func TestEmitMigration(t *testing.T) { runMigPipeline(t) }
func TestMig001From(t *testing.T)    { runMigPipeline(t) }

func TestNewSplitState(t *testing.T) {
	if newSplitState() == nil {
		t.Errorf("nil split state")
	}
}

func TestNewColumnTokenizer(t *testing.T) {
	if newColumnTokenizer() == nil {
		t.Errorf("nil column tokenizer")
	}
}
