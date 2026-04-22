//ff:func feature=validate type=util control=iteration dimension=1 topic=ssac-ddl
//ff:what buildReferencedTables — SSaC 전체에서 참조된 DDL 테이블 이름을 수집

package ssac_ddl

import "github.com/park-jun-woo/yongol/pkg/parser/ssac"

// buildReferencedTables collects DDL table names referenced by SSaC @model/@result.
func buildReferencedTables(funcs []ssac.ServiceFunc) map[string]bool {
	tables := make(map[string]bool)
	for _, fn := range funcs {
		for _, seq := range fn.Sequences {
			collectReferencedTable(seq, tables)
		}
	}
	return tables
}
