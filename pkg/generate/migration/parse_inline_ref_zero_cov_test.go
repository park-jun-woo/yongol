//ff:func feature=migration type=test control=sequence
//ff:what 마이그레이션 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package migration

import (
	"testing"
)

func TestParseInlineRef_ZeroCov(t *testing.T) {
	tbl := &Table{Name: "t"}
	fk, consumed := parseInlineRef(tbl, "org_id", []string{"orgs(id)"})
	if fk == nil || fk.RefTable != "orgs" || consumed == 0 {
		t.Errorf("inline ref wrong: %#v consumed=%d", fk, consumed)
	}
}
