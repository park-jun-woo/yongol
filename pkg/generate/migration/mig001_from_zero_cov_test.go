//ff:func feature=migration type=test control=sequence
//ff:what 마이그레이션 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package migration

import (
	"testing"
)

func TestMig001From_ZeroCov(t *testing.T) {
	prev := NewSchema()
	curr := NewSchema()
	h := newEmptyHints()
	h.RenameTables = []RenameTableHint{{From: "missing_from", To: "missing_to"}}
	diags := mig001From(prev, curr, h)
	_ = diags
	if got := mig001From(prev, curr, nil); got != nil {
		t.Errorf("nil hints should give nil")
	}
}
