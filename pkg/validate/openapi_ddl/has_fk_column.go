//ff:func feature=validate type=util control=iteration dimension=2 topic=openapi-ddl
//ff:what hasFKColumn — srcTable의 colName이 dstTable을 참조하는 FK인지 확인

package openapi_ddl

import "github.com/park-jun-woo/yongol/pkg/yongol"

// hasFKColumn checks if srcTable has a FK column named colName that references dstTable.
func hasFKColumn(fs *yongol.Fullstack, srcTable, colName, dstTable string) bool {
	for _, t := range fs.DDLTables {
		if t.Name != srcTable {
			continue
		}
		for _, fk := range t.ForeignKeys {
			if fk.Column == colName && fk.RefTable == dstTable {
				return true
			}
		}
		return false
	}
	return false
}
