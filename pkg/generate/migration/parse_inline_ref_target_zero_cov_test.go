//ff:func feature=migration type=test control=sequence
//ff:what 마이그레이션 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package migration

import (
	"testing"
)

func TestParseInlineRefTarget_ZeroCov(t *testing.T) {
	rt, rc, c := parseInlineRefTarget([]string{"orgs(id)"})
	if rt != "orgs" || rc != "id" || c != 1 {
		t.Errorf("target wrong: %s %s %d", rt, rc, c)
	}
}
