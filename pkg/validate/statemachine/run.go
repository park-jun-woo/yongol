//ff:func feature=validate type=rule control=sequence topic=statemachine-structural
//ff:what Run — StateMachine 검증 전체 실행 (ST-*)
package statemachine

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// Run executes all StateMachine validation rules.
func Run(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	diags = append(diags, st01Parse(fs)...)
	return diags
}
