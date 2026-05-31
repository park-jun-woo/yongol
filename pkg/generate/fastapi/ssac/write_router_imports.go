//ff:func feature=gen-fastapi type=util control=sequence
//ff:what writeRouterImports — FastAPI Router 파일 상단 import 블록 출력

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// writeRouterImports writes the import block of a FastAPI router file,
// including conditional auth/event-bus/service/schema imports.
func writeRouterImports(b *strings.Builder, feature string, plans []*ir.ServicePlan, needsAuth, needsEventBus bool) {
	b.WriteString("from fastapi import APIRouter, Depends\n")
	b.WriteString("from sqlalchemy.ext.asyncio import AsyncSession\n")
	b.WriteString("from app.dependencies.database import get_session\n")
	if needsAuth {
		b.WriteString("from app.dependencies.auth import get_current_user\n")
	}
	if needsEventBus {
		b.WriteString("from app.dependencies.event_bus import EventBus, get_event_bus\n")
	}
	b.WriteString(fmt.Sprintf("from app.services import %s as svc\n", feature))

	schemaModels := collectSchemaModels(plans)
	if len(schemaModels) > 0 {
		b.WriteString(fmt.Sprintf("from app.schemas.%s import %s\n",
			feature, strings.Join(schemaModels, ", ")))
	}
	b.WriteString("\n")
}
