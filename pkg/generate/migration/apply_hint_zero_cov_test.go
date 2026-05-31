//ff:func feature=migration type=test control=sequence
//ff:what 마이그레이션 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package migration

import (
	"testing"
)

func TestApplyHint_ZeroCov(t *testing.T) {
	h := newEmptyHints()
	h.AllowDestructive["t"] = true
	out := applyHint(DropTable{Name: "t"}, h)
	if dt, ok := out.(DropTable); !ok || !dt.AllowDestructive {
		t.Errorf("DropTable allow-destructive not applied: %#v", out)
	}
	// non-hint-aware op returned unchanged
	if got := applyHint(CreateTable{Table: &Table{Name: "x"}}, h); got == nil {
		t.Errorf("nil returned")
	}
}
