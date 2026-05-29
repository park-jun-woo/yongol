//ff:func feature=validate type=rule control=iteration dimension=1 topic=ssac-ddl
//ff:what XDS-12 — @result 타입 DDL 테이블 없음

package ssac_ddl

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xds12ResultNoDDLTable validates XDS-12: SSaC @result type has no matching DDL table.
// Skip conditions: seq.Type == "call", seq.Package != "", primitive Go types, sqlc row types.
func xds12ResultNoDDLTable(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	tables := canonicalDDLTableSet(fs)
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		diags = append(diags, collectFuncResultDiags(fs, tables, fn)...)
	}
	return diags
}
