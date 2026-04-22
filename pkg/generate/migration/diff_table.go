//ff:func feature=migration type=util control=sequence
//ff:what diffTableBody — 한 테이블의 컬럼/인덱스/FK/CHECK 차이를 Operation 으로 변환
package migration

// diffTableBody compares one pair of existing Tables. currName is the
// *current* canonical name (used as Operation.Table) so rename targets
// still emit well-formed ALTER statements.
func diffTableBody(prev, curr *Table, hints *Hints, currName string) []Operation {
	var ops []Operation
	if prev == nil || curr == nil {
		return nil
	}
	ops = append(ops, diffColumns(prev, curr, hints, currName)...)
	ops = append(ops, diffIndexes(prev, curr, currName)...)
	ops = append(ops, diffForeignKeys(prev, curr, currName)...)
	ops = append(ops, diffChecks(prev, curr, currName)...)
	return ops
}
