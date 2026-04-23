//ff:func feature=migration type=util control=sequence
//ff:what diffColumns — 한 테이블의 컬럼 diff (Drop → Add → Alter 순서)
package migration

// diffColumns compares the columns of prev vs curr inside the same
// (post-rename) table and emits DropColumn / AddColumn / AlterColumn*
// operations. Columns consumed by @rename hints are skipped.
func diffColumns(prev, curr *Table, hints *Hints, tableName string) []Operation {
	prevMap := columnMap(prev.Columns)
	currMap := columnMap(curr.Columns)
	renamedFrom, renamedTo := renamedColumnSets(hints, tableName)

	prevNames := sortedMapKeys(prevMap)
	currNames := sortedMapKeys(currMap)

	var ops []Operation
	ops = append(ops, columnDropOps(tableName, prevNames, currMap, renamedFrom, hints)...)
	ops = append(ops, columnAddOps(tableName, currNames, prevMap, renamedTo, currMap, hints)...)
	ops = append(ops, columnAlterOps(tableName, currNames, prevMap, currMap, renamedTo, hints)...)
	return ops
}
