//ff:func feature=migration type=test control=sequence
//ff:what apply*Hint 단위 테스트 — 함수명 컨벤션에 맞춘 직접 커버 (전 분기)
package migration

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

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
