//ff:func feature=gen-nestjs type=generator control=sequence
//ff:what RenderService — ServicePlan → NestJS service class TypeScript 소스 생성

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// RenderService takes a ServicePlan and produces a TypeScript service class
// file content. Each Op is rendered into the corresponding Prisma query,
// guard, or external call.
func RenderService(plan *ir.ServicePlan, reg ir.TypeRegistry) (string, error) {
	if plan == nil {
		return "", fmt.Errorf("RenderService: nil plan")
	}
	_ = reg // reserved for future type conversion use

	var b strings.Builder
	writeServiceImports(&b, plan)
	writeServiceClass(&b, plan)
	return b.String(), nil
}
