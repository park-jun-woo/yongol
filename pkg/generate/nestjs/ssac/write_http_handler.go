//ff:func feature=gen-nestjs type=util control=sequence
//ff:what writeHTTPHandler — ServicePlan 메타데이터 기반 NestJS HTTP handler 렌더링

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// writeHTTPHandler writes a decorated HTTP handler method. Parameter
// decorators are selected based on ServicePlan metadata:
//   - @Param() for path parameters (all methods)
//   - @Body() for POST/PUT/PATCH (when BodyFields exist)
//   - @Query() for query parameters (when QueryParams exist)
func writeHTTPHandler(b *strings.Builder, plan *ir.ServicePlan) {
	decorator := nestHTTPDecorator(plan.HTTPMethod)
	routeSuffix := controllerRouteSuffix(plan)
	methodName := lcFirst(plan.OperationID)
	method := strings.ToUpper(plan.HTTPMethod)

	hasBody := (method == "POST" || method == "PUT" || method == "PATCH") && len(plan.BodyFields) > 0
	hasQuery := len(plan.QueryParams) > 0
	hasPath := len(plan.PathParams) > 0

	// Skip @Req() and req.user for pre-auth endpoints (login with
	// @verify-password).
	isPreAuth := hasVerifyPasswordOp(plan.Ops)

	b.WriteString(fmt.Sprintf("  @%s('%s')\n", decorator, routeSuffix))
	b.WriteString(fmt.Sprintf("  async %s(\n", methodName))
	if !isPreAuth {
		b.WriteString("    @Req() req: any,\n")
	}
	if hasPath {
		b.WriteString("    @Param() params: any,\n")
	}
	if hasBody {
		b.WriteString("    @Body() body: any,\n")
	}
	if hasQuery {
		b.WriteString("    @Query() query: any,\n")
	}
	b.WriteString("  ) {\n")

	// Build the service call arguments based on what's available.
	var callArgs []string
	if hasPath {
		callArgs = append(callArgs, "params")
	}
	if hasBody {
		callArgs = append(callArgs, "body")
	}
	if hasQuery {
		callArgs = append(callArgs, "query")
	}
	if !isPreAuth {
		callArgs = append(callArgs, "req.user")
	}

	b.WriteString(fmt.Sprintf("    return this.service.%s(%s);\n",
		methodName, strings.Join(callArgs, ", ")))
	b.WriteString("  }\n")
}
