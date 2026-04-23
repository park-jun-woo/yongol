//ff:func feature=migration type=util control=sequence
//ff:what NewSchema — 빈 Schema (Tables 맵 초기화된) 반환
package migration

// NewSchema returns an empty Schema with a non-nil Tables map.
func NewSchema() *Schema {
	return &Schema{Tables: map[string]*Table{}}
}
