//ff:func feature=validate type=rule control=sequence topic=domain-security
//ff:what Run — multi-domain 보안 교차 검증 실행 (XDS-80~82, XDO-90, XMO-20~22)
package domain_security

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// Run executes all domain security cross-validation rules.
// If the project has no domains key in manifest (single-domain), all rules are skipped.
func Run(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs.Manifest == nil || len(fs.Manifest.Domains) == 0 {
		return nil
	}

	var diags []diagnostic.Diagnostic
	diags = append(diags, xds80AdminPublicAccess(fs)...)
	diags = append(diags, xds81InternalSecurity(fs)...)
	diags = append(diags, xds82PublicDeleteNoRego(fs)...)
	diags = append(diags, xdo90DuplicateOperationID(fs)...)
	diags = append(diags, xmo20PublicUnconsumed(fs)...)
	diags = append(diags, xmo21AdminUnconsumed(fs)...)
	diags = append(diags, xmo22CrossDomainCall(fs)...)
	return diags
}
