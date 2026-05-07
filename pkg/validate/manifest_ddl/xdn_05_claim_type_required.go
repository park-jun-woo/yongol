//ff:func feature=validate type=rule control=iteration dimension=1 topic=manifest-infra
//ff:what XDN-05 — claim 타입 선언 필수 (col:type 형식 강제)

package manifest_ddl

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func xdn05ClaimTypeRequired(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if !isAuthActive(fs) {
		return nil
	}
	auth := fs.Manifest.Backend.Auth
	if len(auth.Claims) == 0 {
		return nil
	}
	var diags []diagnostic.Diagnostic
	for _, field := range sortedClaimFields(auth.Claims) {
		def := auth.Claims[field]
		if !def.Typed {
			diags = append(diags, xdn05MissingTypeDiag(field, def))
			continue
		}
		if !allowedClaimTypes[def.GoType] {
			diags = append(diags, xdn05InvalidTypeDiag(field, def))
		}
	}
	return diags
}
