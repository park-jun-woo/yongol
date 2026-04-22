//ff:func feature=validate type=util control=iteration dimension=2 topic=ssac-ddl
//ff:what ddlColumnSet — 특정 테이블의 컬럼 name set (Ground 미적재 fallback 포함)

package ssac_ddl

import "github.com/park-jun-woo/yongol/pkg/yongol"

// ddlColumnSet returns the set of column names for the given table,
// preferring Ground lookup. Returns (set, true) when the table exists.
func ddlColumnSet(fs *yongol.Fullstack, table string) (map[string]bool, bool) {
	if g := fs.Ground(); g != nil {
		if set, ok := g.Lookup["DDL.column."+table]; ok && set != nil {
			return set, true
		}
	}
	for _, t := range fs.DDLTables {
		if t.Name != table {
			continue
		}
		set := make(map[string]bool, len(t.Columns))
		for col := range t.Columns {
			set[col] = true
		}
		return set, true
	}
	return nil, false
}
