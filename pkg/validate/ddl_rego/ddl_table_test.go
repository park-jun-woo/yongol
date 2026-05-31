//ff:func feature=validate type=test control=iteration dimension=1 topic=policy-check
//ff:what XDP-31 test — @ownership table must exist in DDL (Ground 기반)
package ddl_rego

import (
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func ddlTable(name string, cols ...string) ddl.Table {
	m := map[string]ddl.Column{}
	for _, c := range cols {
		m[c] = ddl.Column{Name: c}
	}
	return ddl.Table{Name: name, Columns: m}
}
