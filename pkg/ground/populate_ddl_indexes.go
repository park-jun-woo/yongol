//ff:func feature=rule type=loader control=iteration dimension=1
//ff:what populateDDLIndexes — DDL 인덱스 컬럼을 Ground.Lookup에 등록
package ground

import (
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/rule"
)

func populateDDLIndexes(g *rule.Ground, t ddl.Table) {
	indexed := make(rule.StringSet)
	for _, idx := range t.Indexes {
		for _, col := range idx.Columns {
			indexed[col] = true
		}
	}
	g.Lookup["DDL.index."+t.Name] = indexed
}
