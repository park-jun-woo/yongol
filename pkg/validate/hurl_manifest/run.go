//ff:func feature=validate type=rule control=sequence topic=hurl-manifest
//ff:what Run — Hurl↔Manifest 교차 검증 실행 (XOH-06/07)

package hurl_manifest

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// Run executes every Hurl↔Manifest cross-validation rule.
func Run(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	diags = append(diags, xoh06AuthPrecondition(fs)...)
	diags = append(diags, xoh07CSRFOnMutation(fs)...)
	return diags
}
