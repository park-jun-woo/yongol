//ff:func feature=gen-nestjs type=util control=sequence
//ff:what writeModuleExportsSection — @Module exports: [...] 배열 출력 (cross-module DI 용)

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// writeModuleExportsSection writes the exports array of the @Module decorator.
// All services are exported so cross-module DI works.
func writeModuleExportsSection(b *strings.Builder, plans []*ir.ServicePlan, deps moduleDeps, stubSvcName string) {
	b.WriteString("  exports: [\n")
	writePlanServiceRefs(b, plans)
	if deps.NeedsSameFeatureStub {
		b.WriteString(fmt.Sprintf("    %s,\n", stubSvcName))
	}
	b.WriteString("  ],\n")
}
