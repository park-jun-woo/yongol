//ff:func feature=migration type=test control=sequence
//ff:what TestParseHints_Backfill — @backfill default=... 힌트가 Backfills 맵에 등록
package migration

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestParseHints_Backfill(t *testing.T) {
	comments := []ddl.HintComment{
		{Tag: "backfill", Args: map[string]string{"default": "false"}, TableCtx: "users", ColumnCtx: "email_verified"},
	}
	h := ParseHints(comments)
	if got := h.Backfills[colKey{Table: "users", Column: "email_verified"}]; got != "false" {
		t.Errorf("backfill wrong: %q", got)
	}
}
