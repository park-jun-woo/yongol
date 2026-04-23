//ff:func feature=migration type=util control=iteration dimension=1
//ff:what columnAddOps — curr 에만 있고 rename 된 컬럼이 아닌 경우 AddColumn 연산 생성
package migration

// columnAddOps emits AddColumn for each curr column that is not in prev
// and not the target of a column rename hint.
func columnAddOps(tableName string, currNames []string, prevMap map[string]*Column, renamedTo map[string]bool, currMap map[string]*Column, hints *Hints) []Operation {
	var ops []Operation
	for _, n := range currNames {
		if _, ok := prevMap[n]; ok {
			continue
		}
		if renamedTo[n] {
			continue
		}
		ops = append(ops, buildAddColumnOp(tableName, n, currMap[n], hints))
	}
	return ops
}
