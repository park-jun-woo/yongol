//ff:func feature=gen-nestjs type=generator control=sequence
//ff:what RenderModule — feature 단위 NestJS module TypeScript 소스 생성 (../ 경로)

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// RenderModule produces a NestJS module file that imports the controllers
// and services for a given feature. Each ServicePlan contributes one
// controller and one service to the module.
func RenderModule(feature string, plans []*ir.ServicePlan) (string, error) {
	if feature == "" {
		return "", fmt.Errorf("RenderModule: empty feature name")
	}

	deps := collectModuleDeps(feature, plans)
	stubSvcName := strings.ToUpper(feature[:1]) + feature[1:] + "Service"
	moduleName := strings.ToUpper(feature[:1]) + feature[1:] + "Module"

	var b strings.Builder
	writeModuleImports(&b, feature, plans, deps, stubSvcName)
	writeModuleDecorator(&b, plans, deps, stubSvcName)
	b.WriteString(fmt.Sprintf("export class %s {}\n", moduleName))

	return b.String(), nil
}
