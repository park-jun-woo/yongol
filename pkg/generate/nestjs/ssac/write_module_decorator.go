//ff:func feature=gen-nestjs type=util control=sequence
//ff:what writeModuleDecorator — @Module({...}) 데코레이터(imports/controllers/providers/exports) 출력

package ssac

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// writeModuleDecorator writes the @Module decorator with its imports,
// controllers, providers, and exports sections.
func writeModuleDecorator(b *strings.Builder, plans []*ir.ServicePlan, deps moduleDeps, stubSvcName string) {
	b.WriteString("@Module({\n")
	writeModuleImportsSection(b, deps)
	writeModuleControllersSection(b, plans)
	writeModuleProvidersSection(b, plans, deps, stubSvcName)
	writeModuleExportsSection(b, plans, deps, stubSvcName)
	b.WriteString("})\n")
}
