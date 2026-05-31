//ff:func feature=gen-nestjs type=util control=sequence
//ff:what writeModuleProvidersSection — @Module providers: [...] 배열 출력 (stub service 포함)

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// writeModuleProvidersSection writes the providers array of the @Module
// decorator: one service per plan plus the optional same-feature stub service.
func writeModuleProvidersSection(b *strings.Builder, plans []*ir.ServicePlan, deps moduleDeps, stubSvcName string) {
	b.WriteString("  providers: [\n")
	writePlanServiceRefs(b, plans)
	if deps.NeedsSameFeatureStub {
		b.WriteString(fmt.Sprintf("    %s,\n", stubSvcName))
	}
	b.WriteString("  ],\n")
}
