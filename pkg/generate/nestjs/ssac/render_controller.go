//ff:func feature=gen-nestjs type=generator control=sequence
//ff:what RenderController — ServicePlan → NestJS controller class TypeScript 소스 생성

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// RenderController takes a ServicePlan and produces a NestJS controller
// class with route decorators. Subscribe-triggered plans produce a
// placeholder handler comment instead of HTTP route decorators.
func RenderController(plan *ir.ServicePlan) (string, error) {
	if plan == nil {
		return "", fmt.Errorf("RenderController: nil plan")
	}

	var b strings.Builder
	writeControllerImports(&b, plan)
	writeControllerClass(&b, plan)
	return b.String(), nil
}
