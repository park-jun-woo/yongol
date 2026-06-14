//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=openapi-ddl
//ff:what tbl — 컬럼명만 지정하는 ddl.Table 생성 헬퍼

package openapi_ddl

import "github.com/park-jun-woo/yongol/pkg/parser/ddl"

// tbl is a small ddl.Table constructor with column names only.
func tbl(name string, cols ...string) ddl.Table {
	m := map[string]ddl.Column{}
	for _, c := range cols {
		m[c] = ddl.Column{}
	}
	return ddl.Table{Name: name, Columns: m}
}
