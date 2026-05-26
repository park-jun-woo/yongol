//ff:func feature=gen-nestjs type=util control=sequence
//ff:what writeServiceClass — NestJS service class 본문 작성

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// writeServiceClass writes the class declaration, constructor, and method body.
func writeServiceClass(b *strings.Builder, plan *ir.ServicePlan) {
	className := plan.OperationID + "Service"
	b.WriteString("@Injectable()\n")
	b.WriteString(fmt.Sprintf("export class %s {\n", className))
	b.WriteString("  constructor(\n")
	b.WriteString("    private readonly prisma: PrismaService,\n")
	if hasPublishOp(plan.Ops) {
		b.WriteString("    private readonly queue: QueueService,\n")
	}
	b.WriteString("  ) {}\n\n")

	methodName := lcFirst(plan.OperationID)
	b.WriteString(fmt.Sprintf("  async %s(%s): Promise<any> {\n",
		methodName, renderServiceParams(plan)))

	if plan.UsesTransaction {
		b.WriteString("    return this.prisma.$transaction(async (tx) => {\n")
		renderOpsBody(b, plan.Ops, "      ", "tx")
		b.WriteString("    });\n")
	} else {
		renderOpsBody(b, plan.Ops, "    ", "this.prisma")
	}

	b.WriteString("  }\n")
	b.WriteString("}\n")
}
