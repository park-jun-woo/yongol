//ff:func feature=migration type=test control=sequence
//ff:what TestMigrationE2EZeroCov — ParseHints / BuildASTFromSQL / Diff / ApplyHintsToOps / EmitSQL 풀 파이프라인 커버
package migration

import (
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
