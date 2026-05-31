//ff:func feature=migration type=test control=sequence
//ff:what 마이그레이션 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package migration

import (
	"testing"
)

func TestDispatchColumnAttr_ZeroCov(t *testing.T) {
	tbl := &Table{Name: "t"}
	col := &Column{Name: "c", Nullable: true}
	if step := dispatchColumnAttr(tbl, col, []string{"PRIMARY", "KEY"}, 0); step != 2 {
		t.Errorf("PRIMARY KEY step=%d want 2", step)
	}
	if len(tbl.PrimaryKey) != 1 {
		t.Errorf("PK not set")
	}
}
