//ff:func feature=validate type=rule control=sequence topic=ssac-manifest
//ff:what Run — SSaC↔Manifest 교차 검증 실행 (XNS-*)
package ssac_manifest

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// Run executes all SSaC↔Manifest cross-validation rules.
func Run(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	diags = append(diags, xns48CurrentUserClaims(fs)...)
	diags = append(diags, xns49CurrentUserField(fs)...)
	diags = append(diags, xns56QueueRequired(fs)...)
	diags = append(diags, xns57MemoryTxPublish(fs)...)
	diags = append(diags, xns73JwtCallClaims(fs)...)
	return diags
}
