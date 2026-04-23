//ff:func feature=migration type=parser control=selection topic=migration-hints
//ff:what applyRenameHint — @rename 코멘트를 RenameColumnHint/RenameTableHint 로 분류
package migration

import "github.com/park-jun-woo/yongol/pkg/parser/ddl"

// applyRenameHint interprets a @rename comment as either a column or a
// table rename depending on context (column-attached vs block-above).
func applyRenameHint(h *Hints, c ddl.HintComment) {
	from := c.Args["from"]
	to := c.Args["to"]
	switch {
	case c.ColumnCtx != "" && from != "":
		toName := to
		if toName == "" {
			toName = c.ColumnCtx
		}
		h.RenameColumns = append(h.RenameColumns, RenameColumnHint{
			Table: c.TableCtx, From: from, To: toName,
		})
	case c.BlockAbove && from != "" && to != "":
		h.RenameTables = append(h.RenameTables, RenameTableHint{From: from, To: to})
	case from != "" && to != "" && c.TableCtx != "":
		h.RenameColumns = append(h.RenameColumns, RenameColumnHint{
			Table: c.TableCtx, From: from, To: to,
		})
	}
}
