//ff:func feature=migration type=util control=iteration dimension=1
//ff:what findPrevViaRenameHint — rename 규칙에서 To==n 찾아 From 이름으로 prevMap 조회
package migration

// findPrevViaRenameHint scans rules for a column rename that targets n
// in tableName, and returns the prev column for the matched From name.
func findPrevViaRenameHint(n string, prevMap map[string]*Column, rules []RenameColumnHint, tableName string) *Column {
	for _, r := range rules {
		if r.Table != tableName || r.To != n {
			continue
		}
		if p, ok := prevMap[r.From]; ok {
			return p
		}
		return nil
	}
	return nil
}
