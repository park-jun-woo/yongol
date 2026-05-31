//ff:func feature=validate type=test control=sequence topic=policy-check
//ff:what XDP-65 test — Rego role 값이 DDL role CHECK 제약에 선언되어야 함
package ddl_rego

import (
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func roleTable(roles ...string) ddl.Table {
	return ddl.Table{
		Name: "users",
		Columns: map[string]ddl.Column{
			"role": {Name: "role", CheckEnum: roles},
		},
	}
}
