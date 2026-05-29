//ff:func feature=rule type=loader control=iteration dimension=1
//ff:what populateDDLFuncManaged — DDL Table.FuncManaged 를 Flags["funcManaged.<table>"]로 투영
package ground

import (
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// populateDDLFuncManaged projects parsed `-- @func-managed` table annotations
// captured in ddl.Table onto rule.Ground.Flags:
//
//	Flags["funcManaged.<table>"] = true   // table is managed by a @call'd func/RPC
//
// Only XSD-55 consults this flag to exempt RPC-delegated tables; unlike
// `@archived`, the table is alive and other rules still apply.
func populateDDLFuncManaged(g *rule.Ground, fs *yongol.Fullstack) {
	for _, t := range fs.DDLTables {
		if t.FuncManaged {
			g.Flags["funcManaged."+t.Name] = true
		}
	}
}
