//ff:func feature=migration type=test control=sequence
//ff:what 마이그레이션 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package migration

import (
	"testing"
)

func TestParseNamedConstraint_ZeroCov(t *testing.T) {
	tbl := &Table{Name: "t"}
	parseNamedConstraint(tbl, "CONSTRAINT t_pkey PRIMARY KEY (id)")
	if len(tbl.PrimaryKey) == 0 {
		t.Errorf("named PK not parsed")
	}
}
