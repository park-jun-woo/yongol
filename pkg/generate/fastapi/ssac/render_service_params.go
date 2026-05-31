//ff:func feature=gen-fastapi type=util control=sequence
//ff:what renderServiceParams — ServicePlan 메타데이터 기반 Python service 함수 파라미터 목록 생성

package ssac

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// renderServiceParams produces the Python parameter list for the service
// function based on the plan's trigger kind and OpenAPI metadata. Path,
// body, and query parameters are typed individually.
func renderServiceParams(plan *ir.ServicePlan) string {
	if plan.TriggerKind == ir.TriggerSubscribe {
		return renderSubscribeParams(plan)
	}

	method := strings.ToUpper(plan.HTTPMethod)
	hasBody := (method == "POST" || method == "PUT" || method == "PATCH") && len(plan.BodyFields) > 0

	// Defensive: SSaC Ops may reference body/path/query even when OpenAPI
	// metadata is absent or incomplete. Force body flag when Ops contain
	// FieldArgs with LocBody.
	bodyFallback := false
	if !hasBody && opsReferenceBody(plan.Ops) {
		hasBody = true
		bodyFallback = true
	}

	params := []string{"session: AsyncSession"}
	params = append(params, renderPathParams(plan.PathParams)...)
	params = append(params, renderBodyParam(plan, hasBody, bodyFallback)...)
	params = append(params, renderQueryParams(plan.QueryParams)...)
	params = append(params, "current_user: dict | None = None")

	if hasPublishOp(plan.Ops) {
		params = append(params, "event_bus: EventBus | None = None")
	}

	return strings.Join(params, ", ")
}
