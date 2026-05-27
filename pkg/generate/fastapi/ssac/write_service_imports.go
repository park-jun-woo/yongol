//ff:func feature=gen-fastapi type=util control=sequence
//ff:what writeServiceImports — FastAPI service 파일 통합 import 문 작성

package ssac

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// writeServiceImports writes the consolidated import statements for a feature
// service file. It scans all plans to collect referenced model classes and
// external package functions, emitting each import exactly once.
func writeServiceImports(b *strings.Builder, plans []*ir.ServicePlan) {
	d := collectImportData(plans)
	b.WriteString("from fastapi import HTTPException\n")
	emitSAImports(b, d)
	b.WriteString("from sqlalchemy.ext.asyncio import AsyncSession\n")
	emitModelImports(b, d.Models)
	emitInfraImports(b, d)
	emitExtPkgImports(b, d.ExtPkgs)
	b.WriteString("\n")
}
