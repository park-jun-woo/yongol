//ff:func feature=migration type=util control=iteration dimension=1
//ff:what renamedColumnSets — 특정 테이블 범위의 rename 규칙에서 From/To 이름 집합 2개 반환
package migration

// renamedColumnSets returns two sets (From-columns, To-columns) filtered
// by table scope.
func renamedColumnSets(hints *Hints, tableName string) (map[string]bool, map[string]bool) {
	from := map[string]bool{}
	to := map[string]bool{}
	if hints == nil {
		return from, to
	}
	for _, r := range hints.RenameColumns {
		if r.Table == tableName {
			from[r.From] = true
			to[r.To] = true
		}
	}
	return from, to
}
