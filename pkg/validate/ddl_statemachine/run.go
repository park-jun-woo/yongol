//ff:func feature=validate type=rule control=sequence topic=ddl-statemachine
//ff:what Run — DDL↔StateMachine 교차 검증 실행 (XDM-*)
package ddl_statemachine

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// Run executes all DDL↔StateMachine cross-validation rules.
func Run(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	diags = append(diags, xdm27StateFieldColumn(fs)...)
	diags = append(diags, xdm28DefaultInitialState(fs)...)
	return diags
}
