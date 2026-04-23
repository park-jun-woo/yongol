//ff:func feature=migration type=test control=sequence
//ff:what TestParseHints_AllowDestructive — @allow_destructive 힌트가 테이블에 기록
package migration

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestParseHints_AllowDestructive(t *testing.T) {
	comments := []ddl.HintComment{
		{Tag: "allow_destructive", TableCtx: "old_table", BlockAbove: true},
	}
	h := ParseHints(comments)
	if !h.AllowDestructive["old_table"] {
		t.Errorf("allow_destructive not set: %+v", h.AllowDestructive)
	}
}
