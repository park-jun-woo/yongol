//ff:func feature=migration type=util control=iteration dimension=1
//ff:what mig001CheckRenameColumns — @rename (column) from/to 가 prev/curr 스키마 컬럼에 있는지 점검
package migration

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// mig001CheckRenameColumns emits MIG-001 diagnostics when column rename
// hint from/to columns are missing from prev/curr.
func mig001CheckRenameColumns(prev, curr *Schema, rules []RenameColumnHint) []diagnostic.Diagnostic {
	var out []diagnostic.Diagnostic
	for _, r := range rules {
		if pt, ok := prev.Tables[r.Table]; ok && !tableHasColumn(pt, r.From) {
			out = append(out, diagnostic.Diagnostic{
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: fmt.Sprintf("[MIG-001] @rename from=%s not in previous snapshot %s", r.From, r.Table),
				Advice:  "fix the 'from' value to match the old column name",
			})
		}
		if ct, ok := curr.Tables[r.Table]; ok && !tableHasColumn(ct, r.To) {
			out = append(out, diagnostic.Diagnostic{
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: fmt.Sprintf("[MIG-001] @rename to=%s not in current DDL %s", r.To, r.Table),
				Advice:  "rename the column in DDL to match 'to'",
			})
		}
	}
	return out
}
