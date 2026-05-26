//ff:func feature=gen-fastapi type=util control=sequence
//ff:what writeSubscribeHandler — FastAPI Subscribe handler 함수 작성

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// writeSubscribeHandler writes a queue subscription handler function.
func writeSubscribeHandler(b *strings.Builder, plan *ir.ServicePlan) {
	funcName := snakeCase(plan.OperationID)
	b.WriteString(fmt.Sprintf("# Subscribe handler for topic: %s\n", plan.Topic))
	b.WriteString(fmt.Sprintf("async def handle_%s(session: AsyncSession, payload: dict):\n", funcName))
	b.WriteString(fmt.Sprintf("    return await svc.%s(session, payload)\n", funcName))
}
