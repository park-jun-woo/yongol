//ff:func feature=rule type=loader control=iteration dimension=1
//ff:what populateDDLArchived — DDL Table.Archived/컬럼 @archived/@sensitive 를 Flags로 투영
package ground

import (
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// populateDDLArchived projects parsed `-- @archived` and `-- @sensitive`
// annotations captured in ddl.Table onto rule.Ground.Flags:
//
//	Flags["archived.<table>"]           = true   // table is archived
//	Flags["archived.<table>.<column>"]  = true   // column is archived
//	Flags["sensitive.<table>.<column>"] = true   // column is sensitive
//
// XSD-55 / XOD-10 / XDD-61 consult these flags to exempt opted-out items.
func populateDDLArchived(g *rule.Ground, fs *yongol.Fullstack) {
	for _, t := range fs.DDLTables {
		if t.Archived {
			g.Flags["archived."+t.Name] = true
		}
		populateDDLColumnFlags(g, t)
	}
}
