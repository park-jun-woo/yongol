//ff:func feature=validate type=rule control=iteration dimension=1 topic=manifest-infra
//ff:what XDN-04 — claim Go 타입과 user_table 컬럼 DDL Go 타입 정합 검증

package manifest_ddl

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xdn04ClaimColumnType validates that the Go type declared on each claim
// (`<col>:<go_type>`, default `string`) matches the Go type the DDL
// parser inferred for the corresponding user_table column. Comparison is
// at the Go-type level — this matches how XDO-77 already works and
// follows the parser's published type table (BIGINT/INTEGER/SERIAL →
// int64, VARCHAR/TEXT/UUID/CHAR → string, BOOLEAN/BOOL → bool).
//
// Claims whose column is missing are deferred to XDN-03 (no diagnostic
// here).
func xdn04ClaimColumnType(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if !isAuthActive(fs) {
		return nil
	}
	auth := fs.Manifest.Backend.Auth
	if auth.UserTable == "" {
		return nil
	}
	tbl := findUserTable(fs, auth.UserTable)
	if tbl == nil {
		return nil
	}
	var diags []diagnostic.Diagnostic
	for _, field := range sortedClaimFields(auth.Claims) {
		def := auth.Claims[field]
		if d, ok := xdn04CheckClaim(field, auth.UserTable, def, tbl.Columns); ok {
			diags = append(diags, d)
		}
	}
	return diags
}
