//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestWriteRouterImports — writeRouterImports auth/event_bus/schema 조건부 import 블록 출력 검증
package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestWriteRouterImports(t *testing.T) {
	t.Run("FullImports", func(t *testing.T) {
		plans := []*ir.ServicePlan{
			{OperationID: "CreateOrder", HTTPMethod: "POST", BodyFields: []ir.BodyFieldMeta{{Name: "title"}}},
		}
		var b strings.Builder
		writeRouterImports(&b, "order", plans, true, true)
		got := b.String()
		wants := []string{
			"from fastapi import APIRouter, Depends",
			"from app.dependencies.database import get_session",
			"from app.dependencies.auth import get_current_user",
			"from app.dependencies.event_bus import EventBus, get_event_bus",
			"from app.services import order as svc",
			"from app.schemas.order import CreateOrderRequest",
		}
		for _, w := range wants {
			if !strings.Contains(got, w) {
				t.Errorf("expected %q in:\n%s", w, got)
			}
		}
	})

	t.Run("MinimalImports", func(t *testing.T) {
		plans := []*ir.ServicePlan{
			{OperationID: "ListOrders", HTTPMethod: "GET"},
		}
		var b strings.Builder
		writeRouterImports(&b, "order", plans, false, false)
		got := b.String()
		if strings.Contains(got, "get_current_user") {
			t.Errorf("expected no auth import, got:\n%s", got)
		}
		if strings.Contains(got, "event_bus") {
			t.Errorf("expected no event_bus import, got:\n%s", got)
		}
		if strings.Contains(got, "from app.schemas") {
			t.Errorf("expected no schema import, got:\n%s", got)
		}
	})
}
