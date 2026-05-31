//ff:func feature=migration type=test control=sequence
//ff:what 마이그레이션 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package migration

import (
	"testing"
)

func TestDiffForeignKeys_ZeroCov(t *testing.T) {
	prev, curr := bnSchemas()
	ops := diffForeignKeys(prev, curr, "users")
	if len(ops) == 0 {
		t.Errorf("expected fk diff ops")
	}
}
