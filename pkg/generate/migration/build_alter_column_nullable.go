//ff:func feature=migration type=util control=sequence
//ff:what buildAlterColumnNullable — Backfills 힌트 적용한 AlterColumnNullable Operation 생성
package migration

// buildAlterColumnNullable constructs an AlterColumnNullable, applying
// any matching @backfill hint.
func buildAlterColumnNullable(tableName, column string, from, to bool, hints *Hints) AlterColumnNullable {
	op := AlterColumnNullable{Table: tableName, Column: column, From: from, To: to}
	if hints == nil {
		return op
	}
	if b, ok := hints.Backfills[colKey{Table: tableName, Column: column}]; ok {
		op.Backfill = b
	}
	return op
}
