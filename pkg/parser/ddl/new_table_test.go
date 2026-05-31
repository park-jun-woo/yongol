//ff:func feature=manifest type=test control=sequence
//ff:what dispatchConstraint — 라인 종류별 분기 (FK / PRIMARY / UNIQUE / CHECK / column)
package ddl

func newTable() *Table {
	return &Table{Columns: map[string]Column{}}
}
