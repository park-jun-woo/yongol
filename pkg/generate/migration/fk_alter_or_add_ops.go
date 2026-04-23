//ff:func feature=migration type=util control=iteration dimension=1
//ff:what fkAlterOrAddOps — curr FK 가 prev 와 다르면 Drop+Add, 없으면 Add 만
package migration

// fkAlterOrAddOps emits Add / (Drop+Add) operations for FKs present in
// curr.
func fkAlterOrAddOps(tableName string, currNames []string, prevMap, currMap map[string]*ForeignKey) []Operation {
	var ops []Operation
	for _, n := range currNames {
		ops = append(ops, fkDiffForOne(tableName, n, prevMap, currMap)...)
	}
	return ops
}
