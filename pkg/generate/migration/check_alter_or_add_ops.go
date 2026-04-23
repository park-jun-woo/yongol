//ff:func feature=migration type=util control=iteration dimension=1
//ff:what checkAlterOrAddOps — curr CHECK 가 prev 와 Expression 다르면 Drop+Add, 없으면 Add
package migration

// checkAlterOrAddOps emits Add / (Drop+Add) ops for CHECK constraints
// present in curr.
func checkAlterOrAddOps(tableName string, currNames []string, prevMap, currMap map[string]*CheckConstraint) []Operation {
	var ops []Operation
	for _, n := range currNames {
		ops = append(ops, checkDiffForOne(tableName, n, prevMap, currMap)...)
	}
	return ops
}
