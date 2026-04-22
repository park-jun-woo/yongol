//ff:func feature=validate type=rule control=iteration dimension=1 topic=migration-hints
//ff:what MIG-001 — @rename 힌트의 from/to 가 실제 스키마와 불일치할 때 ERROR
package migration

import (
	"fmt"

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
	}
	for _, r := range hints.RenameColumns {
		pt, pok := prev.Tables[r.Table]
		ct, cok := curr.Tables[r.Table]
		if pok {
			if !hasColumn(pt, r.From) {
				diags = append(diags, diagnostic.Diagnostic{
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: fmt.Sprintf("[MIG-001] @rename from=%s not in previous snapshot %s", r.From, r.Table),
					Advice:  "fix the 'from' value to match the old column name",
				})
			}
		}
		if cok {
			if !hasColumn(ct, r.To) {
				diags = append(diags, diagnostic.Diagnostic{
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: fmt.Sprintf("[MIG-001] @rename to=%s not in current DDL %s", r.To, r.Table),
					Advice:  "rename the column in specs/db/*.sql to match 'to'",
				})
			}
		}
	}
	return diags
}

func hasColumn(t *migration.Table, name string) bool {
	for _, c := range t.Columns {
		if c.Name == name {
			return true
		}
	}
	return false
}
