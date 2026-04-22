//ff:func feature=migration type=util control=sequence
//ff:what diffIndexes — 인덱스 추가/삭제 diff (이름 기준, 컬럼셋 이름 변경은 drop+create)
package migration

import "sort"

func diffIndexes(prev, curr *Table, tableName string) []Operation {
	var ops []Operation
	prevMap := indexMap(prev.Indexes)
	currMap := indexMap(curr.Indexes)

	var prevNames, currNames []string
	for n := range prevMap {
		prevNames = append(prevNames, n)
	}
	for n := range currMap {
		currNames = append(currNames, n)
	}
	sort.Strings(prevNames)
	sort.Strings(currNames)

	for _, n := range prevNames {
		if _, ok := currMap[n]; ok {
			continue
		}
		ops = append(ops, DropIndex{Name: n})
	}
	for _, n := range currNames {
		if _, ok := prevMap[n]; ok {
			// Re-create if definition changed (columns / unique / where).
			if !indexEqual(prevMap[n], currMap[n]) {
				ops = append(ops, DropIndex{Name: n})
				ops = append(ops, CreateIndex{Table: tableName, Index: currMap[n]})
			}
			continue
		}
		ops = append(ops, CreateIndex{Table: tableName, Index: currMap[n]})
	}
	return ops
}

func indexMap(ix []*Index) map[string]*Index {
	m := make(map[string]*Index, len(ix))
	for _, i := range ix {
		m[i.Name] = i
	}
	return m
}

func indexEqual(a, b *Index) bool {
	if a == nil || b == nil {
		return false
	}
	if a.Unique != b.Unique || a.Where != b.Where || len(a.Columns) != len(b.Columns) {
		return false
	}
	for i := range a.Columns {
		if a.Columns[i] != b.Columns[i] {
			return false
		}
	}
	return true
}
