//ff:func feature=gen-fastapi type=util control=sequence
//ff:what writeServiceFunc — FastAPI service async 함수 본문 작성

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// writeServiceFunc writes the async service function definition and body.
func writeServiceFunc(b *strings.Builder, plan *ir.ServicePlan) {
	funcName := snakeCase(plan.OperationID)
	params := renderServiceParams(plan)

	b.WriteString(fmt.Sprintf("async def %s(%s):\n", funcName, params))

	if plan.UsesTransaction {
		b.WriteString("    async with session.begin():\n")
		renderOpsBody(b, plan.Ops, "        ", "session")
	} else {
		renderOpsBody(b, plan.Ops, "    ", "session")
	}
}
