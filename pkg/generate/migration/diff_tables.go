//ff:func feature=migration type=util control=iteration dimension=1
//ff:what diffTables — 테이블 단위 diff 순회 (신규 CREATE / 삭제 DROP / 변경 diffTableBody)
package migration

import "sort"

// diffTables walks all table names present in either prev2 or curr and
// emits the appropriate ops per table. `prev` is the original schema
// (pre-rename) used to look up columns of the rename source.
func diffTables(prev, prev2, curr *Schema, hints *Hints) []Operation {
	names := collectAllTableNames(prev2, curr)
	sort.Strings(names)

	renamed, renamedRev := renameMaps(hints)

	var ops []Operation
	for _, n := range names {
		ops = append(ops, diffOneTable(n, prev, prev2, curr, hints, renamed, renamedRev)...)
	}
	return ops
}
