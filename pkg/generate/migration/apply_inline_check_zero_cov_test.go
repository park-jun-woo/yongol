//ff:func feature=migration type=test control=sequence
//ff:what 마이그레이션 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package migration

import (
	"testing"
)

func TestApplyInlineCheck_ZeroCov(t *testing.T) {
	tbl := &Table{Name: "t"}
	col := &Column{Name: "c"}
	c := applyInlineCheck(tbl, col, []string{"CHECK", "(c > 0)"}, 0)
	if c != 2 || len(tbl.Checks) != 1 {
		t.Errorf("inline check not added: consumed=%d checks=%d", c, len(tbl.Checks))
	}
}
