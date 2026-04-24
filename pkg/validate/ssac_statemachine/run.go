//ff:func feature=validate type=rule control=sequence topic=ssac-statemachine
//ff:what Run — SSaC↔StateMachine 교차 검증 실행 (XMS-*, XSM-*)
package ssac_statemachine

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// Run executes all SSaC↔StateMachine cross-validation rules.
func Run(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	diags = append(diags, xsm23TransitionToFunc(fs)...)
	diags = append(diags, xms24StateDiagramExists(fs)...)
	diags = append(diags, xms25StateEvent(fs)...)
	diags = append(diags, xsm26MissingStateGuard(fs)...)
	diags = append(diags, xsm27StateIntentDeclaration(fs)...)
	return diags
}
