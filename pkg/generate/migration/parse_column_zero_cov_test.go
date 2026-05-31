//ff:func feature=migration type=test control=sequence
//ff:what 마이그레이션 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package migration

import (
	"testing"
)

func TestParseColumn_ZeroCov(t *testing.T) {
	tbl := &Table{Name: "t"}
	parseColumn(tbl, "id BIGINT NOT NULL")
	if len(tbl.Columns) != 1 || tbl.Columns[0].Name != "id" {
		t.Errorf("column not parsed: %#v", tbl.Columns)
	}
}
