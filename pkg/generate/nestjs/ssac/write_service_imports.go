//ff:func feature=gen-nestjs type=util control=iteration dimension=1
//ff:what writeServiceImports — NestJS service 파일 import 문 작성 (../ 경로)

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// writeServiceImports writes the import statements for a service file.
// Import paths use '../' (one level up from src/<feature>/ to src/).
func writeServiceImports(b *strings.Builder, plan *ir.ServicePlan) {
	b.WriteString("import { Injectable, HttpException, HttpStatus } from '@nestjs/common';\n")
	b.WriteString("import { PrismaService } from '../prisma/prisma.service';\n")
	if hasPublishOp(plan.Ops) {
		b.WriteString("import { QueueService } from '../queue/queue.service';\n")
	}
	if hasAuthOp(plan.Ops) {
		b.WriteString("import { AuthzService } from '../authz/authz.service';\n")
	}
	for _, pkg := range collectExternalOpsPackages(plan.Ops) {
		svcName := strings.ToUpper(pkg[:1]) + pkg[1:] + "Service"
		b.WriteString(fmt.Sprintf("import { %s } from '../%s/%s.service';\n", svcName, pkg, pkg))
	}
	b.WriteString("\n")
}
