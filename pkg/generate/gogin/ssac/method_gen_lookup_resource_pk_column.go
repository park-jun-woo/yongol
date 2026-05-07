//ff:func feature=gen-gogin type=util control=sequence
//ff:what lookupResourcePKColumn — 리소스명으로 DDL PK(id) 컬럼 조회

package ssac

import "github.com/park-jun-woo/yongol/pkg/parser/ddl"

// lookupResourcePKColumn resolves the PK (id) column for a resource name.
func (g *methodGen) lookupResourcePKColumn(resource string) *ddl.Column {
	return lookupDDLColumn(g.DDLTables, pascalCase(resource), "id")
}
