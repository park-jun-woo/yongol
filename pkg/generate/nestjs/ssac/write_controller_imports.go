//ff:func feature=gen-nestjs type=util control=sequence
//ff:what writeControllerImports — ServicePlan 메타데이터 기반 NestJS controller import 문 작성

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// writeControllerImports writes the import statements for a controller file.
// Imports @Query when QueryParams exist, @Body for POST/PUT/PATCH with
// BodyFields, and @Param for path parameters.
func writeControllerImports(b *strings.Builder, plan *ir.ServicePlan) {
	method := strings.ToUpper(plan.HTTPMethod)
	hasBody := (method == "POST" || method == "PUT" || method == "PATCH") && len(plan.BodyFields) > 0
	hasQuery := len(plan.QueryParams) > 0
	hasPath := len(plan.PathParams) > 0

	b.WriteString("import {\n")
	b.WriteString("  Controller,\n")
	if plan.TriggerKind == ir.TriggerHTTP {
		b.WriteString(fmt.Sprintf("  %s,\n", nestHTTPDecorator(plan.HTTPMethod)))
		if hasPath {
			b.WriteString("  Param,\n")
		}
		if hasBody {
			b.WriteString("  Body,\n")
		}
		if hasQuery {
			b.WriteString("  Query,\n")
		}
		b.WriteString("  Req,\n")
	}
	b.WriteString("} from '@nestjs/common';\n")

	serviceName := plan.OperationID + "Service"
	b.WriteString(fmt.Sprintf("import { %s } from './%s.service';\n\n",
		serviceName, lcFirst(plan.OperationID)))
}
