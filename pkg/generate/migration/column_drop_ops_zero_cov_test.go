//ff:func feature=migration type=test control=sequence
//ff:what 마이그레이션 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package migration

import (
	"testing"
)

func TestColumnDropOps_ZeroCov(t *testing.T) {
	ops := columnDropOps("t", []string{"old"}, map[string]*Column{}, map[string]bool{}, newEmptyHints())
	if len(ops) != 1 {
		t.Errorf("expected one drop op, got %d", len(ops))
	}
}
