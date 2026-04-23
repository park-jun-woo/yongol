//ff:func feature=migration type=parser control=sequence topic=migration-hints
//ff:what applyAllowDestructiveHint — @allow_destructive 코멘트를 Hints.AllowDestructive 에 등록
package migration

import "github.com/park-jun-woo/yongol/pkg/parser/ddl"

// applyAllowDestructiveHint marks a table as allowed to be dropped.
func applyAllowDestructiveHint(h *Hints, c ddl.HintComment) {
	if c.TableCtx == "" {
		return
	}
	h.AllowDestructive[c.TableCtx] = true
}
