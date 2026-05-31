//ff:func feature=migration type=test control=iteration dimension=1
//ff:what apply*Hint 단위 테스트 — 함수명 컨벤션에 맞춘 직접 커버 (전 분기)
package migration

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

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
