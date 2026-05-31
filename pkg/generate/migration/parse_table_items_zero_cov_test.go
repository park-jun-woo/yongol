//ff:func feature=migration type=test control=sequence
//ff:what 마이그레이션 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package migration

import (
	"testing"
)

func TestParseTableItems_ZeroCov(t *testing.T) {
	tbl := &Table{Name: "t"}
	parseTableItems(tbl, "id BIGINT, name TEXT, PRIMARY KEY (id)")
	if len(tbl.Columns) != 2 || len(tbl.PrimaryKey) == 0 {
		t.Errorf("items not parsed: cols=%d pk=%v", len(tbl.Columns), tbl.PrimaryKey)
	}
}
