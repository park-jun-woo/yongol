//ff:func feature=migration type=util control=iteration dimension=1 topic=migration-hints
//ff:what renameColumnsOf — 한 테이블의 Columns 슬라이스에 @rename 규칙 적용
package migration

// renameColumnsOf returns a copy of t.Columns with matching column-rename
// hints applied. The rule matches either the old (pre-rename) or the new
// table name so users can phrase rules in either frame.
func renameColumnsOf(t *Table, newTableName string, rules []RenameColumnHint) []*Column {
	if len(rules) == 0 {
		return t.Columns
	}
	out := make([]*Column, len(t.Columns))
	for i, c := range t.Columns {
		cc := *c
		cc.Name = renamedColumnName(t.Name, newTableName, c.Name, rules)
		out[i] = &cc
	}
	return out
}
