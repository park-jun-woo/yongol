//ff:func feature=gen-nestjs type=util control=sequence
//ff:what renderServiceParams — ServicePlan 트리거 유형에 따른 메서드 파라미터 목록 생성

package ssac

import "github.com/park-jun-woo/yongol/pkg/generate/ir"

// renderServiceParams produces the TypeScript parameter list for the service
// method based on the plan's trigger kind.
func renderServiceParams(plan *ir.ServicePlan) string {
	if plan.TriggerKind == ir.TriggerSubscribe {
		return "payload: any"
	}
	return "params: any, body: any, user?: any"
}
