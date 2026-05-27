//ff:func feature=gen-nestjs type=util control=sequence
//ff:what writeMethodBody — NestJS service 메서드 본문 (트랜잭션 분기) 작성

package ssac

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// writeMethodBody writes the service method body with optional transaction.
func writeMethodBody(b *strings.Builder, plan *ir.ServicePlan) {
	if plan.UsesTransaction {
		b.WriteString("    return this.prisma.$transaction(async (tx) => {\n")
		renderOpsBody(b, plan.Ops, "      ", "tx")
		b.WriteString("    });\n")
	} else {
		renderOpsBody(b, plan.Ops, "    ", "this.prisma")
	}
}
