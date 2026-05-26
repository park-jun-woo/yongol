//ff:func feature=gen-nestjs type=util control=sequence
//ff:what writeSubscribeHandler — NestJS Subscribe handler 메서드 본문 작성

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// writeSubscribeHandler writes a queue subscription handler method.
func writeSubscribeHandler(b *strings.Builder, plan *ir.ServicePlan) {
	b.WriteString(fmt.Sprintf("  // Subscribe handler for topic: %s\n", plan.Topic))
	b.WriteString(fmt.Sprintf("  async handle%s(payload: any) {\n", plan.OperationID))
	b.WriteString(fmt.Sprintf("    return this.service.%s(payload);\n", lcFirst(plan.OperationID)))
	b.WriteString("  }\n")
}
