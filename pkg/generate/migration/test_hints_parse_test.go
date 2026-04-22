//ff:func feature=migration type=test control=iteration dimension=1
//ff:what Hints 파서 매트릭스 — @rename / @cast / @backfill / @data_migration / @allow_destructive
package migration

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestParseHints_Rename(t *testing.T) {
	comments := []ddl.HintComment{
		{Tag: "rename", Args: map[string]string{"from": "email"}, TableCtx: "users", ColumnCtx: "email_address"},
		{Tag: "rename", Args: map[string]string{"from": "members", "to": "users"}, TableCtx: "users", BlockAbove: true},
	}
	h := ParseHints(comments)
	if len(h.RenameColumns) != 1 {
		t.Fatalf("expected 1 rename column, got %d: %+v", len(h.RenameColumns), h.RenameColumns)
	}
	if h.RenameColumns[0].Table != "users" || h.RenameColumns[0].From != "email" || h.RenameColumns[0].To != "email_address" {
		t.Errorf("column rename wrong: %+v", h.RenameColumns[0])
	}
	if len(h.RenameTables) != 1 || h.RenameTables[0].From != "members" || h.RenameTables[0].To != "users" {
		t.Errorf("table rename wrong: %+v", h.RenameTables)
	}
}

func TestParseHints_Cast(t *testing.T) {
	comments := []ddl.HintComment{
		{Tag: "cast", Args: map[string]string{"using": "col::integer"}, TableCtx: "t", ColumnCtx: "id"},
	}
	h := ParseHints(comments)
	if got := h.Casts[colKey{Table: "t", Column: "id"}]; got != "col::integer" {
		t.Errorf("cast expr wrong: %q", got)
	}
}

func TestParseHints_Backfill(t *testing.T) {
	comments := []ddl.HintComment{
		{Tag: "backfill", Args: map[string]string{"default": "false"}, TableCtx: "users", ColumnCtx: "email_verified"},
	}
	h := ParseHints(comments)
	if got := h.Backfills[colKey{Table: "users", Column: "email_verified"}]; got != "false" {
		t.Errorf("backfill wrong: %q", got)
	}
}

func TestParseHints_DataMigration(t *testing.T) {
	comments := []ddl.HintComment{
		{Tag: "data_migration", Args: map[string]string{"file": "migrations_data/0042.sql"}, TableCtx: "users", BlockAbove: true},
	}
	h := ParseHints(comments)
	if got := h.DataMigrations["users"]; got != "migrations_data/0042.sql" {
		t.Errorf("data_migration wrong: %q", got)
	}
}

func TestParseHints_AllowDestructive(t *testing.T) {
	comments := []ddl.HintComment{
		{Tag: "allow_destructive", TableCtx: "old_table", BlockAbove: true},
	}
	h := ParseHints(comments)
	if !h.AllowDestructive["old_table"] {
		t.Errorf("allow_destructive not set: %+v", h.AllowDestructive)
	}
}
