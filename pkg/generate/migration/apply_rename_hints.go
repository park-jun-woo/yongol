//ff:func feature=migration type=model control=sequence topic=migration-hints
//ff:what Hints 구조체 + applyRenameHints — 실제 rename 규칙 적용은 Diff 호출부
package migration

// colKey is the map key for (table,column) scoped hints.
type colKey struct {
	Table, Column string
}

// Hints carries DDL comment hints consumed by Diff / safety checks.
// Phase003 leaves ParseHints unimplemented; value is set in Phase004.
type Hints struct {
	RenameTables     []RenameTableHint
	RenameColumns    []RenameColumnHint
	Casts            map[colKey]string // (table,col) → USING expr
	Backfills        map[colKey]string // (table,col) → literal default
	DataMigrations   map[string]string // table → sidecar file path
	AllowDestructive map[string]bool   // table → allow DROP TABLE/COLUMN
}

// RenameTableHint maps an old table name to a new one.
type RenameTableHint struct{ From, To string }

// RenameColumnHint maps an old column name to a new one inside a table.
type RenameColumnHint struct{ Table, From, To string }

// applyRenameHints returns a shallow copy of prev with tables/columns
// renamed so Diff can line them up with curr. If hints is nil, returns
// prev unchanged.
func applyRenameHints(prev *Schema, hints *Hints) *Schema {
	if prev == nil || hints == nil {
		return prev
	}
	if len(hints.RenameTables) == 0 && len(hints.RenameColumns) == 0 {
		return prev
	}
	out := &Schema{Tables: make(map[string]*Table, len(prev.Tables))}
	for name, t := range prev.Tables {
		// Table rename: emit under new name.
		newName := name
		for _, r := range hints.RenameTables {
			if r.From == name {
				newName = r.To
				break
			}
		}
		copy := *t
		copy.Name = newName
		copy.Columns = renameColumnsOf(t, newName, hints.RenameColumns)
		out.Tables[newName] = &copy
	}
	return out
}

func renameColumnsOf(t *Table, newTableName string, rules []RenameColumnHint) []*Column {
	if len(rules) == 0 {
		return t.Columns
	}
	out := make([]*Column, len(t.Columns))
	for i, c := range t.Columns {
		cc := *c
		for _, r := range rules {
			// The rule matches either the old (pre-rename) name or the new
			// table name, since some users phrase rules in either frame.
			if (r.Table == t.Name || r.Table == newTableName) && r.From == c.Name {
				cc.Name = r.To
				break
			}
		}
		out[i] = &cc
	}
	return out
}
