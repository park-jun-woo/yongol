//ff:func feature=validate type=rule control=sequence topic=hurl-statemachine
//ff:what Run — Hurl↔StateMachine 교차 검증 실행 (XOH-05)

package hurl_statemachine

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// Run executes every Hurl↔StateMachine cross-validation rule.
func Run(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	return xoh05StateTransitionOrder(fs)
}
