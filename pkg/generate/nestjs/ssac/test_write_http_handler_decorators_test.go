//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestWriteHTTPHandlerDecorators — @Param/@Body/@Query 데코레이터 분기 검증

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestWriteHTTPHandlerDecorators(t *testing.T) {
	t.Run("GETOnlyParam", func(t *testing.T) {
		plan := &ir.ServicePlan{
			OperationID: "GetWorkflow",
			HTTPMethod:  "GET",
			TriggerKind: ir.TriggerHTTP,
			PathParams:  []string{"id"},
		}
		var b strings.Builder
		writeHTTPHandler(&b, plan)
		got := b.String()
		if !strings.Contains(got, "@Param()") {
			t.Errorf("GET should have @Param, got: %s", got)
		}
		if strings.Contains(got, "@Body()") {
			t.Errorf("GET should not have @Body, got: %s", got)
		}
		if strings.Contains(got, "@Query()") {
			t.Errorf("GET without query params should not have @Query, got: %s", got)
		}
	})

	t.Run("GETWithQuery", func(t *testing.T) {
		plan := &ir.ServicePlan{
			OperationID: "ListWorkflows",
			HTTPMethod:  "GET",
			TriggerKind: ir.TriggerHTTP,
			QueryParams: []ir.QueryParamMeta{
				{Name: "per_page", Type: "integer"},
			},
		}
		var b strings.Builder
		writeHTTPHandler(&b, plan)
		got := b.String()
		if strings.Contains(got, "@Param()") {
			t.Errorf("no path params should not have @Param, got: %s", got)
		}
		if !strings.Contains(got, "@Query()") {
			t.Errorf("GET with query params should have @Query, got: %s", got)
		}
	})

	t.Run("POSTWithBody", func(t *testing.T) {
		plan := &ir.ServicePlan{
			OperationID: "CreateWorkflow",
			HTTPMethod:  "POST",
			TriggerKind: ir.TriggerHTTP,
			PathParams:  []string{},
			BodyFields: []ir.BodyFieldMeta{
				{Name: "title"},
			},
		}
		var b strings.Builder
		writeHTTPHandler(&b, plan)
		got := b.String()
		if !strings.Contains(got, "@Body()") {
			t.Errorf("POST with body should have @Body, got: %s", got)
		}
	})

	t.Run("PUTWithParamAndBody", func(t *testing.T) {
		plan := &ir.ServicePlan{
			OperationID: "UpdateWorkflow",
			HTTPMethod:  "PUT",
			TriggerKind: ir.TriggerHTTP,
			PathParams:  []string{"id"},
			BodyFields: []ir.BodyFieldMeta{
				{Name: "title"},
				{Name: "status"},
			},
		}
		var b strings.Builder
		writeHTTPHandler(&b, plan)
		got := b.String()
		if !strings.Contains(got, "@Param()") {
			t.Errorf("PUT should have @Param, got: %s", got)
		}
		if !strings.Contains(got, "@Body()") {
			t.Errorf("PUT should have @Body, got: %s", got)
		}
	})

	t.Run("DELETENoBody", func(t *testing.T) {
		plan := &ir.ServicePlan{
			OperationID: "DeleteWorkflow",
			HTTPMethod:  "DELETE",
			TriggerKind: ir.TriggerHTTP,
			PathParams:  []string{"id"},
		}
		var b strings.Builder
		writeHTTPHandler(&b, plan)
		got := b.String()
		if !strings.Contains(got, "@Param()") {
			t.Errorf("DELETE should have @Param, got: %s", got)
		}
		if strings.Contains(got, "@Body()") {
			t.Errorf("DELETE should not have @Body, got: %s", got)
		}
	})
}
