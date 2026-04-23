//ff:func feature=migration type=util control=sequence
//ff:what diffAddOrRenameTarget — curr 에만 있는 테이블: rename 대상이면 body diff, 아니면 CreateTable + 의존물
package migration

// diffAddOrRenameTarget emits CreateTable for genuinely new tables or
// a body-level diff for rename targets (since the RenameTable op was
// already emitted in Diff).
func diffAddOrRenameTarget(n string, prev *Schema, c *Table, hints *Hints, renamedRev map[string]string) []Operation {
	if from, isRenameTarget := renamedRev[n]; isRenameTarget {
		prevT := prev.Tables[from]
		return diffTableBody(prevT, c, hints, n)
	}
	return createTableWithDeps(c)
}
