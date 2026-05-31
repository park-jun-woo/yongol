//ff:func feature=migration type=test control=sequence
//ff:what 마이그레이션 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package migration

import (
	"testing"
)

func TestDiffTableBody_ZeroCov(t *testing.T) {
	prev, curr := bnSchemas()
	if ops := diffTableBody(prev, curr, newEmptyHints(), "users"); len(ops) == 0 {
		t.Errorf("expected table body ops")
	}
	if ops := diffTableBody(nil, curr, nil, "users"); ops != nil {
		t.Errorf("nil prev should give nil")
	}
}
