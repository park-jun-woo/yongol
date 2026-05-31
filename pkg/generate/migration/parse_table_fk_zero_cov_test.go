//ff:func feature=migration type=test control=sequence
//ff:what 마이그레이션 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package migration

import (
	"testing"
)

func TestParseTableFK_ZeroCov(t *testing.T) {
	tbl := &Table{Name: "t"}
	fk := parseTableFK(tbl, "FOREIGN KEY (org_id) REFERENCES orgs (id)")
	if fk == nil || fk.RefTable != "orgs" {
		t.Errorf("table fk not parsed: %#v", fk)
	}
}
