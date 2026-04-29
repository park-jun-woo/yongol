//ff:func feature=validate type=util control=iteration dimension=1 topic=ssac-sqlc
//ff:what buildDDLTableLookup — fs.DDLTables → table-name 키 / *Table 값 맵 (XQS-20 컬럼 매칭용)

package ssac_sqlc

import (
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// buildDDLTableLookup returns a map from DDL table name → *ddl.Table for
// O(1) lookup during RETURNING shape classification. The values point into
// fs.DDLTables; do not mutate.
func buildDDLTableLookup(fs *yongol.Fullstack) map[string]*ddl.Table {
	out := make(map[string]*ddl.Table, len(fs.DDLTables))
	for i := range fs.DDLTables {
		t := &fs.DDLTables[i]
		out[t.Name] = t
	}
	return out
}
