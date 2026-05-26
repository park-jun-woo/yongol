//ff:func feature=gen-fastapi type=util control=sequence
//ff:what writeHTTPHandler — FastAPI HTTP route handler 함수 작성

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// writeHTTPHandler writes a decorated FastAPI route handler function.
func writeHTTPHandler(b *strings.Builder, plan *ir.ServicePlan) {
	decorator := pyHTTPDecorator(plan.HTTPMethod)
	routePath := routeSuffix(plan)
	funcName := snakeCase(plan.OperationID)

	b.WriteString(fmt.Sprintf("@router.%s(\"%s\")\n", decorator, routePath))
	b.WriteString(fmt.Sprintf("async def %s(\n", funcName))
	b.WriteString("    request: Request,\n")
	b.WriteString("    session: AsyncSession = Depends(get_session),\n")
	b.WriteString("):\n")
	b.WriteString(fmt.Sprintf("    params = request.path_params\n"))
	b.WriteString(fmt.Sprintf("    body = await request.json() if request.method in (\"POST\", \"PUT\", \"PATCH\") else {}\n"))
	b.WriteString(fmt.Sprintf("    user = getattr(request.state, \"user\", None)\n"))
	b.WriteString(fmt.Sprintf("    return await svc.%s(session, params, body, user)\n", funcName))
}
