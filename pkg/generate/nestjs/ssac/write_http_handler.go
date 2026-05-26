//ff:func feature=gen-nestjs type=util control=sequence
//ff:what writeHTTPHandler — NestJS HTTP route handler 메서드 본문 작성

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// writeHTTPHandler writes a decorated HTTP handler method.
func writeHTTPHandler(b *strings.Builder, plan *ir.ServicePlan) {
	decorator := nestHTTPDecorator(plan.HTTPMethod)
	routeSuffix := controllerRouteSuffix(plan)
	methodName := lcFirst(plan.OperationID)

	b.WriteString(fmt.Sprintf("  @%s('%s')\n", decorator, routeSuffix))
	b.WriteString(fmt.Sprintf("  async %s(\n", methodName))
	b.WriteString("    @Req() req: any,\n")
	b.WriteString("    @Param() params: any,\n")
	b.WriteString("    @Body() body: any,\n")
	b.WriteString("  ) {\n")
	b.WriteString(fmt.Sprintf("    return this.service.%s(params, body, req.user);\n", methodName))
	b.WriteString("  }\n")
}
