//ff:func feature=gen-nestjs type=util control=sequence
//ff:what writeControllerImports — NestJS controller 파일 import 문 작성

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// writeControllerImports writes the import statements for a controller file.
func writeControllerImports(b *strings.Builder, plan *ir.ServicePlan) {
	b.WriteString("import {\n")
	b.WriteString("  Controller,\n")
	if plan.TriggerKind == ir.TriggerHTTP {
		b.WriteString(fmt.Sprintf("  %s,\n", nestHTTPDecorator(plan.HTTPMethod)))
		b.WriteString("  Param,\n")
		b.WriteString("  Body,\n")
		b.WriteString("  Req,\n")
	}
	b.WriteString("} from '@nestjs/common';\n")

	serviceName := plan.OperationID + "Service"
	b.WriteString(fmt.Sprintf("import { %s } from './%s.service';\n\n",
		serviceName, lcFirst(plan.OperationID)))
}
