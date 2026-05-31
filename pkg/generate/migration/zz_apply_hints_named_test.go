//ff:func feature=migration type=test control=sequence
//ff:what apply*Hint 단위 테스트 — 함수명 컨벤션에 맞춘 직접 커버 (전 분기)

package migration

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestApplyAllowDestructiveHint(t *testing.T) {
	h := newEmptyHints()
	applyAllowDestructiveHint(h, ddl.HintComment{Tag: "allow_destructive", TableCtx: "users"})
	if !h.AllowDestructive["users"] {
		t.Errorf("expected users marked destructive")
	}
	// empty TableCtx → no-op.
	applyAllowDestructiveHint(h, ddl.HintComment{Tag: "allow_destructive"})
	if len(h.AllowDestructive) != 1 {
		t.Errorf("empty table ctx should be ignored")
	}
}

func TestApplyDataMigrationHint(t *testing.T) {
	h := newEmptyHints()
	applyDataMigrationHint(h, ddl.HintComment{TableCtx: "users", Args: map[string]string{"file": "u.sql"}})
	if h.DataMigrations["users"] != "u.sql" {
		t.Errorf("data migration not stored: %v", h.DataMigrations)
	}
	// missing file → no-op.
	applyDataMigrationHint(h, ddl.HintComment{TableCtx: "x", Args: map[string]string{}})
	// missing table ctx → no-op.
	applyDataMigrationHint(h, ddl.HintComment{Args: map[string]string{"file": "y.sql"}})
	if len(h.DataMigrations) != 1 {
		t.Errorf("invalid data migration hints should be ignored: %v", h.DataMigrations)
	}
}

func TestApplyBackfillHint(t *testing.T) {
	h := newEmptyHints()
	applyBackfillHint(h, ddl.HintComment{TableCtx: "users", ColumnCtx: "age", Args: map[string]string{"default": "0"}})
	if len(h.Backfills) != 1 {
		t.Errorf("backfill not stored: %v", h.Backfills)
	}
	// missing default → no-op.
	applyBackfillHint(h, ddl.HintComment{TableCtx: "users", ColumnCtx: "x", Args: map[string]string{}})
	// missing column ctx → no-op.
	applyBackfillHint(h, ddl.HintComment{TableCtx: "users", Args: map[string]string{"default": "1"}})
	if len(h.Backfills) != 1 {
		t.Errorf("invalid backfill hints should be ignored: %v", h.Backfills)
	}
}

func TestApplyCastHint(t *testing.T) {
	h := newEmptyHints()
	applyCastHint(h, ddl.HintComment{TableCtx: "users", ColumnCtx: "age", Args: map[string]string{"using": "age::int"}})
	if len(h.Casts) != 1 {
		t.Errorf("cast not stored: %v", h.Casts)
	}
	// missing using → no-op.
	applyCastHint(h, ddl.HintComment{TableCtx: "users", ColumnCtx: "x", Args: map[string]string{}})
	// missing column ctx → no-op.
	applyCastHint(h, ddl.HintComment{TableCtx: "users", Args: map[string]string{"using": "x"}})
	if len(h.Casts) != 1 {
		t.Errorf("invalid cast hints should be ignored: %v", h.Casts)
	}
}

func TestApplyRenameHint(t *testing.T) {
	h := newEmptyHints()
	// column rename (column ctx).
	applyRenameHint(h, ddl.HintComment{TableCtx: "users", ColumnCtx: "fullname", Args: map[string]string{"from": "name"}})
	if len(h.RenameColumns) != 1 {
		t.Errorf("column rename not stored")
	}
	// table rename (block above).
	applyRenameHint(h, ddl.HintComment{BlockAbove: true, Args: map[string]string{"from": "old", "to": "new"}})
	if len(h.RenameTables) != 1 {
		t.Errorf("table rename not stored")
	}
	// table-context column rename (from+to+tablectx, no column ctx).
	applyRenameHint(h, ddl.HintComment{TableCtx: "users", Args: map[string]string{"from": "a", "to": "b"}})
	if len(h.RenameColumns) != 2 {
		t.Errorf("table-ctx column rename not stored: %v", h.RenameColumns)
	}
}

func TestApplyHintComment_Dispatch(t *testing.T) {
	h := newEmptyHints()
	for _, c := range []ddl.HintComment{
		{Tag: "rename", TableCtx: "t", ColumnCtx: "c", Args: map[string]string{"from": "x"}},
		{Tag: "cast", TableCtx: "t", ColumnCtx: "c", Args: map[string]string{"using": "u"}},
		{Tag: "backfill", TableCtx: "t", ColumnCtx: "c", Args: map[string]string{"default": "0"}},
		{Tag: "data_migration", TableCtx: "t", Args: map[string]string{"file": "f"}},
		{Tag: "allow_destructive", TableCtx: "t"},
		{Tag: "unknown"}, // default branch — no-op.
	} {
		applyHintComment(h, c)
	}
	if !h.AllowDestructive["t"] {
		t.Errorf("dispatch failed for allow_destructive")
	}
}

func TestApplyRenameHints(t *testing.T) {
	prev := mustAST(t, `CREATE TABLE old_users (id BIGINT PRIMARY KEY, name TEXT NOT NULL);`)
	h := newEmptyHints()
	h.RenameTables = []RenameTableHint{{From: "old_users", To: "users"}}
	h.RenameColumns = []RenameColumnHint{{Table: "users", From: "name", To: "full_name"}}
	out := applyRenameHints(prev, h)
	if _, ok := out.Tables["users"]; !ok {
		t.Errorf("table not renamed to users: %v", out.Tables)
	}
	// nil hints → unchanged.
	if applyRenameHints(prev, nil) != prev {
		t.Errorf("nil hints should return prev unchanged")
	}
	// no rename rules → unchanged.
	if applyRenameHints(prev, newEmptyHints()) != prev {
		t.Errorf("empty rules should return prev unchanged")
	}
}
