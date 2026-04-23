//ff:func feature=migration type=util control=iteration dimension=1
//ff:what checkDropOps — prev 에만 있는 CHECK 이름에 대해 DropCheck 연산 생성
package migration

// checkDropOps emits DropCheck for each CHECK name present in prev but
// absent in curr.
func checkDropOps(tableName string, prevNames []string, currMap map[string]*CheckConstraint) []Operation {
	var ops []Operation
	for _, n := range prevNames {
		if _, ok := currMap[n]; ok {
			continue
		}
		ops = append(ops, DropCheck{Table: tableName, Name: n})
	}
	return ops
}
