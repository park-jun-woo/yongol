//ff:func feature=validate type=rule control=iteration dimension=1 topic=manifest-infra
//ff:what XDN-06 — claim 타입 ↔ DDL 컬럼 타입 정합 검증 (UUID 포함 매트릭스)

package manifest_ddl

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func xdn06ClaimDDLType(fs *yongol.Fullstack) []diagnostic.Diagnostic {
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
		if !def.Typed {
			continue
		}
		col, ok := tbl.Columns[def.Key]
		if !ok {
			continue
		}
		if !claimDDLTypeCompatible(def.GoType, col) {
			diags = append(diags, xdn06TypeMismatchDiag(field, auth.UserTable, def, col))
		}
	}
	return diags
}
