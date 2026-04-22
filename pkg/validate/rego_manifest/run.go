//ff:func feature=validate type=rule control=sequence topic=rego-manifest
//ff:what Run — Rego↔Manifest 교차 검증 실행 (XNP-*, XPN-*)
package rego_manifest

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// Run executes all Rego↔Manifest cross-validation rules.
func Run(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	diags = append(diags, xnp53InputClaimsValues(fs)...)
	diags = append(diags, xnp63RoleManifest(fs)...)
	diags = append(diags, xpn54ClaimsToRego(fs)...)
	diags = append(diags, xpn64RolesToRego(fs)...)
	return diags
}
