//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestEmitSchemaImports — Pydantic 스키마 import 출력 검증

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestEmitSchemaImports(t *testing.T) {
	t.Run("POSTWithBody", func(t *testing.T) {
		plans := []*ir.ServicePlan{
			{
				OperationID: "Login",
				HTTPMethod:  "POST",
				Feature:     "auth",
				BodyFields: []ir.BodyFieldMeta{
					{Name: "email"},
					{Name: "password"},
				},
			},
		}
		var b strings.Builder
		emitSchemaImports(&b, plans, "auth")
		got := b.String()
		if !strings.Contains(got, "from app.schemas.auth import LoginRequest") {
			t.Errorf("expected LoginRequest import, got:\n%s", got)
		}
	})

	t.Run("GETNoBody", func(t *testing.T) {
		plans := []*ir.ServicePlan{
			{
				OperationID: "GetProfile",
				HTTPMethod:  "GET",
				Feature:     "auth",
			},
		}
		var b strings.Builder
		emitSchemaImports(&b, plans, "auth")
		got := b.String()
		if got != "" {
			t.Errorf("expected no schema import for GET, got:\n%s", got)
		}
	})

	t.Run("MultipleBodies", func(t *testing.T) {
		plans := []*ir.ServicePlan{
			{
				OperationID: "CreateWorkflow",
				HTTPMethod:  "POST",
				Feature:     "workflow",
				BodyFields:  []ir.BodyFieldMeta{{Name: "title"}},
			},
			{
				OperationID: "UpdateWorkflow",
				HTTPMethod:  "PUT",
				Feature:     "workflow",
				BodyFields:  []ir.BodyFieldMeta{{Name: "title"}, {Name: "status"}},
			},
		}
		var b strings.Builder
		emitSchemaImports(&b, plans, "workflow")
		got := b.String()
		if !strings.Contains(got, "CreateWorkflowRequest") {
			t.Errorf("expected CreateWorkflowRequest, got:\n%s", got)
		}
		if !strings.Contains(got, "UpdateWorkflowRequest") {
			t.Errorf("expected UpdateWorkflowRequest, got:\n%s", got)
		}
	})

	t.Run("DELETENoSchema", func(t *testing.T) {
		plans := []*ir.ServicePlan{
			{
				OperationID: "DeleteWorkflow",
				HTTPMethod:  "DELETE",
				Feature:     "workflow",
				PathParams:  []string{"id"},
			},
		}
		var b strings.Builder
		emitSchemaImports(&b, plans, "workflow")
		got := b.String()
		if got != "" {
			t.Errorf("expected no schema import for DELETE, got:\n%s", got)
		}
	})
}
