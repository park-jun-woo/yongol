//ff:func feature=migration type=test control=sequence
//ff:what 마이그레이션 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package migration

import (
	"testing"
)

func TestDiffTables_ZeroCov(t *testing.T) {
	prev, curr := bnSchemas()
	ps := NewSchema()
	ps.Tables["users"] = prev
	cs := NewSchema()
	cs.Tables["users"] = curr
	if ops := diffTables(ps, ps, cs, newEmptyHints()); len(ops) == 0 {
		t.Errorf("expected diff ops")
	}
}
