//ff:func feature=gen-ir type=util control=iteration dimension=1
//ff:what pkColumnSet -- primary key 컬럼 슬라이스를 lookup 집합으로 변환

package ir

// pkColumnSet returns a set of the given primary key column names for O(1)
// membership lookup.
func pkColumnSet(primaryKey []string) map[string]bool {
	set := make(map[string]bool, len(primaryKey))
	for _, pk := range primaryKey {
		set[pk] = true
	}
	return set
}
