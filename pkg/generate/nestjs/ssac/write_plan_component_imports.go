//ff:func feature=gen-nestjs type=util control=iteration dimension=1
//ff:what writePlanComponentImports — plan 별 Controller/Service import 문 출력

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// writePlanComponentImports writes the controller and service import statements
// for each plan.
func writePlanComponentImports(b *strings.Builder, plans []*ir.ServicePlan) {
	for _, p := range plans {
		baseName := lcFirst(p.OperationID)
		b.WriteString(fmt.Sprintf("import { %sController } from './%s.controller';\n",
			p.OperationID, baseName))
		b.WriteString(fmt.Sprintf("import { %sService } from './%s.service';\n",
			p.OperationID, baseName))
	}
}
