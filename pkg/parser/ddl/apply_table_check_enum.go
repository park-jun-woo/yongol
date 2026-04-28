//ff:func feature=manifest type=util control=sequence
//ff:what applyTableCheckEnum — 테이블 레벨 CHECK 절을 해당 Column에 적용
package ddl

// applyTableCheckEnum handles a table-level CHECK constraint line. The
// column name is parsed from the CHECK clause and the enum values are
// attached to that Column entry on Table. No-op when the column is not
// in the table or the enum is empty.
func applyTableCheckEnum(line string, t *Table) {
	colName, vals := parseCheckEnum(line)
	if colName == "" || len(vals) == 0 {
		return
	}
	col, ok := t.Columns[colName]
	if !ok {
		return
	}
	col.CheckEnum = vals
	t.Columns[colName] = col
}
