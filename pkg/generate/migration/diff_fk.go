//ff:func feature=migration type=util control=sequence
//ff:what diffForeignKeys — FK 추가/삭제/변경 diff
package migration

// diffForeignKeys compares the FKs of prev and curr and emits
// DropForeignKey / AddForeignKey as needed.
func diffForeignKeys(prev, curr *Table, tableName string) []Operation {
	prevMap := fkMap(prev.ForeignKeys)
	currMap := fkMap(curr.ForeignKeys)
	prevNames := sortedMapKeys(prevMap)
	currNames := sortedMapKeys(currMap)

	var ops []Operation
	ops = append(ops, fkDropOps(tableName, prevNames, currMap)...)
	ops = append(ops, fkAlterOrAddOps(tableName, currNames, prevMap, currMap)...)
	return ops
}
