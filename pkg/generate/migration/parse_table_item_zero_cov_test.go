//ff:func feature=migration type=test control=sequence
//ff:what 마이그레이션 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package migration

import (
	"testing"
)

func TestParseTableItem_ZeroCov(t *testing.T) {
	tbl := &Table{Name: "t"}
	parseTableItem(tbl, "PRIMARY KEY (id)")
	if len(tbl.PrimaryKey) == 0 {
		t.Errorf("PK item not handled")
	}
	parseTableItem(tbl, "name TEXT")
	if len(tbl.Columns) == 0 {
		t.Errorf("column item not handled")
	}
}
