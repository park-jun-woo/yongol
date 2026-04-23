//ff:func feature=migration type=util control=selection
//ff:what diffOneTable — 한 테이블 이름에 대한 diff 분기 (신규/삭제/rename target/양쪽 존재)
package migration

// diffOneTable handles the cases for a single table name `n`:
//   - exists only in curr → CreateTable (or rename target)
//   - exists only in prev2 → DropTable (or rename source — already handled)
//   - exists in both      → body diff
func diffOneTable(n string, prev, prev2, curr *Schema, hints *Hints, renamed, renamedRev map[string]string) []Operation {
	p, pok := prev2.Tables[n]
	c, cok := curr.Tables[n]
	switch {
	case !pok && cok:
		return diffAddOrRenameTarget(n, prev, c, hints, renamedRev)
	case pok && !cok:
		return diffDropTable(n, p, renamed)
	case pok && cok:
		return diffTableBody(p, c, hints, n)
	}
	return nil
}
