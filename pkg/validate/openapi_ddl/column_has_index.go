//ff:func feature=validate type=util control=iteration dimension=2 topic=openapi-ddl
//ff:what columnHasIndex — 컬럼이 PK 또는 인덱스의 선행 컬럼인지 확인

package openapi_ddl

import "github.com/park-jun-woo/yongol/pkg/yongol"

// columnHasIndex reports whether col is the primary key or the leading column
// of any index in the given table.
func columnHasIndex(fs *yongol.Fullstack, table, col string) bool {
	for _, t := range fs.DDLTables {
		if t.Name != table {
			continue
		}
		for _, pk := range t.PrimaryKey {
			if pk == col {
				return true
			}
		}
		for _, idx := range t.Indexes {
			if len(idx.Columns) > 0 && idx.Columns[0] == col {
				return true
			}
		}
		return false
	}
	return false
}
