//ff:func feature=migration type=parser control=sequence
//ff:what ParseHints — pkg/parser/ddl 의 HintComment 로부터 Hints 구조체 구성
package migration

import (
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// ParseHints converts raw hint comments extracted from DDL files into
// the Hints structure consumed by Diff / check_safety. An empty list
// yields a non-nil but empty Hints.
func ParseHints(comments []ddl.HintComment) *Hints {
	h := &Hints{
		Casts:            map[colKey]string{},
		Backfills:        map[colKey]string{},
		DataMigrations:   map[string]string{},
		AllowDestructive: map[string]bool{},
	}
	for _, c := range comments {
		switch c.Tag {
		case "rename":
			from := c.Args["from"]
			to := c.Args["to"]
			switch {
			case c.ColumnCtx != "" && from != "":
				// Column rename (attached to column line). 'to' defaults to column name.
				toName := to
				if toName == "" {
					toName = c.ColumnCtx
				}
				h.RenameColumns = append(h.RenameColumns, RenameColumnHint{
					Table: c.TableCtx, From: from, To: toName,
				})
			case c.BlockAbove && from != "" && to != "":
				// Above CREATE TABLE — table rename.
				h.RenameTables = append(h.RenameTables, RenameTableHint{From: from, To: to})
			case from != "" && to != "" && c.TableCtx != "":
				// Inline same-line rename: treat as column rename if the 'to'
				// matches a column.
				h.RenameColumns = append(h.RenameColumns, RenameColumnHint{
					Table: c.TableCtx, From: from, To: to,
				})
			}
		case "cast":
			if using := c.Args["using"]; using != "" && c.ColumnCtx != "" {
				h.Casts[colKey{Table: c.TableCtx, Column: c.ColumnCtx}] = using
			}
		case "backfill":
			if def := c.Args["default"]; def != "" && c.ColumnCtx != "" {
				h.Backfills[colKey{Table: c.TableCtx, Column: c.ColumnCtx}] = def
			}
		case "data_migration":
			if file := c.Args["file"]; file != "" && c.TableCtx != "" {
				h.DataMigrations[c.TableCtx] = file
			}
		case "allow_destructive":
			if c.TableCtx != "" {
				h.AllowDestructive[c.TableCtx] = true
			}
		}
	}
	return h
}
