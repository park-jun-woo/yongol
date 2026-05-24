//ff:func feature=validate type=util control=iteration dimension=1 topic=ssac-structural
//ff:what xss60BuildTableMap — DDLTables → tableName → *ddl.Table O(1) lookup 맵 생성

package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xss60BuildTableMap creates a tableName → *ddl.Table map for O(1) DDL column
// lookups.
func xss60BuildTableMap(fs *yongol.Fullstack) map[string]*ddl.Table {
	out := make(map[string]*ddl.Table, len(fs.DDLTables))
	for i := range fs.DDLTables {
		t := &fs.DDLTables[i]
		out[t.Name] = t
	}
	return out
}
