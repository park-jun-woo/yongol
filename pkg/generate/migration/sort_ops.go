//ff:func feature=migration type=util control=iteration dimension=1
//ff:what sortByDependency — 11단계 의존성 정렬 (DROP FK → DROP INDEX → … → ADD FK)
package migration

import "sort"

// sortByDependency places ops into the right execution order so FKs are
// torn down before referenced tables / columns change, and so FKs are
// re-attached only after their dependencies exist.
//
//   1. RenameTable / RenameColumn        (executed first so later ops use new names)
//   2. DropForeignKey
//   3. DropIndex
//   4. DropCheck
//   5. DropColumn
//   6. DropTable
//   7. CreateTable
//   8. AddColumn
//   9. AlterColumn* (Type / Nullable / Default)
//  10. AddCheck
//  11. CreateIndex
//  12. AddForeignKey
func sortByDependency(ops []Operation) []Operation {
	keyed := make([]keyedOp, len(ops))
	for i, op := range ops {
		keyed[i] = keyedOp{phase: phaseOf(op), order: i, op: op}
	}
	sort.SliceStable(keyed, func(i, j int) bool {
		if keyed[i].phase != keyed[j].phase {
			return keyed[i].phase < keyed[j].phase
		}
		return keyed[i].order < keyed[j].order
	})
	out := make([]Operation, len(ops))
	for i, k := range keyed {
		out[i] = k.op
	}
	return out
}
