//ff:func feature=validate type=util control=iteration dimension=2 topic=ssac-ddl
//ff:what ddlTableSet — fs.DDLTables → name set (Ground 미적재 fallback 포함)

package ssac_ddl

import "github.com/park-jun-woo/yongol/pkg/yongol"

// ddlTableSet returns the set of DDL table names, preferring Ground lookup.
func ddlTableSet(fs *yongol.Fullstack) map[string]bool {
	if g := fs.Ground(); g != nil {
		if set, ok := g.Lookup["DDL.table"]; ok && set != nil {
			return set
		}
	}
	set := make(map[string]bool, len(fs.DDLTables))
	for _, t := range fs.DDLTables {
		set[t.Name] = true
	}
	return set
}
