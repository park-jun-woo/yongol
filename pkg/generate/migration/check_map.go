//ff:func feature=migration type=util control=iteration dimension=1
//ff:what checkMap — []*CheckConstraint 를 이름 기준 맵으로 변환
package migration

// checkMap returns a lookup map keyed by CHECK constraint name.
func checkMap(cs []*CheckConstraint) map[string]*CheckConstraint {
	m := make(map[string]*CheckConstraint, len(cs))
	for _, c := range cs {
		m[c.Name] = c
	}
	return m
}
