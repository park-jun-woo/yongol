//ff:func feature=migration type=util control=sequence
//ff:what ensureTable — Schema.Tables[name] 이 없으면 새로 만들고 반환
package migration

// ensureTable returns the existing Table for `name` or creates one.
func ensureTable(s *Schema, name string) *Table {
	if t, ok := s.Tables[name]; ok {
		return t
	}
	t := &Table{Name: name}
	s.Tables[name] = t
	return t
}
