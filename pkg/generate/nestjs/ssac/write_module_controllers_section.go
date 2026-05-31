//ff:func feature=gen-nestjs type=util control=iteration dimension=1
//ff:what writeModuleControllersSection — @Module controllers: [...] 배열 출력

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// writeModuleControllersSection writes the controllers array of the @Module
// decorator, one entry per plan.
func writeModuleControllersSection(b *strings.Builder, plans []*ir.ServicePlan) {
	b.WriteString("  controllers: [\n")
	for _, p := range plans {
		b.WriteString(fmt.Sprintf("    %sController,\n", p.OperationID))
	}
	b.WriteString("  ],\n")
}
