//ff:func feature=migration type=parser control=sequence topic=migration-hints
//ff:what applyBackfillHint — @backfill default=... 코멘트를 Hints.Backfills 에 등록
package migration

import "github.com/park-jun-woo/yongol/pkg/parser/ddl"

// applyBackfillHint stores a column-level backfill default literal.
func applyBackfillHint(h *Hints, c ddl.HintComment) {
	def := c.Args["default"]
	if def == "" || c.ColumnCtx == "" {
		return
	}
	h.Backfills[colKey{Table: c.TableCtx, Column: c.ColumnCtx}] = def
}
