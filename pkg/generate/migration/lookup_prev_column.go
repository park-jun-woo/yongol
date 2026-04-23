//ff:func feature=migration type=util control=sequence
//ff:what lookupPrevColumn — 현재 컬럼명에 대응하는 prev 컬럼 탐색 (rename 고려)
package migration

// lookupPrevColumn returns the prev-side *Column that matches current
// name n, following a column-rename hint if n is a rename target.
func lookupPrevColumn(n string, prevMap map[string]*Column, renamedTo map[string]bool, hints *Hints, tableName string) *Column {
	if p, ok := prevMap[n]; ok {
		return p
	}
	if hints == nil || !renamedTo[n] {
		return nil
	}
	return findPrevViaRenameHint(n, prevMap, hints.RenameColumns, tableName)
}
