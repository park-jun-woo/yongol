//ff:func feature=migration type=util control=iteration dimension=1 topic=migration-hints
//ff:what renamedColumnName — 컬럼명에 해당하는 rename 규칙 검색 (매칭 없으면 원본)
package migration

// renamedColumnName finds a matching RenameColumnHint and returns the
// target column name. If no rule matches, it returns colName unchanged.
func renamedColumnName(prevTable, newTable, colName string, rules []RenameColumnHint) string {
	for _, r := range rules {
		if (r.Table == prevTable || r.Table == newTable) && r.From == colName {
			return r.To
		}
	}
	return colName
}
