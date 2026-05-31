//ff:func feature=migration type=test control=sequence
//ff:what 마이그레이션 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package migration

import (
	"testing"
)

func TestDiffAddOrRenameTarget_ZeroCov(t *testing.T) {
	prev := NewSchema()
	c := &Table{Name: "newt", Columns: []*Column{{Name: "id"}}}
	ops := diffAddOrRenameTarget("newt", prev, c, newEmptyHints(), map[string]string{})
	if len(ops) == 0 {
		t.Errorf("expected create table ops")
	}
}
