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
	writeConstructorParams(b, plan)
	methodName := lcFirst(plan.OperationID)
	b.WriteString(fmt.Sprintf("  async %s(%s): Promise<any> {\n",
		methodName, renderServiceParams(plan)))
	writeMethodBody(b, plan)
	b.WriteString("  }\n")
	b.WriteString("}\n")
}
