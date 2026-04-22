//ff:func feature=validate type=util control=iteration dimension=1 topic=openapi-ddl
//ff:what tableHasColumn — 특정 테이블에 컬럼이 있는지 확인

package openapi_ddl

import "github.com/park-jun-woo/yongol/pkg/yongol"

// tableHasColumn reports whether the given table contains the given column.
func tableHasColumn(fs *yongol.Fullstack, table, col string) bool {
	for _, t := range fs.DDLTables {
		if t.Name != table {
			continue
		}
		_, ok := t.Columns[col]
		return ok
	}
	return false
}
