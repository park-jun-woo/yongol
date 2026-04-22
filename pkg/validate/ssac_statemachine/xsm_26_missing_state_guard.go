//ff:func feature=validate type=rule control=iteration dimension=1 topic=states
//ff:what XSM-26 — 상태 전이 참여하는데 @state 없음

package ssac_statemachine

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xsm26MissingStateGuard validates XSM-26: 상태 전이 참여하는데 @state 없음
func xsm26MissingStateGuard(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	funcByName := buildFuncByName(fs.ServiceFuncs)
	guardFuncs := collectGuardStateFuncs(fs.ServiceFuncs)
	var diags []diagnostic.Diagnostic
	for _, d := range fs.StateDiagrams {
		diags = append(diags, collectMissingGuards(d.ID, d.Events(), funcByName, guardFuncs)...)
	}
	return diags
}
