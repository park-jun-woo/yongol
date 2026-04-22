//ff:func feature=migration type=util control=sequence
//ff:what diffColumns — 컬럼 추가/삭제/타입·Nullable·Default 변경 diff
package migration

import "sort"

func diffColumns(prev, curr *Table, hints *Hints, tableName string) []Operation {
	var ops []Operation
	prevMap := columnMap(prev.Columns)
	currMap := columnMap(curr.Columns)

	// Columns renamed within this table are handled as RenameColumn ops
	// emitted in diff.go already; remove them from add/drop consideration.
	renamedFrom := map[string]bool{}
	renamedTo := map[string]bool{}
	if hints != nil {
		for _, r := range hints.RenameColumns {
			if r.Table == tableName {
				renamedFrom[r.From] = true
				renamedTo[r.To] = true
			}
		}
	}

	// Deterministic iteration: sorted prev names first, then sorted curr
	// names (for adds).
	var prevNames, currNames []string
	for n := range prevMap {
		prevNames = append(prevNames, n)
	}
	for n := range currMap {
		currNames = append(currNames, n)
	}
	sort.Strings(prevNames)
	sort.Strings(currNames)

	// Drops — columns in prev but not in curr (and not consumed by rename).
	for _, n := range prevNames {
		if _, ok := currMap[n]; ok {
			continue
		}
		if renamedFrom[n] {
			continue
		}
		op := DropColumn{Table: tableName, Column: n}
		if hints != nil {
			// @allow_destructive is table-scoped.
			if hints.AllowDestructive[tableName] {
				op.AllowDestructive = true
			}
		}
		ops = append(ops, op)
	}

	// Adds — in curr but not in prev.
	for _, n := range currNames {
		if _, ok := prevMap[n]; ok {
			continue
		}
		if renamedTo[n] {
			// Rename already handled; still compare types/nullable/default.
			// The rename hint maps to an old column — find it and diff that.
			continue
		}
		col := currMap[n]
		op := AddColumn{Table: tableName, Column: col}
		if hints != nil {
			if b, ok := hints.Backfills[colKey{Table: tableName, Column: n}]; ok {
				op.Backfill = b
			}
		}
		ops = append(ops, op)
	}

	// Alters — present in both.
	for _, n := range currNames {
		cc, ok := currMap[n]
		if !ok {
			continue
		}
		var pc *Column
		if p, ok := prevMap[n]; ok {
			pc = p
		} else if hints != nil && renamedTo[n] {
			// Find prev column under old name.
			for _, r := range hints.RenameColumns {
				if r.Table == tableName && r.To == n {
					if p, ok2 := prevMap[r.From]; ok2 {
						pc = p
					}
					break
				}
			}
		}
		if pc == nil {
			continue
		}

		if !pc.Type.Equal(cc.Type) {
			aop := AlterColumnType{
				Table: tableName, Column: n,
				From: pc.Type, To: cc.Type,
			}
			if hints != nil {
				if using, ok := hints.Casts[colKey{Table: tableName, Column: n}]; ok {
					aop.Using = using
				}
			}
			ops = append(ops, aop)
		}
		if pc.Nullable != cc.Nullable {
			op := AlterColumnNullable{
				Table: tableName, Column: n,
				From: pc.Nullable, To: cc.Nullable,
			}
			if hints != nil {
				if b, ok := hints.Backfills[colKey{Table: tableName, Column: n}]; ok {
					op.Backfill = b
				}
			}
			ops = append(ops, op)
		}
		if pc.Default != cc.Default {
			ops = append(ops, AlterColumnDefault{
				Table: tableName, Column: n,
				From: pc.Default, To: cc.Default,
			})
		}
	}

	return ops
}

func columnMap(cols []*Column) map[string]*Column {
	m := make(map[string]*Column, len(cols))
	for _, c := range cols {
		m[c.Name] = c
	}
	return m
}
