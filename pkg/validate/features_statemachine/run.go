//ff:func feature=validate type=rule control=sequence topic=features-statemachine
//ff:what Run — Features↔StateMachine 교차 검증 실행 (XFS-*)
package features_statemachine

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// Run executes all Features↔StateMachine cross-validation rules.
func Run(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	diags = append(diags, xfs01StatesInDiagram(fs)...)
	return diags
}
