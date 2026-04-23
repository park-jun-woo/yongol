//ff:type feature=migration type=model
//ff:what colKey — (table, column) 스코프 힌트 맵 키
package migration

// colKey is the map key for (table,column) scoped hints.
type colKey struct {
	Table, Column string
}
