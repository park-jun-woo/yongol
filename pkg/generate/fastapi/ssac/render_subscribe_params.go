//ff:func feature=gen-fastapi type=util control=sequence
//ff:what renderSubscribeParams — subscribe trigger service 함수 파라미터 목록 생성

package ssac

import "github.com/park-jun-woo/yongol/pkg/generate/ir"

// renderSubscribeParams produces the Python parameter list for a subscribe
// trigger service function.
func renderSubscribeParams(plan *ir.ServicePlan) string {
	base := "session: AsyncSession, payload: dict"
	if hasPublishOp(plan.Ops) {
		base += ", event_bus: EventBus | None = None"
	}
	return base
}
