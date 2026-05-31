//ff:func feature=migration type=test control=sequence
//ff:what apply*Hint 단위 테스트 — 함수명 컨벤션에 맞춘 직접 커버 (전 분기)
package migration

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

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
