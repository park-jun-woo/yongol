//ff:func feature=validate type=rule control=iteration dimension=1 topic=migration-hints
//ff:what MIG-001 — @rename 힌트의 from/to 가 실제 스키마와 불일치할 때 ERROR
package migration

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/generate/migration"
)

// Mig001RenameMismatch checks each @rename hint:
//   - for RenameColumn: prev must contain (table, from) and curr must contain (table, to)
//   - for RenameTable:  prev must contain from and curr must contain to
func Mig001RenameMismatch(prev, curr *migration.Schema, hints *migration.Hints) []diagnostic.Diagnostic {
	if hints == nil {
		return nil
	}
	var diags []diagnostic.Diagnostic
	for _, r := range hints.RenameTables {
		diags = append(diags, mig001CheckRenameTable(prev, curr, r)...)
	}
	for _, r := range hints.RenameColumns {
		diags = append(diags, mig001CheckRenameColumn(prev, curr, r)...)
	}
	return diags
}
