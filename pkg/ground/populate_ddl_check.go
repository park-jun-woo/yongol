//ff:func feature=rule type=loader control=iteration dimension=1
//ff:what populateDDLCheck — DDL CHECK enum 값을 Ground에 등록
package ground

import (
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/rule"
)

func populateDDLCheck(g *rule.Ground, t ddl.Table) {
	for col, vals := range t.CheckEnums {
		checkSet := make(rule.StringSet, len(vals))
		for _, v := range vals {
			checkSet[v] = true
		}
		g.Lookup["DDL.check."+t.Name+"."+col] = checkSet
		g.Schemas["DDL.check."+t.Name+"."+col] = vals
	}
}
