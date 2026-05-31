//ff:func feature=gen-fastapi type=util control=iteration dimension=1
//ff:what writeRouterHandlers — plan 별 HTTP/subscribe handler 함수를 Router 본문에 출력

package ssac

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// writeRouterHandlers writes one handler function per plan, choosing HTTP or
// subscribe handler based on the plan's trigger kind.
func writeRouterHandlers(b *strings.Builder, plans []*ir.ServicePlan) {
	for _, plan := range plans {
		if plan.TriggerKind == ir.TriggerHTTP {
			writeHTTPHandler(b, plan)
		} else {
			writeSubscribeHandler(b, plan)
		}
		b.WriteString("\n")
	}
}
