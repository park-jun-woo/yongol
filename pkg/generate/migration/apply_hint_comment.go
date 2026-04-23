//ff:func feature=migration type=parser control=selection topic=migration-hints
//ff:what applyHintComment — 단일 HintComment 의 Tag 에 따라 Hints 업데이트
package migration

import "github.com/park-jun-woo/yongol/pkg/parser/ddl"

// applyHintComment routes c to the specific tag handler.
func applyHintComment(h *Hints, c ddl.HintComment) {
	switch c.Tag {
	case "rename":
		applyRenameHint(h, c)
	case "cast":
		applyCastHint(h, c)
	case "backfill":
		applyBackfillHint(h, c)
	case "data_migration":
		applyDataMigrationHint(h, c)
	case "allow_destructive":
		applyAllowDestructiveHint(h, c)
	}
}
