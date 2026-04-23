//ff:func feature=migration type=util control=iteration dimension=1
//ff:what fkDropOps — prev 에만 있는 FK 이름에 대해 DropForeignKey 연산 생성
package migration

// fkDropOps emits DropForeignKey for each FK in prev that is not in curr.
func fkDropOps(tableName string, prevNames []string, currMap map[string]*ForeignKey) []Operation {
	var ops []Operation
	for _, n := range prevNames {
		if _, ok := currMap[n]; ok {
			continue
		}
		ops = append(ops, DropForeignKey{Table: tableName, Name: n})
	}
	return ops
}
