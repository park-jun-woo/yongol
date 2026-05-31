//ff:func feature=migration type=test control=sequence
//ff:what 마이그레이션 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package migration

import (
	"testing"
)

func TestFkAlterOrAddOps_ZeroCov(t *testing.T) {
	currMap := map[string]*ForeignKey{"fk": {Name: "fk", Columns: []string{"a"}, RefTable: "o", RefColumns: []string{"id"}}}
	ops := fkAlterOrAddOps("t", []string{"fk"}, map[string]*ForeignKey{}, currMap)
	if len(ops) != 1 {
		t.Errorf("expected add fk op, got %d", len(ops))
	}
}
