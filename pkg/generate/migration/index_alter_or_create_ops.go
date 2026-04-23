//ff:func feature=migration type=util control=iteration dimension=1
//ff:what indexAlterOrCreateOps — 인덱스 정의가 바뀌면 Drop+Create, 없으면 Create 만
package migration

// indexAlterOrCreateOps emits Create / (Drop+Create) ops for indexes in
// curr.
func indexAlterOrCreateOps(tableName string, currNames []string, prevMap, currMap map[string]*Index) []Operation {
	var ops []Operation
	for _, n := range currNames {
		ops = append(ops, indexDiffForOne(tableName, n, prevMap, currMap)...)
	}
	return ops
}
