//ff:func feature=rule type=loader control=iteration dimension=1
//ff:what populateDDLVarchar — DDL VARCHAR 길이를 Ground.Types에 등록
package ground

import (
	"strconv"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/rule"
)

func populateDDLVarchar(g *rule.Ground, t ddl.Table) {
	for col, c := range t.Columns {
		if c.VarcharLen <= 0 {
			continue
		}
		g.Types["DDL.varchar."+t.Name+"."+col] = strconv.Itoa(c.VarcharLen)
	}
}
