//ff:func feature=validate type=rule control=sequence topic=features-structural
//ff:what Run — Features 검증 전체 실행 (FT-*)
package features

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// Run executes all Features validation rules.
func Run(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	diags = append(diags, ft01DuplicateOp(fs)...)
	diags = append(diags, ft02DuplicatePath(fs)...)
	return diags
}
