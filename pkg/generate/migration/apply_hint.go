//ff:func feature=migration type=util control=selection
//ff:what applyHint — 단일 Operation 에 Hints 를 type switch 로 반영
package migration

// applyHint merges relevant hint values into op and returns the updated
// copy. Non-hint-aware op types are returned unchanged.
func applyHint(op Operation, hints *Hints) Operation {
	switch v := op.(type) {
	case DropTable:
		if hints.AllowDestructive[v.Name] {
			v.AllowDestructive = true
		}
		return v
	case DropColumn:
		if hints.AllowDestructive[v.Table] {
			v.AllowDestructive = true
		}
		return v
	case AddColumn:
		if b, ok := hints.Backfills[colKey{Table: v.Table, Column: v.Column.Name}]; ok {
			v.Backfill = b
		}
		return v
	case AlterColumnNullable:
		if b, ok := hints.Backfills[colKey{Table: v.Table, Column: v.Column}]; ok {
			v.Backfill = b
		}
		return v
	case AlterColumnType:
		if using, ok := hints.Casts[colKey{Table: v.Table, Column: v.Column}]; ok {
			v.Using = using
		}
		return v
	}
	return op
}
