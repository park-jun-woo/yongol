//ff:func feature=migration type=util control=iteration dimension=1
//ff:what columnAlterOps — prev/curr 양쪽에 있는 컬럼의 Type/Nullable/Default 변경 Alter 연산 생성
package migration

// columnAlterOps emits ALTER ops for columns present in both prev and
// curr when their Type / Nullable / Default differs.
func columnAlterOps(tableName string, currNames []string, prevMap, currMap map[string]*Column, renamedTo map[string]bool, hints *Hints) []Operation {
	var ops []Operation
	for _, n := range currNames {
		cc, ok := currMap[n]
		if !ok {
			continue
		}
		pc := lookupPrevColumn(n, prevMap, renamedTo, hints, tableName)
		if pc == nil {
			continue
		}
		ops = append(ops, columnAlterForPair(tableName, n, pc, cc, hints)...)
	}
	return ops
}
