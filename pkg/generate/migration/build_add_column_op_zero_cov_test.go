//ff:func feature=migration type=test control=sequence
//ff:what 마이그레이션 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package migration

import (
	"testing"
)

func TestBuildAddColumnOp_ZeroCov(t *testing.T) {
	col := &Column{Name: "age", Type: CanonicalType{Base: "INTEGER"}}
	h := newEmptyHints()
	h.Backfills[colKey{Table: "t", Column: "age"}] = "0"
	op := buildAddColumnOp("t", "age", col, h)
	if op.Backfill != "0" || op.Table != "t" {
		t.Errorf("add col op wrong: %#v", op)
	}
}
