//ff:func feature=validate type=util control=selection topic=manifest-infra
//ff:what xdn04CheckClaim — 단일 claim 의 Go 타입↔컬럼 Go 타입 정합 검사 (mismatch 시 1건 diag)

package manifest_ddl

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

// xdn04CheckClaim returns (diag, true) when the given claim's Go type
// disagrees with the DDL column's Go type. Missing column is silently
// deferred to XDN-03 ((Diagnostic{}, false)).
func xdn04CheckClaim(field, userTable string, def pmanifest.ClaimDef, columns map[string]ddl.Column) (diagnostic.Diagnostic, bool) {
	col, ok := columns[def.Key]
	if !ok {
		return diagnostic.Diagnostic{}, false
	}
	ddlGoType := ddl.GoTypeOf(col)
	claimGoType := def.GoType
	if claimGoType == "" {
		claimGoType = "string"
	}
	if claimTypeCompatible(claimGoType, ddlGoType) {
		return diagnostic.Diagnostic{}, false
	}
	return xdn04TypeMismatchDiag(field, userTable, claimGoType, ddlGoType, def), true
}
