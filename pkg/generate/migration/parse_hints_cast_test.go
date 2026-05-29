//ff:func feature=migration type=test control=sequence
//ff:what TestParseHints_Cast — @cast using=... 힌트가 Casts 맵에 등록
package migration

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestParseHints_Cast(t *testing.T) {
	comments := []ddl.HintComment{
		{Tag: "cast", Args: map[string]string{"using": "col::integer"}, TableCtx: "t", ColumnCtx: "id"},
	}
	h := ParseHints(comments)
	if got := h.Casts[colKey{Table: "t", Column: "id"}]; got != "col::integer" {
		t.Errorf("cast expr wrong: %q", got)
	}
}
