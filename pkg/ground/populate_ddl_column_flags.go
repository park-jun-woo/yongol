//ff:func feature=rule type=loader control=iteration dimension=1
//ff:what populateDDLColumnFlags — 한 테이블의 archived/sensitive 컬럼 어노테이션을 Flags 에 등록
package ground

import (
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/rule"
)

// populateDDLColumnFlags writes per-column archived / sensitive markers
// for a single table onto Ground.Flags.
func populateDDLColumnFlags(g *rule.Ground, t ddl.Table) {
	for col, c := range t.Columns {
		if c.Archived {
			g.Flags["archived."+t.Name+"."+col] = true
		}
		if c.Sensitive {
			g.Flags["sensitive."+t.Name+"."+col] = true
		}
	}
}
