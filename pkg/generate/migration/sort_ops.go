//ff:func feature=migration type=util control=sequence
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
	// Use stable sort with a phase-based key. Ties are broken by
	// description (also deterministic).
	keyed := make([]keyedOp, len(ops))
	for i, op := range ops {
		keyed[i] = keyedOp{phase: phaseOf(op), order: i, op: op}
	}
	sort.SliceStable(keyed, func(i, j int) bool {
		if keyed[i].phase != keyed[j].phase {
			return keyed[i].phase < keyed[j].phase
		}
		// Within the same phase preserve original order for determinism.
		return keyed[i].order < keyed[j].order
	})
	out := make([]Operation, len(ops))
	for i, k := range keyed {
		out[i] = k.op
	}
	return out
}

type keyedOp struct {
	phase int
	order int
	op    Operation
}

func phaseOf(op Operation) int {
	switch op.(type) {
	case RenameTable, RenameColumn:
		return 1
	case DropForeignKey:
		return 2
	case DropIndex:
		return 3
	case DropCheck:
		return 4
	case DropColumn:
		return 5
	case DropTable:
		return 6
	case CreateTable:
		return 7
	case AddColumn:
		return 8
	case AlterColumnType, AlterColumnNullable, AlterColumnDefault:
		return 9
	case AddCheck:
		return 10
	case CreateIndex:
		return 11
	case AddForeignKey:
		return 12
	}
	return 99
}
