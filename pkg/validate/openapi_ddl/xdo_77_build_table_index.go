//ff:func feature=validate type=util control=iteration dimension=2 topic=openapi-ddl
//ff:what xdo77BuildTableIndex — DDLTables → tableName → column → Go 타입 맵 빌드

package openapi_ddl

import (
	"github.com/park-jun-woo/yongol/pkg/generate/gogin/types"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xdo77BuildTableIndex builds a table-name → column → Go-type index from
// fs.DDLTables for fast XDO-77 lookup.
func xdo77BuildTableIndex(fs *yongol.Fullstack) map[string]map[string]string {
	tableIndex := make(map[string]map[string]string, len(fs.DDLTables))
	for _, t := range fs.DDLTables {
		cols := make(map[string]string, len(t.Columns))
		for col, c := range t.Columns {
			goType := types.GoTypeOf(c)
			if goType == "string" && isUUIDRawType(c) {
				goType = "uuid"
			}
			cols[col] = goType
		}
		tableIndex[t.Name] = cols
	}
	return tableIndex
}
