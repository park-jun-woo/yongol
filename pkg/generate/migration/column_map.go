//ff:func feature=migration type=util control=iteration dimension=1
//ff:what columnMap — []*Column 을 이름 기준 맵으로 변환
package migration

// columnMap returns a lookup map keyed by column name.
func columnMap(cols []*Column) map[string]*Column {
	m := make(map[string]*Column, len(cols))
	for _, c := range cols {
		m[c.Name] = c
	}
	return m
}
