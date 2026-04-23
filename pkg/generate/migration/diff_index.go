//ff:func feature=migration type=util control=sequence
//ff:what diffIndexes — 인덱스 추가/삭제/재생성 diff
package migration

// diffIndexes compares the indexes of prev and curr and emits
// DropIndex / CreateIndex as needed.
func diffIndexes(prev, curr *Table, tableName string) []Operation {
	prevMap := indexMap(prev.Indexes)
	currMap := indexMap(curr.Indexes)
	prevNames := sortedMapKeys(prevMap)
	currNames := sortedMapKeys(currMap)

	var ops []Operation
	ops = append(ops, indexDropOps(prevNames, currMap)...)
	ops = append(ops, indexAlterOrCreateOps(tableName, currNames, prevMap, currMap)...)
	return ops
}
