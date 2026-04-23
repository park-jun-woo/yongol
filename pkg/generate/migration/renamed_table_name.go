//ff:func feature=migration type=util control=iteration dimension=1 topic=migration-hints
//ff:what renamedTableName — 기존 테이블명을 rename 규칙에 따라 새 이름으로 매핑 (해당 없으면 원본)
package migration

// renamedTableName returns the new table name when a rule matches `name`,
// otherwise returns `name` unchanged.
func renamedTableName(name string, rules []RenameTableHint) string {
	for _, r := range rules {
		if r.From == name {
			return r.To
		}
	}
	return name
}
