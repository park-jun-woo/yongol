//ff:func feature=migration type=test control=sequence
//ff:what 마이그레이션 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package migration

import (
	"testing"
)

func TestFkDiffForOne_ZeroCov(t *testing.T) {
	prevMap := map[string]*ForeignKey{"fk": {Name: "fk", Columns: []string{"a"}, RefTable: "o", RefColumns: []string{"id"}}}
	currMap := map[string]*ForeignKey{"fk": {Name: "fk", Columns: []string{"a"}, RefTable: "o", RefColumns: []string{"id"}, OnDelete: "CASCADE"}}
	ops := fkDiffForOne("t", "fk", prevMap, currMap)
	if len(ops) != 2 {
		t.Errorf("changed fk should drop+add, got %d", len(ops))
	}
	// equal → nil
	if got := fkDiffForOne("t", "fk", prevMap, prevMap); got != nil {
		t.Errorf("equal fk should be nil")
	}
}
