//ff:func feature=validate type=util control=iteration dimension=1 topic=openapi-ddl
//ff:what findTableWithColumn — 주어진 컬럼이 존재하는 첫 번째 DDL 테이블명 반환

package openapi_ddl

import "github.com/park-jun-woo/yongol/pkg/yongol"

// findTableWithColumn returns the first DDL table name that contains col.
// Returns "" if no table has the column.
func findTableWithColumn(fs *yongol.Fullstack, col string) string {
	for _, t := range fs.DDLTables {
		if _, ok := t.Columns[col]; ok {
			return t.Name
		}
	}
	return ""
}
