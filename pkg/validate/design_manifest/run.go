//ff:func feature=validate type=rule control=sequence topic=design-manifest
//ff:what Run — DESIGN.md↔Manifest 교차 검증 실행 (XNV-01 ~ XNV-02)
package design_manifest

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// Run executes all DESIGN.md↔Manifest cross-validation rules.
func Run(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	diags = append(diags, xnv01PathExists(fs)...)
	diags = append(diags, xnv02Undeclared(fs)...)
	return diags
}
