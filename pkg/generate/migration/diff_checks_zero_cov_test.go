//ff:func feature=migration type=test control=sequence
//ff:what 마이그레이션 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package migration

import (
	"testing"
)

func TestDiffChecks_ZeroCov(t *testing.T) {
	prev, curr := bnSchemas()
	ops := diffChecks(prev, curr, "users")
	if len(ops) == 0 {
		t.Errorf("expected check diff ops")
	}
}
