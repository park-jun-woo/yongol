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
