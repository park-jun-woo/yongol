//ff:func feature=gen-fastapi type=util control=sequence
//ff:what writeHTTPHandler — ServicePlan 메타데이터 기반 FastAPI HTTP route handler 함수 작성

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// writeHTTPHandler writes a decorated FastAPI route handler function.
// Parameter declarations are driven by ServicePlan metadata:
//   - PathParams → typed path parameters as function args
//   - BodyFields → Pydantic model parameter
//   - QueryParams → typed query parameters as function args
func writeHTTPHandler(b *strings.Builder, plan *ir.ServicePlan) {
	decorator := pyHTTPDecorator(plan.HTTPMethod)
	routePath := routeSuffix(plan)
	funcName := snakeCase(plan.OperationID)

	method := strings.ToUpper(plan.HTTPMethod)
	hasBody := (method == "POST" || method == "PUT" || method == "PATCH") && len(plan.BodyFields) > 0
	isPreAuth := hasVerifyPasswordOp(plan.Ops)
	needsEventBus := hasPublishOp(plan.Ops)

	b.WriteString(fmt.Sprintf("@router.%s(\"%s\")\n", decorator, routePath))
	b.WriteString(fmt.Sprintf("async def %s(\n", funcName))

	writeHandlerParamDecls(b, plan, hasBody)
	writeHandlerDeps(b, isPreAuth, needsEventBus)
	b.WriteString("):\n")

	callArgs := buildHandlerCallArgs(plan, hasBody, isPreAuth, needsEventBus)
	b.WriteString(fmt.Sprintf("    return await svc.%s(%s)\n",
		funcName, strings.Join(callArgs, ", ")))
}
