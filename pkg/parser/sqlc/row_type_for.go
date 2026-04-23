//ff:func feature=orchestrator type=util control=selection
//ff:what rowTypeFor — sqlc 쿼리 cardinality 에 따른 row struct 이름 결정

package sqlc

// rowTypeFor returns the sqlc row struct name for a query, or "" when the
// query's cardinality produces no rows. sqlc default naming is "<Name>Row".
func rowTypeFor(name, cardinality string) string {
	switch cardinality {
	case "one", "many":
		return name + "Row"
	default:
		// "exec" / "execresult" / unknown → no row struct.
		return ""
	}
}
