//ff:func feature=gen-nestjs type=util control=iteration dimension=1
//ff:what writeConstructorParams — NestJS service 생성자 DI 파라미터 작성

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// writeConstructorParams writes the constructor DI parameters.
func writeConstructorParams(b *strings.Builder, plan *ir.ServicePlan) {
	b.WriteString("  constructor(\n")
	b.WriteString("    private readonly prisma: PrismaService,\n")
	if hasPublishOp(plan.Ops) {
		b.WriteString("    private readonly queue: QueueService,\n")
	}
	if hasAuthOp(plan.Ops) {
		b.WriteString("    private readonly authz: AuthzService,\n")
	}
	for _, pkg := range collectExternalOpsPackages(plan.Ops) {
		svcName := pkg + "Service"
		b.WriteString(fmt.Sprintf("    private readonly %s: %s,\n",
			svcName, strings.ToUpper(pkg[:1])+pkg[1:]+"Service"))
	}
	b.WriteString("  ) {}\n\n")
}
