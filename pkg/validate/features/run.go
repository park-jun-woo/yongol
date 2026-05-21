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
	diags = append(diags, ft03HashMismatch(fs)...)
	diags = append(diags, ft16TablesRequired(fs)...)
	diags = append(diags, ft10HasManyRef(fs)...)
	diags = append(diags, ft11BelongsToRef(fs)...)
	diags = append(diags, ft12Bidirectional(fs)...)
	diags = append(diags, ft17FeatureTableRequired(fs)...)
	diags = append(diags, ft13FeatureTableRef(fs)...)
	return diags
}
