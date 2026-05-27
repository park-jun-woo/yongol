//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestRenderFeatureSchemas — Pydantic BaseModel 스키마 생성 검증

package fastapi

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderFeatureSchemas(t *testing.T) {
	t.Run("POSTWithBody", func(t *testing.T) {
		plans := []*ir.ServicePlan{
			{
				OperationID: "CreateWorkflow",
				HTTPMethod:  "POST",
				TriggerKind: ir.TriggerHTTP,
				BodyFields: []ir.BodyFieldMeta{
					{Name: "title", Required: true},
					{Name: "description", Required: false},
				},
			},
		}
		got := renderFeatureSchemas(plans)
		if !strings.Contains(got, "from pydantic import BaseModel") {
			t.Errorf("expected pydantic import, got: %s", got)
		}
		if !strings.Contains(got, "class CreateWorkflowRequest(BaseModel):") {
			t.Errorf("expected class definition, got: %s", got)
		}
		if !strings.Contains(got, "title: str") {
			t.Errorf("expected required field, got: %s", got)
		}
		if !strings.Contains(got, "description: Optional[str] = None") {
			t.Errorf("expected optional field, got: %s", got)
		}
	})

	t.Run("GETNoBody", func(t *testing.T) {
		plans := []*ir.ServicePlan{
			{
				OperationID: "GetWorkflow",
				HTTPMethod:  "GET",
				TriggerKind: ir.TriggerHTTP,
			},
		}
		got := renderFeatureSchemas(plans)
		if got != "" {
			t.Errorf("expected empty for GET plan, got: %s", got)
		}
	})

	t.Run("SubscribeNoBody", func(t *testing.T) {
		plans := []*ir.ServicePlan{
			{
				OperationID: "HandleEvent",
				TriggerKind: ir.TriggerSubscribe,
				Topic:       "event.created",
			},
		}
		got := renderFeatureSchemas(plans)
		if got != "" {
			t.Errorf("expected empty for subscribe plan, got: %s", got)
		}
	})
}
