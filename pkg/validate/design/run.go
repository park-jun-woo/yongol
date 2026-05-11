//ff:func feature=validate type=rule control=sequence topic=design-structural
//ff:what Run — DESIGN.md 내부 정합성 검증 전체 실행 (V-01~V-07)
package design

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// Run executes all DESIGN.md internal consistency rules.
func Run(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs.DesignSpec == nil {
		return nil
	}
	var diags []diagnostic.Diagnostic
	diags = append(diags, v01NameRequired(fs)...)
	diags = append(diags, v02HexValid(fs)...)
	diags = append(diags, v03TypographyRequired(fs)...)
	diags = append(diags, v04DimensionValid(fs)...)
	diags = append(diags, v05TokenRefResolve(fs)...)
	diags = append(diags, v06DuplicateHeading(fs)...)
	diags = append(diags, v07UnknownProp(fs)...)
	return diags
}
