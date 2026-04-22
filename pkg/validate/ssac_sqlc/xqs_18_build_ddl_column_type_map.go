//ff:func feature=validate type=util control=iteration dimension=2 topic=ssac-sqlc
//ff:what buildXqs18DDLColumnTypeMap — build a tableName → column → Go-type map from DDLTables

package ssac_sqlc

import (
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// buildXqs18DDLColumnTypeMap builds tableName → columnName → Go type map from DDLTables.
func buildXqs18DDLColumnTypeMap(fs *yongol.Fullstack) map[string]map[string]string {
	result := make(map[string]map[string]string, len(fs.DDLTables))
	for _, t := range fs.DDLTables {
		cols := make(map[string]string, len(t.Columns))
		for col, goType := range t.Columns {
			cols[col] = goType
		}
		result[t.Name] = cols
	}
	return result
}
