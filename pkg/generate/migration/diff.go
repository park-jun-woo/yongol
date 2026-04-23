//ff:func feature=migration type=util control=sequence
//ff:what Diff — prev/curr Schema AST 를 Operation 리스트로 변환 (hints 지원)
package migration

// Diff compares two schemas and returns the minimal ordered list of
// Operations that transform prev into curr. hints may be nil. Ordering
// is deterministic — callers can rely on stable output for the same
// inputs.
func Diff(prev, curr *Schema, hints *Hints) []Operation {
	if prev == nil {
		prev = NewSchema()
	}
	if curr == nil {
		curr = NewSchema()
	}
	prev2 := applyRenameHints(prev, hints)

	var ops []Operation
	ops = append(ops, collectRenameOps(hints)...)
	ops = append(ops, diffTables(prev, prev2, curr, hints)...)
	return sortByDependency(ops)
}
