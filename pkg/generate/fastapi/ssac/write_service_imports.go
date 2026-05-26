//ff:func feature=gen-fastapi type=util control=sequence
//ff:what writeServiceImports — FastAPI service 파일 import 문 작성

package ssac

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// writeServiceImports writes the import statements for a service file.
func writeServiceImports(b *strings.Builder, plan *ir.ServicePlan) {
	b.WriteString("from fastapi import HTTPException\n")
	b.WriteString("from sqlalchemy import select, update, delete\n")
	b.WriteString("from sqlalchemy.ext.asyncio import AsyncSession\n")
	if hasPublishOp(plan.Ops) {
		b.WriteString("from app.dependencies.event_bus import EventBus\n")
	}
	b.WriteString("\n")
}
