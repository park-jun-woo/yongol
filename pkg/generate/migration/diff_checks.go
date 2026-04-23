//ff:func feature=migration type=util control=sequence
//ff:what diffChecks — CHECK 제약 추가/삭제/변경 diff
package migration

// diffChecks compares the CHECK constraints of prev and curr and emits
// DropCheck / AddCheck as needed.
func diffChecks(prev, curr *Table, tableName string) []Operation {
	prevMap := checkMap(prev.Checks)
	currMap := checkMap(curr.Checks)
	prevNames := sortedMapKeys(prevMap)
	currNames := sortedMapKeys(currMap)

	var ops []Operation
	ops = append(ops, checkDropOps(tableName, prevNames, currMap)...)
	ops = append(ops, checkAlterOrAddOps(tableName, currNames, prevMap, currMap)...)
	return ops
}
