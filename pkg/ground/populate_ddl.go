//ff:func feature=rule type=loader control=iteration dimension=1
//ff:what populateDDL — DDL Table에서 테이블명, 컬럼, FK, 인덱스, CHECK 추출
package ground

import (
	"github.com/park-jun-woo/yongol/pkg/yongol"
	"github.com/park-jun-woo/yongol/pkg/rule"
)

func populateDDL(g *rule.Ground, fs *yongol.Fullstack) {
	tables := make(rule.StringSet)
	for _, t := range fs.DDLTables {
		tables[t.Name] = true
		cols := make(rule.StringSet, len(t.Columns))
		for col := range t.Columns {
			cols[col] = true
		}
		g.Lookup["DDL.column."+t.Name] = cols
		populateDDLIndexes(g, t)
		populateDDLCheck(g, t)
		populateDDLVarchar(g, t)
		for col, def := range t.Defaults {
			g.Flags["DDL.default."+t.Name+"."+col] = true
			g.Types["DDL.default.value."+t.Name+"."+col] = def
		}
	}
	g.Lookup["DDL.table"] = tables
}
