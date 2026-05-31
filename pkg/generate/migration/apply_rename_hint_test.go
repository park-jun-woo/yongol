//ff:func feature=migration type=test control=sequence
//ff:what apply*Hint 단위 테스트 — 함수명 컨벤션에 맞춘 직접 커버 (전 분기)
package migration

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestApplyRenameHint(t *testing.T) {
	h := newEmptyHints()
	// column rename (column ctx).
	applyRenameHint(h, ddl.HintComment{TableCtx: "users", ColumnCtx: "fullname", Args: map[string]string{"from": "name"}})
	if len(h.RenameColumns) != 1 {
		t.Errorf("column rename not stored")
	}
	// table rename (block above).
	applyRenameHint(h, ddl.HintComment{BlockAbove: true, Args: map[string]string{"from": "old", "to": "new"}})
	if len(h.RenameTables) != 1 {
		t.Errorf("table rename not stored")
	}
	// table-context column rename (from+to+tablectx, no column ctx).
	applyRenameHint(h, ddl.HintComment{TableCtx: "users", Args: map[string]string{"from": "a", "to": "b"}})
	if len(h.RenameColumns) != 2 {
		t.Errorf("table-ctx column rename not stored: %v", h.RenameColumns)
	}
}
