//ff:func feature=gen-nestjs type=util control=sequence
//ff:what writeModuleImports — NestJS module 파일 상단 import 블록 출력

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// writeModuleImports writes the import statements of a NestJS module file:
// base modules, conditional queue/authz, cross-feature modules, per-plan
// controllers/services, and the optional same-feature stub service.
func writeModuleImports(b *strings.Builder, feature string, plans []*ir.ServicePlan, deps moduleDeps, stubSvcName string) {
	b.WriteString("import { Module } from '@nestjs/common';\n")
	b.WriteString("import { PrismaModule } from '../prisma/prisma.module';\n")
	if deps.NeedsQueue {
		b.WriteString("import { QueueModule } from '../queue/queue.module';\n")
	}
	if deps.NeedsAuthz {
		b.WriteString("import { AuthzModule } from '../authz/authz.module';\n")
	}
	writeCrossFeatureImports(b, deps.CrossFeatures)
	writePlanComponentImports(b, plans)
	if deps.NeedsSameFeatureStub {
		b.WriteString(fmt.Sprintf("import { %s } from './%s.service';\n",
			stubSvcName, strings.ToLower(feature)))
	}
	b.WriteString("\n")
}
