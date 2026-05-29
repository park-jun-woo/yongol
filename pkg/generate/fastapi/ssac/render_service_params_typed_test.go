//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestRenderServiceParamsTyped — ServicePlan 메타데이터 기반 typed 파라미터 목록 검증

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderServiceParamsTyped(t *testing.T) {
	t.Run("GETWithPath", func(t *testing.T) {
		plan := &ir.ServicePlan{
			OperationID: "GetWorkflow",
			HTTPMethod:  "GET",
			TriggerKind: ir.TriggerHTTP,
			PathParams:  []string{"id"},
		}
		got := renderServiceParams(plan)
		if !strings.Contains(got, "session: AsyncSession") {
			t.Errorf("expected session param, got: %s", got)
		}
		if !strings.Contains(got, "id: int") {
			t.Errorf("expected typed path param, got: %s", got)
		}
		if !strings.Contains(got, "current_user: dict | None = None") {
			t.Errorf("expected current_user, got: %s", got)
		}
	})

	t.Run("POSTWithBody", func(t *testing.T) {
		plan := &ir.ServicePlan{
			OperationID: "CreateWorkflow",
			HTTPMethod:  "POST",
			TriggerKind: ir.TriggerHTTP,
			BodyFields: []ir.BodyFieldMeta{
				{Name: "title"},
			},
		}
		got := renderServiceParams(plan)
		if !strings.Contains(got, "body: CreateWorkflowRequest") {
			t.Errorf("expected Pydantic model body, got: %s", got)
		}
	})

	t.Run("GETWithQueryParams", func(t *testing.T) {
		plan := &ir.ServicePlan{
			OperationID: "ListWorkflows",
			HTTPMethod:  "GET",
			TriggerKind: ir.TriggerHTTP,
			QueryParams: []ir.QueryParamMeta{
				{Name: "per_page", Type: "integer", Required: false},
				{Name: "status", Type: "string", Required: true},
			},
		}
		got := renderServiceParams(plan)
		if !strings.Contains(got, "per_page: int | None = None") {
			t.Errorf("expected optional int query param, got: %s", got)
		}
		if !strings.Contains(got, "status: str") {
			t.Errorf("expected required str query param, got: %s", got)
		}
	})

	t.Run("SubscribeTrigger", func(t *testing.T) {
		plan := &ir.ServicePlan{
			OperationID: "HandleExecuted",
			TriggerKind: ir.TriggerSubscribe,
			Topic:       "workflow.executed",
		}
		got := renderServiceParams(plan)
		if !strings.Contains(got, "session: AsyncSession, payload: dict") {
			t.Errorf("expected subscribe params, got: %s", got)
		}
	})
}
