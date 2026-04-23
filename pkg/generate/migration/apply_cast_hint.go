//ff:func feature=migration type=parser control=sequence topic=migration-hints
//ff:what applyCastHint — @cast using=... 코멘트를 Hints.Casts 에 등록
package migration

import "github.com/park-jun-woo/yongol/pkg/parser/ddl"

// applyCastHint stores a column-level USING expression.
func applyCastHint(h *Hints, c ddl.HintComment) {
	using := c.Args["using"]
	if using == "" || c.ColumnCtx == "" {
		return
	}
	h.Casts[colKey{Table: c.TableCtx, Column: c.ColumnCtx}] = using
}
