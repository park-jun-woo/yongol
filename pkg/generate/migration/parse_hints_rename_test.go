//ff:func feature=migration type=test control=sequence
//ff:what TestParseHints_Rename — @rename 컬럼/테이블 분리 파싱
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
