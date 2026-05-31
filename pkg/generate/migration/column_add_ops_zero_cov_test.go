//ff:func feature=migration type=test control=sequence
//ff:what 마이그레이션 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package migration

import (
	"testing"
)

func TestColumnAddOps_ZeroCov(t *testing.T) {
	currMap := map[string]*Column{"age": {Name: "age"}}
	ops := columnAddOps("t", []string{"age"}, map[string]*Column{}, map[string]bool{}, currMap, newEmptyHints())
	if len(ops) != 1 {
		t.Errorf("expected one add op, got %d", len(ops))
	}
}
