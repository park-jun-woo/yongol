//ff:func feature=migration type=test control=sequence
//ff:what 마이그레이션 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package migration

import (
	"testing"
)

func TestDiffOneTable_ZeroCov(t *testing.T) {
	prev, curr := bnSchemas()
	ps := NewSchema()
	ps.Tables["users"] = prev
	cs := NewSchema()
	cs.Tables["users"] = curr
	ops := diffOneTable("users", ps, ps, cs, newEmptyHints(), map[string]string{}, map[string]string{})
	if len(ops) == 0 {
		t.Errorf("expected body diff ops")
	}
}
