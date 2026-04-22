//ff:func feature=migration type=util control=sequence
//ff:what Diff — prev/curr Schema AST 를 Operation 리스트로 변환 (hints 지원)
package migration

import "sort"

// Diff compares two schemas and returns the minimal ordered list of
// Operations that transform prev into curr. hints may be nil (Phase003
// standalone). Ordering is deterministic — callers can rely on stable
// output for the same inputs.
func Diff(prev, curr *Schema, hints *Hints) []Operation {
	if prev == nil {
		prev = NewSchema()
	}
	if curr == nil {
		curr = NewSchema()
	}

	// Apply rename hints by rewriting prev so diff matches renamed items.
	prev2 := applyRenameHints(prev, hints)

	var ops []Operation

	// Collect rename-generated operations first (they must appear in SQL
	// order 1 — before column adds / drops).
	ops = append(ops, collectRenameOps(hints)...)

	// Table-level diff.
	allTables := map[string]bool{}
	for n := range prev2.Tables {
		allTables[n] = true
	}
	for n := range curr.Tables {
		allTables[n] = true
	}
	names := make([]string, 0, len(allTables))
	for n := range allTables {
		names = append(names, n)
	}
	sort.Strings(names)

	// Skip tables involved in rename so we don't emit drop+create.
	renamed := map[string]string{} // prev name -> curr name
	renamedRev := map[string]string{}
	if hints != nil {
		for _, r := range hints.RenameTables {
			renamed[r.From] = r.To
			renamedRev[r.To] = r.From
		}
	}

	for _, n := range names {
		p, pok := prev2.Tables[n]
		c, cok := curr.Tables[n]
		switch {
		case !pok && cok:
			if _, isRenameTarget := renamedRev[n]; isRenameTarget {
				// Handled by rename; still need to diff columns if any differ.
				from := renamedRev[n]
				prevT := prev.Tables[from]
				ops = append(ops, diffTableBody(prevT, c, hints, n)...)
				continue
			}
			ops = append(ops, CreateTable{Table: c})
			for _, idx := range c.Indexes {
				ops = append(ops, CreateIndex{Table: c.Name, Index: idx})
			}
			for _, fk := range c.ForeignKeys {
				ops = append(ops, AddForeignKey{Table: c.Name, FK: fk})
			}
			for _, chk := range c.Checks {
				ops = append(ops, AddCheck{Table: c.Name, Check: chk})
			}
		case pok && !cok:
			if _, isRenameSource := renamed[n]; isRenameSource {
				continue // the rename op already handles it
			}
			// Emit FK drops before table drop.
			for _, fk := range p.ForeignKeys {
				ops = append(ops, DropForeignKey{Table: p.Name, Name: fk.Name})
			}
			ops = append(ops, DropTable{Name: n})
		case pok && cok:
			ops = append(ops, diffTableBody(p, c, hints, n)...)
		}
	}

	return sortByDependency(ops)
}

// collectRenameOps turns rename hints into explicit RenameTable /
// RenameColumn operations. Called before diff so body-level diffs see
// the post-rename shape.
func collectRenameOps(hints *Hints) []Operation {
	if hints == nil {
		return nil
	}
	var ops []Operation
	for _, r := range hints.RenameTables {
		ops = append(ops, RenameTable{From: r.From, To: r.To})
	}
	for _, r := range hints.RenameColumns {
		ops = append(ops, RenameColumn{Table: r.Table, From: r.From, To: r.To})
	}
	return ops
}
