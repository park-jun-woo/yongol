//ff:func feature=migration type=test control=sequence
//ff:what 마이그레이션 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package migration

import (
	"testing"
)

func TestApplyIdentityAttr_ZeroCov(t *testing.T) {
	tbl := &Table{Name: "t"}
	col := &Column{Name: "id", Nullable: true}
	rest := []string{"GENERATED", "ALWAYS", "AS", "IDENTITY"}
	if c := applyIdentityAttr(tbl, col, rest, 0); c != 4 {
		t.Errorf("consumed=%d want 4", c)
	}
	if col.Identity == nil || !col.Identity.Always || col.Nullable {
		t.Errorf("identity not set: %#v", col)
	}
	// non-GENERATED returns 0
	if c := applyIdentityAttr(tbl, col, []string{"FOO"}, 0); c != 0 {
		t.Errorf("expected 0 for non-GENERATED")
	}
}
