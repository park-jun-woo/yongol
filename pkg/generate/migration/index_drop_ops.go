//ff:func feature=migration type=util control=iteration dimension=1
//ff:what indexDropOps — prev 에만 있는 인덱스 이름에 대해 DropIndex 연산 생성
package migration

// indexDropOps emits DropIndex for each index name in prev but absent
// from curr.
func indexDropOps(prevNames []string, currMap map[string]*Index) []Operation {
	var ops []Operation
	for _, n := range prevNames {
		if _, ok := currMap[n]; ok {
			continue
		}
		ops = append(ops, DropIndex{Name: n})
	}
	return ops
}
