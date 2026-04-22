//ff:func feature=validate type=util control=iteration dimension=2 topic=policy-check
//ff:what buildDDLColumnIndex — build a local table → column-set map (inner-loop optimisation for rules)
package ddl_rego

import "github.com/park-jun-woo/yongol/pkg/yongol"

// buildDDLColumnIndex returns a table -> column-set map derived from
// fs.DDLTables. Ground.Lookup stores columns as "DDL.column.<table>" which
// requires one lookup per table; building a local map keeps the per-rule
// inner loop simple and avoids repeated key string concatenation.
func buildDDLColumnIndex(fs *yongol.Fullstack) map[string]map[string]bool {
	idx := make(map[string]map[string]bool, len(fs.DDLTables))
	for _, t := range fs.DDLTables {
		cols := make(map[string]bool, len(t.Columns))
		for col := range t.Columns {
			cols[col] = true
		}
		idx[t.Name] = cols
	}
	return idx
}
