//ff:func feature=migration type=test control=sequence
//ff:what 마이그레이션 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package migration

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestCollectDiags_ZeroCov(t *testing.T) {
	prev := NewSchema()
	curr := NewSchema()
	diags := collectDiags([]diagnostic.Diagnostic{{Message: "X"}}, prev, curr, newEmptyHints(), nil, nil)
	if len(diags) == 0 {
		t.Errorf("expected at least the seeded diag")
	}
}
