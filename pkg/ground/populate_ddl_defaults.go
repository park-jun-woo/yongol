//ff:func feature=rule type=loader control=iteration dimension=1
//ff:what populateDDLDefaults — DDL DEFAULT 'literal' 을 Flags / Types 에 등록
package ground

import (
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/rule"
)

// populateDDLDefaults registers each column with a DEFAULT '<literal>' on a
// table to Ground.Flags ("DDL.default.<table>.<col>") and Ground.Types
// ("DDL.default.value.<table>.<col>" → literal).
func populateDDLDefaults(g *rule.Ground, t ddl.Table) {
	for col, c := range t.Columns {
		if !c.HasDefault {
			continue
		}
		g.Flags["DDL.default."+t.Name+"."+col] = true
		g.Types["DDL.default.value."+t.Name+"."+col] = c.DefaultLiteral
	}
}
