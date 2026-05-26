//ff:func feature=gen-fastapi type=util control=sequence
//ff:what renderServiceParams — ServicePlan 트리거 유형에 따른 Python 함수 파라미터 목록 생성

package ssac

import "github.com/park-jun-woo/yongol/pkg/generate/ir"

// renderServiceParams produces the Python parameter list for the service
// function based on the plan's trigger kind.
func renderServiceParams(plan *ir.ServicePlan) string {
	if plan.TriggerKind == ir.TriggerSubscribe {
		return "session: AsyncSession, payload: dict"
	}
	return "session: AsyncSession, params: dict, body: dict, user: dict | None = None"
}
