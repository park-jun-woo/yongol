//ff:func feature=migration type=util control=iteration dimension=1
//ff:what indexMap — []*Index 를 이름 기준 맵으로 변환
package migration

// indexMap returns a lookup map keyed by index name.
func indexMap(ix []*Index) map[string]*Index {
	m := make(map[string]*Index, len(ix))
	for _, i := range ix {
		m[i.Name] = i
	}
	return m
}
