//ff:func feature=validate type=util control=sequence topic=migration-hints
//ff:what MIG-001 헬퍼 — @rename RenameTable 단건 검증

package migration

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/generate/migration"
)

// mig001CheckRenameTable validates a single RenameTableHint and returns
// zero, one, or two diagnostics depending on which end is missing.
func mig001CheckRenameTable(prev, curr *migration.Schema, r migration.RenameTableHint) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	if _, ok := prev.Tables[r.From]; !ok {
		diags = append(diags, diagnostic.Diagnostic{
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: fmt.Sprintf("[MIG-001] @rename from=%s (table) not present in previous snapshot", r.From),
			Advice:  "remove the hint or fix the 'from' to match the old snapshot table name",
		})
	}
	if _, ok := curr.Tables[r.To]; !ok {
		diags = append(diags, diagnostic.Diagnostic{
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: fmt.Sprintf("[MIG-001] @rename to=%s (table) not present in current DDL", r.To),
			Advice:  "rename the CREATE TABLE in specs/db/*.sql to match 'to'",
		})
	}
	return diags
}
