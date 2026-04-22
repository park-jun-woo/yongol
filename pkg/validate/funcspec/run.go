//ff:func feature=validate type=rule control=sequence topic=funcspec-structural
//ff:what Run — FuncSpec 검증 전체 실행 (F-*, XFF-*)
package funcspec

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// Run executes all FuncSpec validation rules.
func Run(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	diags = append(diags, f01BuiltinOverride(fs)...)
	diags = append(diags, xff40FuncBodyTodo(fs)...)
	diags = append(diags, xff41FuncForbiddenImport(fs)...)
	return diags
}
