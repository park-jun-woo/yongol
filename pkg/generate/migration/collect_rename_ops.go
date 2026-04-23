//ff:func feature=migration type=util control=iteration dimension=1
//ff:what collectRenameOps — rename 힌트를 RenameTable/RenameColumn Operation 으로 변환
package migration

// collectRenameOps turns rename hints into explicit RenameTable /
// RenameColumn operations. Called before diff so body-level diffs see
// the post-rename shape.
func collectRenameOps(hints *Hints) []Operation {
	if hints == nil {
		return nil
	}
	var ops []Operation
	for _, r := range hints.RenameTables {
		ops = append(ops, RenameTable{From: r.From, To: r.To})
	}
	for _, r := range hints.RenameColumns {
		ops = append(ops, RenameColumn{Table: r.Table, From: r.From, To: r.To})
	}
	return ops
}
