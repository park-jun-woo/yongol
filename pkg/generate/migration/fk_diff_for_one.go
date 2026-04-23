//ff:func feature=migration type=util control=sequence
//ff:what fkDiffForOne — 한 이름의 FK 가 추가/변경됐는지 판정해 ops 반환
package migration

// fkDiffForOne returns Add / (Drop+Add) ops for a single FK name.
func fkDiffForOne(tableName, n string, prevMap, currMap map[string]*ForeignKey) []Operation {
	p, ok := prevMap[n]
	if !ok {
		return []Operation{AddForeignKey{Table: tableName, FK: currMap[n]}}
	}
	if fkEqual(p, currMap[n]) {
		return nil
	}
	return []Operation{
		DropForeignKey{Table: tableName, Name: n},
		AddForeignKey{Table: tableName, FK: currMap[n]},
	}
}
