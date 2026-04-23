//ff:func feature=migration type=util control=iteration dimension=1
//ff:what columnDropOps — prev 에만 있고 rename 되지 않은 컬럼에 대한 DropColumn 연산 리스트
package migration

// columnDropOps emits DropColumn for each prev column not present in
// curr and not consumed by a rename hint.
func columnDropOps(tableName string, prevNames []string, currMap map[string]*Column, renamedFrom map[string]bool, hints *Hints) []Operation {
	var ops []Operation
	for _, n := range prevNames {
		if _, ok := currMap[n]; ok {
			continue
		}
		if renamedFrom[n] {
			continue
		}
		op := DropColumn{Table: tableName, Column: n}
		if hints != nil && hints.AllowDestructive[tableName] {
			op.AllowDestructive = true
		}
		ops = append(ops, op)
	}
	return ops
}
