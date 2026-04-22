//ff:func feature=validate type=rule control=sequence topic=hurl-structural
//ff:what Run — Hurl 검증 전체 실행 (H-*)
package hurl

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// Run executes all Hurl validation rules.
func Run(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	diags = append(diags, h01DeprecatedFeature(fs)...)
	diags = append(diags, h02EmptyTestsDir(fs)...)
	return diags
}
