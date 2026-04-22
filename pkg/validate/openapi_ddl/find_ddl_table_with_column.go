//ff:func feature=validate type=util control=iteration dimension=1 topic=openapi-ddl
//ff:what findDDLTableWithColumn — 특정 컬럼을 포함한 첫 DDL 테이블 조회

package openapi_ddl

import (
	"github.com/park-jun-woo/yongol/pkg/yongol"
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// findDDLTableWithColumn returns a pointer to the first DDL table that contains
// a column named col, or nil when no table matches.
func findDDLTableWithColumn(fs *yongol.Fullstack, col string) *ddl.Table {
	for i := range fs.DDLTables {
		if _, ok := fs.DDLTables[i].Columns[col]; ok {
			return &fs.DDLTables[i]
		}
	}
	return nil
}
