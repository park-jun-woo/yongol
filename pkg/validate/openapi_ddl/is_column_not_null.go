//ff:func feature=validate type=util control=iteration dimension=1 topic=openapi-ddl
//ff:what isColumnNotNull — 컬럼이 NOT NULL (제약 또는 PK) 인지 판정

package openapi_ddl

import (
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// isColumnNotNull reports whether col is NOT NULL in the given table (either
// via an explicit NOT NULL constraint or because it is part of the primary key).
func isColumnNotNull(tbl *ddl.Table, col string) bool {
	if c, ok := tbl.Columns[col]; ok && c.NotNull {
		return true
	}
	for _, pk := range tbl.PrimaryKey {
		if pk == col {
			return true
		}
	}
	return false
}
