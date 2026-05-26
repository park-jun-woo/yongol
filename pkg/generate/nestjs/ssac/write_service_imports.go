//ff:func feature=gen-nestjs type=util control=sequence
//ff:what writeServiceImports — NestJS service 파일 import 문 작성

package ssac

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// writeServiceImports writes the import statements for a service file.
func writeServiceImports(b *strings.Builder, plan *ir.ServicePlan) {
	b.WriteString("import { Injectable, HttpException, HttpStatus } from '@nestjs/common';\n")
	b.WriteString("import { PrismaService } from '../../prisma/prisma.service';\n")
	if hasPublishOp(plan.Ops) {
		b.WriteString("import { QueueService } from '../../queue/queue.service';\n")
	}
	b.WriteString("\n")
}
