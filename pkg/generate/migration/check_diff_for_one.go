//ff:func feature=migration type=util control=sequence
//ff:what checkDiffForOne — 한 이름의 CHECK 가 추가/변경됐는지 판정해 ops 반환
package migration

// checkDiffForOne returns Add / (Drop+Add) ops for the CHECK with the
// given name.
func checkDiffForOne(tableName, n string, prevMap, currMap map[string]*CheckConstraint) []Operation {
	p, ok := prevMap[n]
	if !ok {
		return []Operation{AddCheck{Table: tableName, Check: currMap[n]}}
	}
	if p.Expression == currMap[n].Expression {
		return nil
	}
	return []Operation{
		DropCheck{Table: tableName, Name: n},
		AddCheck{Table: tableName, Check: currMap[n]},
	}
}
