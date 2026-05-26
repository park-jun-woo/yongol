//ff:func feature=gen-nestjs type=util control=sequence
//ff:what writeControllerClass — NestJS controller class 본문 작성

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// writeControllerClass writes the class declaration and route handler(s).
func writeControllerClass(b *strings.Builder, plan *ir.ServicePlan) {
	controllerName := plan.OperationID + "Controller"
	serviceName := plan.OperationID + "Service"
	routePrefix := controllerRoutePrefix(plan)
	b.WriteString(fmt.Sprintf("@Controller('%s')\n", routePrefix))
	b.WriteString(fmt.Sprintf("export class %s {\n", controllerName))
	b.WriteString(fmt.Sprintf("  constructor(private readonly service: %s) {}\n\n", serviceName))

	if plan.TriggerKind == ir.TriggerHTTP {
		writeHTTPHandler(b, plan)
	} else {
		writeSubscribeHandler(b, plan)
	}

	b.WriteString("}\n")
}
