//ff:func feature=migration type=util control=iteration dimension=1 topic=migration-hints
//ff:what applyRenameHints — prev Schema 를 rename 규칙 적용한 얕은 복사로 변환
package migration

// applyRenameHints returns a shallow copy of prev with tables/columns
// renamed so Diff can line them up with curr. If hints is nil, returns
// prev unchanged.
func applyRenameHints(prev *Schema, hints *Hints) *Schema {
	if prev == nil || hints == nil {
		return prev
	}
	if len(hints.RenameTables) == 0 && len(hints.RenameColumns) == 0 {
		return prev
	}
	out := &Schema{Tables: make(map[string]*Table, len(prev.Tables))}
	for name, t := range prev.Tables {
		newName := renamedTableName(name, hints.RenameTables)
		cp := *t
		cp.Name = newName
		cp.Columns = renameColumnsOf(t, newName, hints.RenameColumns)
		out.Tables[newName] = &cp
	}
	return out
}
