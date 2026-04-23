//ff:func feature=migration type=util control=sequence
//ff:what buildAlterColumnType — Casts 힌트 적용한 AlterColumnType Operation 생성
package migration

// buildAlterColumnType constructs an AlterColumnType, applying any
// matching @cast USING hint.
func buildAlterColumnType(tableName, column string, from, to CanonicalType, hints *Hints) AlterColumnType {
	op := AlterColumnType{Table: tableName, Column: column, From: from, To: to}
	if hints == nil {
		return op
	}
	if using, ok := hints.Casts[colKey{Table: tableName, Column: column}]; ok {
		op.Using = using
	}
	return op
}
