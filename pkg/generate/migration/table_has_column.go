//ff:func feature=migration type=util control=iteration dimension=1
//ff:what tableHasColumn — Table.Columns 에서 주어진 이름 존재 여부 선형 탐색
package migration

// tableHasColumn reports whether t has a column named `name`.
func tableHasColumn(t *Table, name string) bool {
	for _, c := range t.Columns {
		if c.Name == name {
			return true
		}
	}
	return false
}
