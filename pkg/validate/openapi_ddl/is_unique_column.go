//ff:func feature=validate type=util control=iteration dimension=2 topic=openapi-ddl
//ff:what isUniqueColumn — 컬럼이 PRIMARY KEY 또는 단일 컬럼 UNIQUE 인덱스인지 확인

package openapi_ddl

import "github.com/park-jun-woo/yongol/pkg/yongol"

// isUniqueColumn reports whether col is the (single-column) primary key or
// has a (single-column) UNIQUE index in the given table.
func isUniqueColumn(fs *yongol.Fullstack, table, col string) bool {
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
			if idx.IsUnique && len(idx.Columns) == 1 && idx.Columns[0] == col {
				return true
			}
		}
		return false
	}
	return false
}
