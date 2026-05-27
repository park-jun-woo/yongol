//ff:func feature=gen-nestjs type=util control=iteration dimension=1
//ff:what isColumnUnique — 단일 컬럼 unique 인덱스 존재 여부 확인

package prisma

import "github.com/park-jun-woo/yongol/pkg/parser/ddl"

// isColumnUnique checks if the column has a single-column unique index.
func isColumnUnique(colName string, indexes []ddl.Index) bool {
	for _, idx := range indexes {
		if idx.IsUnique && len(idx.Columns) == 1 && idx.Columns[0] == colName {
			return true
		}
	}
	return false
}
