//ff:func feature=migration type=test control=sequence
//ff:what 마이그레이션 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package migration

import (
	"testing"
)

func TestFkDropOps_ZeroCov(t *testing.T) {
	ops := fkDropOps("t", []string{"fk"}, map[string]*ForeignKey{})
	if len(ops) != 1 {
		t.Errorf("expected drop fk op, got %d", len(ops))
	}
}
