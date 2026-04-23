//ff:func feature=migration type=util control=sequence
//ff:what indexDiffForOne — 한 이름의 Index 가 추가/변경됐는지 판정해 ops 반환
package migration

// indexDiffForOne returns Create / (Drop+Create) ops for a single index.
func indexDiffForOne(tableName, n string, prevMap, currMap map[string]*Index) []Operation {
	_, ok := prevMap[n]
	if !ok {
		return []Operation{CreateIndex{Table: tableName, Index: currMap[n]}}
	}
	if indexEqual(prevMap[n], currMap[n]) {
		return nil
	}
	return []Operation{
		DropIndex{Name: n},
		CreateIndex{Table: tableName, Index: currMap[n]},
	}
}
