//ff:func feature=gen-fastapi type=generator control=iteration dimension=1
//ff:what RenderRouter — feature 단위 FastAPI Router Python 소스 생성

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// RenderRouter produces a FastAPI router file for a given feature. Each
// ServicePlan contributes one route handler decorated with the appropriate
// HTTP method decorator. Parameters are typed using PathParams, QueryParams,
// and BodyFields from the ServicePlan.
func RenderRouter(feature string, plans []*ir.ServicePlan) (string, error) {
	if feature == "" {
		return "", fmt.Errorf("RenderRouter: empty feature name")
	}

	var b strings.Builder

	b.WriteString("from fastapi import APIRouter, Depends\n")
	b.WriteString("from sqlalchemy.ext.asyncio import AsyncSession\n")
	b.WriteString("from app.dependencies.database import get_session\n")
	b.WriteString("from app.dependencies.auth import get_current_user\n")

	// Import the service module
	b.WriteString(fmt.Sprintf("from app.services import %s as svc\n", feature))
	b.WriteString("\n")

	b.WriteString(fmt.Sprintf("router = APIRouter(prefix=\"/%s\", tags=[\"%s\"])\n\n", feature, feature))

	for _, plan := range plans {
		if plan.TriggerKind == ir.TriggerHTTP {
			writeHTTPHandler(&b, plan)
		} else {
			writeSubscribeHandler(&b, plan)
		}
		b.WriteString("\n")
	}

	return b.String(), nil
}
