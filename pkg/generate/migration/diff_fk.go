//ff:func feature=migration type=util control=sequence
//ff:what diffForeignKeys / diffChecks — FK·CHECK 추가/삭제 diff
package migration

import "sort"

func diffForeignKeys(prev, curr *Table, tableName string) []Operation {
	var ops []Operation
	prevMap := fkMap(prev.ForeignKeys)
	currMap := fkMap(curr.ForeignKeys)

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
		ops = append(ops, DropForeignKey{Table: tableName, Name: n})
	}
	for _, n := range currNames {
		if p, ok := prevMap[n]; ok {
			if !fkEqual(p, currMap[n]) {
				ops = append(ops, DropForeignKey{Table: tableName, Name: n})
				ops = append(ops, AddForeignKey{Table: tableName, FK: currMap[n]})
			}
			continue
		}
		ops = append(ops, AddForeignKey{Table: tableName, FK: currMap[n]})
	}
	return ops
}

func fkMap(fks []*ForeignKey) map[string]*ForeignKey {
	m := make(map[string]*ForeignKey, len(fks))
	for _, fk := range fks {
		m[fk.Name] = fk
	}
	return m
}

func fkEqual(a, b *ForeignKey) bool {
	if a == nil || b == nil {
		return false
	}
	if a.RefTable != b.RefTable || a.OnDelete != b.OnDelete || a.OnUpdate != b.OnUpdate {
		return false
	}
	if len(a.Columns) != len(b.Columns) || len(a.RefColumns) != len(b.RefColumns) {
		return false
	}
	for i := range a.Columns {
		if a.Columns[i] != b.Columns[i] {
			return false
		}
	}
	for i := range a.RefColumns {
		if a.RefColumns[i] != b.RefColumns[i] {
			return false
		}
	}
	return true
}

func diffChecks(prev, curr *Table, tableName string) []Operation {
	var ops []Operation
	prevMap := checkMap(prev.Checks)
	currMap := checkMap(curr.Checks)

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
		ops = append(ops, DropCheck{Table: tableName, Name: n})
	}
	for _, n := range currNames {
		if p, ok := prevMap[n]; ok {
			if p.Expression != currMap[n].Expression {
				ops = append(ops, DropCheck{Table: tableName, Name: n})
				ops = append(ops, AddCheck{Table: tableName, Check: currMap[n]})
			}
			continue
		}
		ops = append(ops, AddCheck{Table: tableName, Check: currMap[n]})
	}
	return ops
}

func checkMap(cs []*CheckConstraint) map[string]*CheckConstraint {
	m := make(map[string]*CheckConstraint, len(cs))
	for _, c := range cs {
		m[c.Name] = c
	}
	return m
}
