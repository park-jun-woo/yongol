//ff:func feature=validate type=util control=sequence topic=migration-hints
//ff:what MIG-001 헬퍼 — @rename RenameColumn 단건 검증

package migration

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/generate/migration"
)

// mig001CheckRenameColumn validates a single RenameColumnHint and returns
// zero, one, or two diagnostics depending on which end is missing.
func mig001CheckRenameColumn(prev, curr *migration.Schema, r migration.RenameColumnHint) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	if pt, pok := prev.Tables[r.Table]; pok && !hasColumn(pt, r.From) {
		diags = append(diags, diagnostic.Diagnostic{
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: fmt.Sprintf("[MIG-001] @rename from=%s not in previous snapshot %s", r.From, r.Table),
			Advice:  "fix the 'from' value to match the old column name",
		})
	}
	if ct, cok := curr.Tables[r.Table]; cok && !hasColumn(ct, r.To) {
		diags = append(diags, diagnostic.Diagnostic{
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: fmt.Sprintf("[MIG-001] @rename to=%s not in current DDL %s", r.To, r.Table),
			Advice:  "rename the column in specs/db/*.sql to match 'to'",
		})
	}
	return diags
}
