//ff:func feature=migration type=test control=sequence
//ff:what TestTableHasColumn — 컬럼 이름 존재 여부 선형 탐색
package migration

import (
	"testing"
)

func TestTableHasColumn(t *testing.T) {
	tbl := &Table{Columns: []*Column{{Name: "id"}, {Name: "email"}}}
	if !tableHasColumn(tbl, "email") {
		t.Error("tableHasColumn(email) = false, want true")
	}
	if tableHasColumn(tbl, "missing") {
		t.Error("tableHasColumn(missing) = true, want false")
	}
}
