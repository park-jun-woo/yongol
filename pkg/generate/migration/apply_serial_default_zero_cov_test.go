//ff:func feature=migration type=test control=sequence
//ff:what 마이그레이션 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package migration

import (
	"testing"
)

func TestApplySerialDefault_ZeroCov(t *testing.T) {
	tbl := &Table{Name: "t"}
	col := &Column{Name: "id", Nullable: true}
	applySerialDefault(tbl, col, true)
	if col.Default == "" || col.Nullable {
		t.Errorf("serial default not applied: %#v", col)
	}
}
