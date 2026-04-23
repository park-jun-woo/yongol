//ff:func feature=migration type=util control=sequence
//ff:what buildAddColumnOp — AddColumn Operation 생성 + Backfills 힌트 반영
package migration

// buildAddColumnOp constructs an AddColumn, applying any matching
// @backfill hint.
func buildAddColumnOp(tableName, column string, col *Column, hints *Hints) AddColumn {
	op := AddColumn{Table: tableName, Column: col}
	if hints == nil {
		return op
	}
	if b, ok := hints.Backfills[colKey{Table: tableName, Column: column}]; ok {
		op.Backfill = b
	}
	return op
}
