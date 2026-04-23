//ff:func feature=migration type=parser control=sequence topic=migration-hints
//ff:what applyDataMigrationHint — @data_migration file=... 코멘트를 Hints.DataMigrations 에 등록
package migration

import "github.com/park-jun-woo/yongol/pkg/parser/ddl"

// applyDataMigrationHint stores a sidecar file path for a table.
func applyDataMigrationHint(h *Hints, c ddl.HintComment) {
	file := c.Args["file"]
	if file == "" || c.TableCtx == "" {
		return
	}
	h.DataMigrations[c.TableCtx] = file
}
