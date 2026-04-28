//ff:func feature=migration type=test-helper control=iteration dimension=1
//ff:what findIndexByName — 테스트용: 주어진 이름의 Index 를 배열에서 찾는다

package migration

// findIndexByName returns the first index in indexes matching name, or nil
// when none matches. Used by BuildASTFromSQL USING-method tests to pull a
// specific index out of the parsed table.
func findIndexByName(indexes []*Index, name string) *Index {
	for _, idx := range indexes {
		if idx.Name == name {
			return idx
		}
	}
	return nil
}
