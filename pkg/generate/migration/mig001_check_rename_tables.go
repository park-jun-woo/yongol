//ff:func feature=migration type=util control=iteration dimension=1
//ff:what mig001CheckRenameTables — @rename (table) from/to 가 prev/curr 스키마에 있는지 점검
package migration

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// mig001CheckRenameTables emits MIG-001 diagnostics when the from/to
// values of table rename hints don't match prev/curr schemas.
func mig001CheckRenameTables(prev, curr *Schema, rules []RenameTableHint) []diagnostic.Diagnostic {
	var out []diagnostic.Diagnostic
	for _, r := range rules {
		if _, ok := prev.Tables[r.From]; !ok {
			out = append(out, diagnostic.Diagnostic{
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: fmt.Sprintf("[MIG-001] @rename from=%s (table) not in previous snapshot", r.From),
				Advice:  "fix the 'from' value or remove the hint",
			})
		}
		if _, ok := curr.Tables[r.To]; !ok {
			out = append(out, diagnostic.Diagnostic{
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: fmt.Sprintf("[MIG-001] @rename to=%s (table) not in current DDL", r.To),
				Advice:  "rename the CREATE TABLE to match 'to'",
			})
		}
	}
	return out
}
