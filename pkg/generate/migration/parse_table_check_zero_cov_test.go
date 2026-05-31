//ff:func feature=migration type=test control=sequence
//ff:what 마이그레이션 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package migration

import (
	"testing"
)

func TestParseTableCheck_ZeroCov(t *testing.T) {
	tbl := &Table{Name: "t"}
	c := parseTableCheck(tbl, "", "CHECK (age > 0)")
	if c == nil || c.Name == "" || c.Expression == "" {
		t.Errorf("check not parsed: %#v", c)
	}
}
