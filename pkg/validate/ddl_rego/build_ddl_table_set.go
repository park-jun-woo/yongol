//ff:func feature=validate type=util control=iteration dimension=1 topic=policy-check
//ff:what buildDDLTableSet — build the set of DDL table names (prefers Ground.Lookup, falls back to fs.DDLTables)
package ddl_rego

import (
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// buildDDLTableSet returns the set of DDL table names. Prefers Ground.Lookup
// ("DDL.table") when populated, else falls back to iterating fs.DDLTables so
// rules remain usable in unit tests that construct Fullstack without Ground.
func buildDDLTableSet(fs *yongol.Fullstack, g *rule.Ground) map[string]bool {
	if g != nil {
		if tables := g.Lookup["DDL.table"]; tables != nil {
			return tables
		}
	}
	tables := make(map[string]bool, len(fs.DDLTables))
	for _, t := range fs.DDLTables {
		tables[t.Name] = true
	}
	return tables
}
