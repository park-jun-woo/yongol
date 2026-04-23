//ff:func feature=migration type=util control=iteration dimension=1
//ff:what fkMap — []*ForeignKey 를 이름 기준 맵으로 변환
package migration

// fkMap returns a lookup map keyed by FK name.
func fkMap(fks []*ForeignKey) map[string]*ForeignKey {
	m := make(map[string]*ForeignKey, len(fks))
	for _, fk := range fks {
		m[fk.Name] = fk
	}
	return m
}
