//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestWriteHTTPHandlerTyped — PathParams/BodyFields/QueryParams 기반 typed 핸들러 검증

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestWriteHTTPHandlerTyped(t *testing.T) {
	t.Run("GETOnlyPath", func(t *testing.T) {
		plan := &ir.ServicePlan{
			OperationID: "GetWorkflow",
			HTTPMethod:  "GET",
			TriggerKind: ir.TriggerHTTP,
			URLPath:     "/workflow/:id",
			Feature:     "workflow",
			PathParams:  []string{"id"},
		}
		var b strings.Builder
		writeHTTPHandler(&b, plan)
		got := b.String()
		if !strings.Contains(got, "id: int,") {
			t.Errorf("GET should have typed path param, got: %s", got)
		}
		if strings.Contains(got, "body:") {
			t.Errorf("GET should not have body, got: %s", got)
		}
		if !strings.Contains(got, "Depends(get_current_user)") {
			t.Errorf("expected Depends(get_current_user), got: %s", got)
		}
		if !strings.Contains(got, "Depends(get_session)") {
			t.Errorf("expected Depends(get_session), got: %s", got)
		}
	})

	t.Run("GETWithQuery", func(t *testing.T) {
		plan := &ir.ServicePlan{
			OperationID: "ListWorkflows",
			HTTPMethod:  "GET",
			TriggerKind: ir.TriggerHTTP,
			URLPath:     "/workflow",
			Feature:     "workflow",
			QueryParams: []ir.QueryParamMeta{
				{Name: "per_page", Type: "integer"},
				{Name: "cursor", Type: "string", Required: false},
			},
		}
		var b strings.Builder
		writeHTTPHandler(&b, plan)
		got := b.String()
		if !strings.Contains(got, "per_page: int | None = None") {
			t.Errorf("expected optional int query param, got: %s", got)
		}
		if !strings.Contains(got, "cursor: str | None = None") {
			t.Errorf("expected optional str query param, got: %s", got)
		}
	})

	t.Run("POSTWithBody", func(t *testing.T) {
		plan := &ir.ServicePlan{
			OperationID: "CreateWorkflow",
			HTTPMethod:  "POST",
			TriggerKind: ir.TriggerHTTP,
			URLPath:     "/workflow",
			Feature:     "workflow",
			BodyFields: []ir.BodyFieldMeta{
				{Name: "title"},
			},
		}
		var b strings.Builder
		writeHTTPHandler(&b, plan)
		got := b.String()
		if !strings.Contains(got, "body: CreateWorkflowRequest") {
			t.Errorf("POST with body should have Pydantic model, got: %s", got)
		}
	})

	t.Run("PUTWithPathAndBody", func(t *testing.T) {
		plan := &ir.ServicePlan{
			OperationID: "UpdateWorkflow",
			HTTPMethod:  "PUT",
			TriggerKind: ir.TriggerHTTP,
			URLPath:     "/workflow/:id",
			Feature:     "workflow",
			PathParams:  []string{"id"},
			BodyFields: []ir.BodyFieldMeta{
				{Name: "title"},
				{Name: "status"},
			},
		}
		var b strings.Builder
		writeHTTPHandler(&b, plan)
		got := b.String()
		if !strings.Contains(got, "id: int,") {
			t.Errorf("PUT should have path param, got: %s", got)
		}
		if !strings.Contains(got, "body: UpdateWorkflowRequest") {
			t.Errorf("PUT should have body, got: %s", got)
		}
	})

	t.Run("DELETENoBody", func(t *testing.T) {
		plan := &ir.ServicePlan{
			OperationID: "DeleteWorkflow",
			HTTPMethod:  "DELETE",
			TriggerKind: ir.TriggerHTTP,
			URLPath:     "/workflow/:id",
			Feature:     "workflow",
			PathParams:  []string{"id"},
		}
		var b strings.Builder
		writeHTTPHandler(&b, plan)
		got := b.String()
		if !strings.Contains(got, "id: int,") {
			t.Errorf("DELETE should have path param, got: %s", got)
		}
		if strings.Contains(got, "body:") {
			t.Errorf("DELETE should not have body, got: %s", got)
		}
	})
}
