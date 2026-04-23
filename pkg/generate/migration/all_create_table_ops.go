//ff:func feature=migration type=util control=iteration dimension=1
//ff:what allCreateTableOps — ops 전부가 CreateTable 인지 순회 검사
package migration

// allCreateTableOps reports whether every op in ops is a CreateTable.
func allCreateTableOps(ops []Operation) bool {
	for _, op := range ops {
		if _, ok := op.(CreateTable); !ok {
			return false
		}
	}
	return true
}
