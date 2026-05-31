//ff:func feature=validate type=test control=iteration dimension=1 topic=ssac-sqlc
//ff:what TestSSaCSqlcHelpers — unit tests for the pure ssac_sqlc helper functions
package ssac_sqlc

import (
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func ddlTable(cols ...string) *ddl.Table {
	m := make(map[string]ddl.Column, len(cols))
	for _, c := range cols {
		m[c] = ddl.Column{Name: c}
	}
	return &ddl.Table{Columns: m}
}
